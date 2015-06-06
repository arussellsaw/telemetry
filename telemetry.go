package telemetry

import (
	"time"
)

//Telemetry container struct
type Telemetry struct {
	registry map[string]metricInterface
	duration time.Duration
}

type metricInterface interface {
	GetName() string
	Add(*Telemetry, float64) error
	Get(*Telemetry) float64
	Maintain()
}

//New create new telemetry struct
func New(duration time.Duration) *Telemetry {
	tel := Telemetry{
		registry: make(map[string]metricInterface),
		duration: duration,
	}
	go tel.maintainance()
	return &tel
}

//GetAll get all metrics registered to Telemetry
func (t *Telemetry) GetAll() map[string]float64 {
	metrics := make(map[string]float64)
	for _, metric := range t.registry {
		metrics[metric.GetName()] = metric.Get(t)
	}
	return metrics
}

//maintainance run management on existing metrics
func (t *Telemetry) maintainance() {
	for {
		for _, metric := range t.registry {
			metric.Maintain()
		}
		time.Sleep(t.duration)
	}
}
