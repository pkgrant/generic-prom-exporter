/*
 * Exports the metrics into Prometheus format
 */

package collector

import (

    "fmt"
    "strings"

    "static-exporter/client"
    "github.com/prometheus/client_golang/prometheus"
)

type thresholdValues struct {
    firstValue     *prometheus.Desc
    secondValue    *prometheus.Desc
    thirdValue     *prometheus.Desc
    fourthValue    *prometheus.Desc
}

type recordedValues struct {
    firstValue     int
    secondValue    int
    thirdValue     int
    fourthValue    int
}

type Collector struct {
    up         *prometheus.Desc
    thresholds thresholdValues
}

func newThresholdValueMetrics(item string, labels ...string) thresholdValues {
    return thresholdValues{
        firstValue: prometheus.NewDesc(fmt.Sprintf("first_%s_test_value", strings.ToLower(item)),
            fmt.Sprintf("%s first value", item), labels, nil),
        secondValue: prometheus.NewDesc(fmt.Sprintf("second_%s_test_value", strings.ToLower(item)),
            fmt.Sprintf("%s second value", item), labels, nil),
        thirdValue: prometheus.NewDesc(fmt.Sprintf("third_%s_test_value", strings.ToLower(item)),
            fmt.Sprintf("%s third value", item), labels, nil),
        fourthValue: prometheus.NewDesc(fmt.Sprintf("fourth_%s_test_value", strings.ToLower(item)),
            fmt.Sprintf("%s fourth value", item), labels, nil),
    }
}

func New() *Collector {
    return &Collector {
        up: prometheus.NewDesc("staticexporter_up", "Whether the config file was read successfully", nil, nil),
        thresholds: newThresholdValueMetrics("threshold", "machine"),
    }
}

func describeThresholdValues(ch chan<- *prometheus.Desc, thresholds *thresholdValues) {
    ch <- thresholds.firstValue
    ch <- thresholds.secondValue
    ch <- thresholds.thirdValue
    ch <- thresholds.fourthValue
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {

    ch <- c.up

    describeThresholdValues(ch, &c.thresholds)
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
    if thresholds, err := client.GetThresholds(); err != nil {
        // client call failed, set the up metric value to 0
        ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
    } else {
        //client call succeeded, set the up metric value to 1
        ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)

        // ... collect other metrics ...
        fmt.Printf("80: %+v\n\n\n", thresholds)
        collectThresholds(c, ch, thresholds)
    }
}

func collectThresholds(c *Collector, ch chan<- prometheus.Metric, thresholds *client.Config) {
    fmt.Printf("85: %+v\n\n\n", *thresholds)
    for _, host := range thresholds.Targets {
        //fmt.Printf("%+v\n", host)
        currentThresh := recordedValues {
            firstValue:     host.Firstvalue,
            secondValue:    host.Secondvalue,
            thirdValue:     host.Thirdvalue,
            fourthValue:    host.Fourthvalue,
        }
        collectThresholdValues(ch, &c.thresholds, currentThresh, host.Host)
    }
}

func collectThresholdValues(ch chan<- prometheus.Metric, thresholds *thresholdValues, values recordedValues, labelValues ...string) {
    ch <- prometheus.MustNewConstMetric(thresholds.firstValue, prometheus.GaugeValue, float64(values.firstValue), labelValues...)
    ch <- prometheus.MustNewConstMetric(thresholds.secondValue, prometheus.GaugeValue, float64(values.fourthValue), labelValues...)
    ch <- prometheus.MustNewConstMetric(thresholds.thirdValue, prometheus.GaugeValue, float64(values.thirdValue), labelValues...)
    ch <- prometheus.MustNewConstMetric(thresholds.fourthValue, prometheus.GaugeValue, float64(values.fourthValue), labelValues...)
}
