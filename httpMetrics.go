package telemetry

import (
	"fmt"
	"net/http"
)

type httpMetrics struct {
	metrics *dataPoints
}

func (h httpMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key, value := range h.metrics.floatMetrics {
		w.Write([]byte(fmt.Sprintf("%s %v \n", key, value)))
	}
	for key, value := range h.metrics.intMetrics {
		w.Write([]byte(fmt.Sprintf("%s %v \n", key, value)))
	}
}
