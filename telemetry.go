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
	Current     *Current
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
func New(listen string, cullSchedule time.Duration) *Telemetry {
	var constructed Telemetry
	counter := new(Counter)
	counterMetric := make(map[string]metric)
	counter.metric = counterMetric
	constructed.Counter = counter

	average := new(Average)
	averageMetric := make(map[string]metric)
	average.metric = averageMetric
	constructed.Average = average

	total := new(Total)
	totalMetric := make(map[string]float32)
	total.metric = totalMetric
	constructed.Total = total

	current := new(Current)
	currentMetric := make(map[string]float32)
	current.metric = currentMetric
	constructed.Current = current

	constructed.httpMetrics.metrics = map[string]metricInterface{
		"averages": average,
		"counters": counter,
		"totals":   total,
		"currents": current,
	}

	go http.ListenAndServe(listen, constructed.httpMetrics)
	go constructed.cullScheduler(cullSchedule)
	return &constructed
}

func (t *Telemetry) cullScheduler(schedule time.Duration) {
	for {
		for key := range t.Average.metric {
			points := cull(t.Average.metric[key].points, t.Average.metric[key].duration)
			t.Average.metric[key] = metric{points, t.Average.metric[key].duration}
		}
		for key := range t.Counter.metric {
			points := cull(t.Counter.metric[key].points, t.Counter.metric[key].duration)
			t.Counter.metric[key] = metric{points, t.Counter.metric[key].duration}
		}
		time.Sleep(schedule)
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
