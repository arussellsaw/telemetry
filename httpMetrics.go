package telemetry

import (
	"net/http"
)

type httpMetrics struct {
	metrics map[string]metricInterface
}

func (h httpMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, metric := range h.metrics {
		w.Write([]byte(metric.GetAll()))
	}
}
