package main

import (
	"strconv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "ClpMetric"
)

type myCollector struct{}

var (
	exampleCount = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "example_count",
		Help:      "example counter help",
	})
	exampleGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "example_gauge",
		Help:      "example gauge help",
	})
)

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- exampleCount.Desc()
	ch <- exampleGauge.Desc()
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/bin/sh", "./getelaps.sh").Output()
	if err != nil {
//		fmt.Println("ERROR!")
	}
	fmt.Println(out)
	fmt.Printf("%s\n", string(out))
	f, err := strconv.ParseFloat(strings.TrimRight(*(*string)(unsafe.Pointer(&out)), "\n"), 64)
	fmt.Println(f)

	ch <- prometheus.MustNewConstMetric(
		exampleCount.Desc(),
		prometheus.CounterValue,
		float64(f),
	)
	ch <- prometheus.MustNewConstMetric(
		exampleGauge.Desc(),
		prometheus.GaugeValue,
		float64(f),
	)
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
