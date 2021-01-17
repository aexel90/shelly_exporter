package collector

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"

	"github.com/aexel90/shelly_exporter/metric"
	"github.com/aexel90/shelly_exporter/shelly"
)

// Collector instance
type Collector struct {
	exporter *shelly.Exporter
	metrics  []*metric.Metric
}

// NewShellyCollector initialization
func NewShellyCollector(metricsFile *metric.MetricsFile) (*Collector, error) {

	hueExporter := shelly.Exporter{Products: metricsFile.Products}
	return &Collector{&hueExporter, metricsFile.Metrics}, nil
}

// Describe for prometheus
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	collector.initDescAndType()
}

// Collect for prometheus
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {

	err := collector.collect()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for _, m := range collector.metrics {
		for _, promResult := range m.PromResult {
			ch <- prometheus.MustNewConstMetric(promResult.PromDesc, promResult.PromValueType, promResult.Value, promResult.LabelValues...)
		}
	}
}

//Test collector metrics
func (collector *Collector) Test() {

	collector.initDescAndType()

	err := collector.collect()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = collector.printResult()
}

func (collector *Collector) printResult() error {

	for _, m := range collector.metrics {
		fmt.Printf("Metric: %v\n", m.FqName)
		fmt.Printf(" - Exporter Result:\n")

		for i, result := range m.MetricResult {
			fmt.Printf("   - Exporter Result %v:\n", i)
			for key, value := range result {
				fmt.Printf("     - %s=\"%v\"\n", key, value)
			}
		}

		for _, promResult := range m.PromResult {

			fmt.Printf("   - prom desc: %v\n", promResult.PromDesc)
			fmt.Printf("     - prom metric type: %v\n", promResult.PromValueType)
			fmt.Printf("     - prom metric value: %v\n", uint64(promResult.Value))
			fmt.Printf("     - prom label values: %v\n", promResult.LabelValues)
		}
	}

	return nil
}

func (collector *Collector) collect() (err error) {

	err = collector.exporter.Collect(collector.metrics)
	if err != nil {
		return err
	}

	err = collector.getResult()
	if err != nil {
		return err
	}
	return nil
}

func (collector *Collector) getResult() (err error) {

	for _, m := range collector.metrics {
		m.PromResult = nil
		for _, metricResult := range m.MetricResult {

			labelValues, err := getLabelValues(m.Labels, metricResult)
			if err != nil {
				return err
			}
			resultValue, err := getResultValue(m.ResultKey, metricResult)
			if err != nil {
				return err
			}
			if resultValue == nil {
				continue
			}

			result := metric.PrometheusResult{PromDesc: m.PromDesc, PromValueType: m.PromType, Value: *resultValue, LabelValues: labelValues}
			m.PromResult = append(m.PromResult, &result)
		}
	}
	return nil
}

func (collector *Collector) initDescAndType() {

	for _, metric := range collector.metrics {

		metric.PromType = prometheus.GaugeValue

		labels := []string{}
		for _, label := range metric.Labels {
			labels = append(labels, strings.ToLower(label))
		}

		metric.PromDesc = prometheus.NewDesc(metric.FqName, metric.Help, labels, nil)
	}
}

func getResultValue(resultKey string, result map[string]interface{}) (*float64, error) {

	var floatValue float64

	if resultKey == "" {
		floatValue = 1
		return &floatValue, nil
	}

	value := result[resultKey]
	if value == nil {
		return nil, nil
	}

	switch tval := value.(type) {
	case gjson.Result:
		floatValue = tval.Float()
	default:
		return nil, fmt.Errorf("[getResultValue] %v in %v - unknown type: %T", resultKey, result, value)
	}
	return &floatValue, nil
}

func getLabelValues(labelNames []string, result map[string]interface{}) ([]string, error) {

	labelValues := []string{}
	for _, labelname := range labelNames {

		labelValue := result[labelname]
		if labelValue == nil {
			labelValue = ""
		}

		labelValueString := fmt.Sprintf("%v", labelValue)
		switch labelValueString {
		case "true":
			labelValueString = "1"
		case "false":
			labelValueString = "0"
		}
		labelValues = append(labelValues, labelValueString)
	}
	return labelValues, nil
}
