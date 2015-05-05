package telemetry

import (
	"net/http"
	"time"
)

//Telemetry struct
type Telemetry struct {
	Counter *Counter
	Average *Average
	Total   *Total
	Current *Current
}

//MetricInterface any type of metric
type MetricInterface interface {
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
func New(cullSchedule time.Duration) (*Telemetry, http.Handler) {
	var constructed Telemetry

	counter := &Counter{
		metric: make(map[string]metric),
	}
	constructed.Counter = counter

	average := &Average{
		metric: make(map[string]metric),
	}
	constructed.Average = average

	total := &Total{
		metric: make(map[string]float32),
	}
	constructed.Total = total

	current := &Current{
		metric: make(map[string]float32),
	}
	constructed.Current = current

	handler := HTTPMetrics{
		Metrics: map[string]MetricInterface{
			"averages": average,
			"counters": counter,
			"totals":   total,
			"currents": current,
		},
	}
	go constructed.cullScheduler(cullSchedule)
	return &constructed, &handler
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
