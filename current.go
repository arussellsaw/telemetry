package telemetry

import (
	"sync"
	"time"
)

//Current - a metric containing the value most recently passed to it
type Current struct {
	Name  string
	value float64
	lock  sync.Mutex
}

//NewCurrent - Create a new current metric and add it to the telemetry register
func NewCurrent(tel *Telemetry, name string, _ time.Duration) *Current {
	current := Current{
		Name:  name,
		value: float64(0),
	}
	tel.lock.Lock()
	defer tel.lock.Unlock()
	tel.registry[name] = &current
	return &current
}

//Add - set the value of the Current metric
func (c *Current) Add(tel *Telemetry, value float64) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	tel.lock.Lock()
	defer tel.lock.Unlock()
	tel.registry[c.Name].(*Current).value = value
	return nil
}

//Get - return the value of the metric
func (c *Current) Get(tel *Telemetry) float64 {
	return tel.registry[c.Name].(*Current).value
}

//GetName - return the human readable name of the metric
func (c *Current) GetName() string {
	return c.Name
}

//Maintain - stub method for interface, metric is so simple that it isn't needed
func (c *Current) Maintain() {
	return
}
