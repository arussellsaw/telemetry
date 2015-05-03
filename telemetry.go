package telemetry

import (
	"errors"
	"net/http"
)

//Telemetry struct
type Telemetry struct {
	httpMetrics httpMetrics
	metrics     *dataPoints
}

type dataPoints struct {
	intMetrics   map[string]int64
	floatMetrics map[string]float32
}

//Initialize metric reporting
func (t *Telemetry) Initialize() error {
	ints := make(map[string]int64)
	floats := make(map[string]float32)
	var data = dataPoints{ints, floats}
	t.httpMetrics.metrics = &data
	t.metrics = &data
	go http.ListenAndServe(":9000", t.httpMetrics)
	return nil
}

//SetFloat set float telemetry value
func (t *Telemetry) SetFloat(name string, value float32) {
	t.metrics.floatMetrics[name] = value
}

//SetInt set int telemetry value
func (t *Telemetry) SetInt(name string, value int64) {
	t.metrics.intMetrics[name] = value
}

//IncrementMetric add value to existing metric
func (t *Telemetry) IncrementMetric(name string, value interface{}) error {
	switch value.(type) {
	default:
		return errors.New("unexpected type")
	case float32:
		t.metrics.floatMetrics[name] = (t.metrics.floatMetrics[name] + value.(float32))
	case int64:
		t.metrics.intMetrics[name] = (t.metrics.intMetrics[name] + value.(int64))
	}
	return nil
}
