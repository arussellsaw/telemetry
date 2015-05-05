package telemetry

import (
	"encoding/json"
	"net/http"
)

//HTTPMetrics request handler for metrics
type HTTPMetrics struct {
	Metrics map[string]MetricInterface
}

//MetricsHandler
func (h HTTPMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	container := make(map[string]map[string]float32)
	for key, metric := range h.Metrics {
		container[key] = metric.GetAll()
	}
	output, err := json.MarshalIndent(container, "", "  ")
	if err != nil {
		w.Write([]byte("failed to encode JSON"))
		return
	}
	w.Write(output)
}
