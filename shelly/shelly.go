package shelly

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aexel90/shelly_exporter/metric"
	"github.com/tidwall/gjson"
)

// Exporter data
type Exporter struct {
	Products []*metric.Products
}

// Collect metrics
func (exporter *Exporter) Collect(metrics []*metric.Metric) (err error) {

	productResults := make(map[string][]map[string]interface{})
	for _, product := range exporter.Products {

		deviceResults := []map[string]interface{}{}
		for _, device := range product.Devices {
			apiURL := "https://" + device.URL + "/device/Status" + "&auth_key=" + device.AuthKey + "&id=" + device.ID
			response, err := http.Get(apiURL)
			if err != nil {
				return err
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				return fmt.Errorf("Response not OK: %v", response.Status)
			}

			jsonResponse, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			jsonString := string(jsonResponse[:])

			result := make(map[string]interface{})
			for key, value := range product.Export {
				result[key] = gjson.Get(jsonString, value)
			}

			result["name"] = device.Name
			result["shelly_name"] = device.ShellyName

			deviceResults = append(deviceResults, result)
		}
		productResults[product.Type] = deviceResults
	}

	for _, metric := range metrics {
		metric.MetricResult = productResults[metric.ShellyType]
	}
	return nil
}
