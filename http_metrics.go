package telemetry

import (
	"encoding/json"
	"net/http"
)

// HTTPMetrics http.Handler for telemetry data
type HTTPMetrics struct {
	Metrics map[string]MetricInterface
}

//MetricsHandler returns JSON encoded output of application metrics
func (h HTTPMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	container := make(map[string]map[string]float32)
	for key, metric := range h.Metrics {
		container[key] = metric.GetAll()
	}
	output, _ := json.MarshalIndent(container, "", "  ")
	w.Write(output)
}
