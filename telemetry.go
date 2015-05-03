package telemetry

import (
	"net/http"
	"time"
)

//Telemetry struct
type Telemetry struct {
	httpMetrics httpMetrics
	metrics     *dataPoints
}

type dataPoints struct {
	points   map[string]float32
	averages map[string]avgMetric
	counters map[string]counterMetric
}

type point struct {
	value     float32
	timestamp time.Time
}

type counterMetric struct {
	points   []point
	duration time.Duration
}

type avgMetric struct {
	points   []point
	duration time.Duration
}

//Initialize metric reporting
func (t *Telemetry) Initialize() error {
	points := make(map[string]float32)
	averages := make(map[string]avgMetric)
	counters := make(map[string]counterMetric)
	var data = dataPoints{points, averages, counters}
	t.httpMetrics.metrics = &data
	t.metrics = &data
	go http.ListenAndServe(":9000", t.httpMetrics)
	go t.cullScheduler()
	return nil
}

//Set set telemetry value
func (t *Telemetry) Set(name string, value float32) {
	t.metrics.points[name] = value
}

//IncrementMetric add value to existing metric
func (t *Telemetry) IncrementMetric(name string, value float32) error {
	t.metrics.points[name] = (t.metrics.points[name] + value)
	return nil
}

//CreateAvg create new averaged metric
func (t *Telemetry) CreateAvg(name string, duration time.Duration) {
	average := avgMetric{duration: duration}
	t.metrics.averages[name] = average
}

//CreateCounter create new counter metric
func (t *Telemetry) CreateCounter(name string, duration time.Duration) {
	counter := counterMetric{duration: duration}
	t.metrics.counters[name] = counter
}

//AppendCounter add value to counter
func (t *Telemetry) AppendCounter(name string, value float32) {
	point := point{value: value, timestamp: time.Now()}
	//this ugly section is because we cannot assign to properties of a
	//struct within a map, so have to create the entire struct again
	counter := t.cull(t.metrics.counters[name].points, t.metrics.counters[name].duration)
	points := counterMetric{append(counter, point), t.metrics.counters[name].duration}

	t.metrics.counters[name] = points
}

//AppendAvg add value to averaged metric
func (t *Telemetry) AppendAvg(name string, value float32) {
	point := point{value: value, timestamp: time.Now()}
	//this ugly section is because we cannot assign to properties of a
	//struct within a map, so have to create the entire struct again
	averages := t.cull(t.metrics.averages[name].points, t.metrics.averages[name].duration)
	points := avgMetric{append(averages, point), t.metrics.averages[name].duration}

	t.metrics.averages[name] = points
}

func (t *Telemetry) cullScheduler() {
	for {
		for key := range t.metrics.averages {
			points := t.cull(t.metrics.averages[key].points, t.metrics.averages[key].duration)
			t.metrics.averages[key] = avgMetric{points, t.metrics.averages[key].duration}
		}
		for key := range t.metrics.counters {
			points := t.cull(t.metrics.counters[key].points, t.metrics.counters[key].duration)
			t.metrics.counters[key] = counterMetric{points, t.metrics.counters[key].duration}
		}
		time.Sleep(time.Second * 5)
	}
}

func (t *Telemetry) cull(points []point, ttl time.Duration) []point {
	var culled []point
	for i := range points {
		if time.Since(points[i].timestamp) < ttl {
			culled = append(culled, points[i])
		}
	}
	return culled
}
