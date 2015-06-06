package telemetry

import (
	"time"
)

//Counter - A running total of values provided to it, over a time period
type Counter struct {
	Name     string
	value    float64
	points   map[time.Time]float64
	duration time.Duration
}

//NewCounter - Create new Counter metric with a duration for keeping points
func NewCounter(tel *Telemetry, name string, duration time.Duration) Counter {
	count := Counter{
		Name:     name,
		value:    0,
		points:   make(map[time.Time]float64),
		duration: duration,
	}
	tel.registry[name] = &count
	return count
}

//Add - Add a value to the metric
func (c *Counter) Add(tel *Telemetry, value float64) error {
	tel.registry[c.Name].(*Counter).points[time.Now()] = value
	return nil
}

//Get - Fetch the metric value
func (c *Counter) Get(tel *Telemetry) float64 {
	return tel.registry[c.Name].(*Counter).value
}

//GetName - get the name of the metric
func (c *Counter) GetName() string {
	return c.Name
}

//Maintain - maintain metric value
func (c *Counter) Maintain() {
	points := make(map[time.Time]float64)
	for pointTime, point := range c.points {
		if time.Since(pointTime) < c.duration {
			points[pointTime] = point
		}
	}
	c.points = points
	var count float64
	for _, point := range c.points {
		count = count + point
	}
	c.value = count
}
