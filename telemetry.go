package telemetry

import (
	"net/http"
	"time"
)

//Telemetry struct
type Telemetry struct {
	httpMetrics httpMetrics
	Counter     *Counter
	Average     *Average
	Total       *Total
}

type metricInterface interface {
	New(string, time.Duration)
	Add(string, float32)
	Get(string) string
	GetAll() map[string]float32
}

type metric struct {
	points   []point
	duration time.Duration
}

type point struct {
	value     float32
	timestamp time.Time
}

//New init metric reporting
func (t *Telemetry) New() error {
	counter := new(Counter)
	counterMetric := make(map[string]metric)
	counter.metric = counterMetric
	t.Counter = counter

	average := new(Average)
	averageMetric := make(map[string]metric)
	average.metric = averageMetric
	t.Average = average

	total := new(Total)
	totalMetric := make(map[string]float32)
	total.metric = totalMetric
	t.Total = total

	t.httpMetrics.metrics = map[string]metricInterface{
		"averages": average,
		"counters": counter,
		"totals":   total,
	}

	go http.ListenAndServe(":9000", t.httpMetrics)
	go t.cullScheduler()
	return nil
}

func (t *Telemetry) cullScheduler() {
	for {
		for key := range t.Average.metric {
			points := cull(t.Average.metric[key].points, t.Average.metric[key].duration)
			t.Average.metric[key] = metric{points, t.Average.metric[key].duration}
		}
		for key := range t.Counter.metric {
			points := cull(t.Counter.metric[key].points, t.Counter.metric[key].duration)
			t.Counter.metric[key] = metric{points, t.Counter.metric[key].duration}
		}
		time.Sleep(time.Second * 5)
	}
}

func cull(points []point, ttl time.Duration) []point {
	var culled []point
	for i := range points {
		if time.Since(points[i].timestamp) < ttl {
			culled = append(culled, points[i])
		}
	}
	return culled
}
