package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusResult struct
type PrometheusResult struct {
	PromDesc      *prometheus.Desc
	PromValueType prometheus.ValueType
	Value         float64
	LabelValues   []string
}

// Metric struct
type Metric struct {
	ShellyType string   `json:"type"`
	ResultKey  string   `json:"resultKey"`
	FqName     string   `json:"fqName"`
	Help       string   `json:"help"`
	Labels     []string `json:"labels"`

	MetricResult []map[string]interface{}

	PromType   prometheus.ValueType
	PromDesc   *prometheus.Desc
	PromResult []*PrometheusResult
}

// Device struct
type Device struct {
	ID         string `json:"id"`
	ShellyName string `json:"shelly_name"`
	Name       string `json:"name"`
}

// Products struct
type Products struct {
	Type    string            `json:"type"`
	Devices []*Device         `json:"devices"`
	Export  map[string]string `json:"export"`
}

// Account struct
type Account struct {
	AuthKey string `json:"auth_key"`
	URL     string `json:"url"`
}

// MetricsFile struct
type MetricsFile struct {
	Metrics  []*Metric   `json:"metrics"`
	Products []*Products `json:"products"`
	Account  *Account    `json:"account"`
}
