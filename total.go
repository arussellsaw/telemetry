package telemetry

import (
	"sync"
	"time"
)

//Total a total sum of a metric over the lifetime of the process
type Total struct {
	Name  string
	value float64
	lock  sync.Mutex
}

//NewTotal - create new total metric type, add it to telemetry register
func NewTotal(tel *Telemetry, name string, duration time.Duration) *Total {
	total := Total{
		Name:  name,
		value: 0,
	}
	tel.lock.Lock()
	defer tel.lock.Unlock()
	tel.registry[name] = &total
	return &total
}

//Add - add value to total
func (t *Total) Add(tel *Telemetry, value float64) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	tel.lock.Lock()
	defer tel.lock.Unlock()
	tel.registry[t.Name].(*Total).value += value
	return nil
}

//Get - get current value
func (t *Total) Get(tel *Telemetry) float64 {
	return tel.registry[t.Name].(*Total).value
}

//GetName - get metric name
func (t *Total) GetName() string {
	return t.Name
}

//Maintain - stub method for interface, does nothing
func (t *Total) Maintain() {
	return
}
