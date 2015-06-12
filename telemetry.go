package telemetry

import (
	"fmt"
	"sync"
	"time"
)

//Telemetry container struct
type Telemetry struct {
	registry map[string]Metric
	prefix   string
	lock     sync.Mutex
	duration time.Duration
}

type Metric interface {
	GetName() string
	Add(*Telemetry, float64) error
	Get(*Telemetry) float64
	Maintain()
}

//New create new telemetry struct
func New(prefix string, duration time.Duration) *Telemetry {
	tel := Telemetry{
		registry: make(map[string]Metric),
		prefix:   prefix,
		duration: duration,
	}
	go tel.maintainance()
	return &tel
}

//GetAll get all metrics registered to Telemetry
func (t *Telemetry) GetAll() map[string]float64 {
	metrics := make(map[string]float64)
	for _, metric := range t.registry {
		metrics[fmt.Sprintf("%s%s", t.prefix, metric.GetName())] = metric.Get(t)
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
