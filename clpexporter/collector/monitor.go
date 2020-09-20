package collector

import (
//	"strconv"
//	"log"
//	"os/exec"
//	"strings"
//	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	subsystem = "monitor"
)

type monitorCollector struct{}

func init() {
	registCollector(subsystem, NewMonitorCollector)
}

func NewMonitorCollector() (Collector, error) {
	return &monitorCollector{}, nil
}

func (c *monitorCollector) Update(ch chan<- prometheus.Metric) error {
/*
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
*/
	return nil
}

