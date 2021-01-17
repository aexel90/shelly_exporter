package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/aexel90/shelly_exporter/collector"
	"github.com/aexel90/shelly_exporter/metric"
)

var (
	flagAddress     = flag.String("listen-address", "127.0.0.1:9784", "The address to listen on for HTTP requests.") //TODO port?
	flagMetricsFile = flag.String("metrics-file", "shelly_metrics.json", "The JSON file with the metric definitions.")

	flagTest = flag.Bool("test", false, "Test configured metrics")
)

func main() {

	flag.Parse()

	var metricsFile *metric.MetricsFile

	err := readAndParseFile(*flagMetricsFile, &metricsFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	shellyCollector, err := collector.NewShellyCollector(metricsFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *flagTest {
		shellyCollector.Test()
	} else {
		prometheus.MustRegister(shellyCollector)
		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("metrics available at http://%s/metrics\n", *flagAddress)
		log.Fatal(http.ListenAndServe(*flagAddress, nil))
	}
}

func readAndParseFile(file string, v interface{}) error {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading metric file: %v", err)
	}

	err = json.Unmarshal(jsonData, v)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}
	return nil
}
