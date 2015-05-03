package telemetry

import (
	"encoding/json"
	"net/http"
)

type httpMetrics struct {
	metrics map[string]metricInterface
}

func (h httpMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	container := make(map[string]map[string]float32)
	for key, metric := range h.metrics {
		container[key] = metric.GetAll()
	}
	output, err := json.MarshalIndent(container, "", "  ")
	if err != nil {
		w.Write([]byte("failed to encode JSON"))
		return
	}
	w.Write(output)
}
