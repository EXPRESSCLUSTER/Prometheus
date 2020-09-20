package collector

import (
	"strconv"
	"fmt"
	"os/exec"
	"strings"
	"unsafe"

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
	var metricType prometheus.ValueType
	metricType = prometheus.GaugeValue
	out, err := exec.Command("/bin/sh", "./getelaps.sh").Output()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, "test"),
				fmt.Sprintf("Monitor"),	
				nil, nil,
			),
			metricType, float64(0),
		)
	} else {
		f, err := strconv.ParseFloat(strings.TrimRight(*(*string)(unsafe.Pointer(&out)), "\n"), 64)
		fmt.Printf("f: %v\n", float64(f))
		if err != nil {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "test"),
					fmt.Sprintf("Monitor"),	
					nil, nil,
				),
				metricType, float64(0),
			)
		} else {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "test"),
					fmt.Sprintf("Monitor"),	
					nil, nil,
				),
				metricType, float64(f),
			)
		}
	}

	return nil
}

