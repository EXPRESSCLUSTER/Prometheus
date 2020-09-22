package collector

import (
	"bufio"
	"strconv"
	"fmt"
//	"io"
//	"log"
	"os"
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


	cmdstr := "clpmonctrl -v|grep Resource|awk '{print $3}'"
	cmd := exec.Command("sh", "-c", cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		monitor := scanner.Text()
		fmt.Printf("scan: %v\n", monitor)

		out, err := exec.Command("clpperfc", "-m", monitor).Output()
		fmt.Printf("ret: %s\n", out)
		if err != nil {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, monitor),
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
						prometheus.BuildFQName(namespace, subsystem, monitor),
						fmt.Sprintf("Monitor"),	
						nil, nil,
					),
					metricType, float64(0),
				)
			} else {
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(namespace, subsystem, monitor),
						fmt.Sprintf("Monitor"),	
						nil, nil,
					),
					metricType, float64(f),
				)
			}
		}
	}
	cmd.Wait()

	return nil
}

