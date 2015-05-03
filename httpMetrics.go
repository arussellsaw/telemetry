package telemetry

import (
	"fmt"
	"net/http"
)

type httpMetrics struct {
	metrics *dataPoints
}

func (h httpMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key, value := range h.metrics.points {
		w.Write([]byte(fmt.Sprintf("%s %v \n", key, value)))
	}
	for key, value := range h.metrics.averages {
		var sum = float32(0)
		if len(value.points) > 0 {
			for i := range value.points {
				sum = sum + value.points[i].value
			}
			avg := (sum / float32(len(value.points)))
			w.Write([]byte(fmt.Sprintf("%s %v \n", key, avg)))
		}
	}
	for key, value := range h.metrics.counters {
		var sum = float32(0)
		if len(value.points) > 0 {
			for i := range value.points {
				sum = sum + value.points[i].value
			}
			w.Write([]byte(fmt.Sprintf("%s %v \n", key, sum)))
		}
	}
}
