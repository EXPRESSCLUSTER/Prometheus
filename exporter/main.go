package main

import (
	"strconv"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "ecx"
)

type myCollector struct{}

var (
	exampleGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "monitor_elapsed_time",
		Help:      "monitor elapsed time  [msec]",
	})
)

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- exampleGauge.Desc()
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/bin/sh", "./getelaps.sh").Output()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			exampleGauge.Desc(),
			prometheus.GaugeValue,
			float64(0),
		)
	} else {
		f, err := strconv.ParseFloat(strings.TrimRight(*(*string)(unsafe.Pointer(&out)), "\n"), 64)
		if err != nil {
			ch <- prometheus.MustNewConstMetric(
				exampleGauge.Desc(),
				prometheus.GaugeValue,
				float64(0),
			)
		} else {
			ch <- prometheus.MustNewConstMetric(
				exampleGauge.Desc(),
				prometheus.GaugeValue,
				float64(f),
			)
		}
	}
}

var (
	addr = flag.String("listen address", ":9090", "The Address to listen on for HTTP Requests.")
)

func main() {
	flag.Parse()

	var c myCollector
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
