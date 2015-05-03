package telemetry

import (
	"time"
)

//Counter metric type for total over a time.Duration
type Counter struct {
	metric map[string]metric
}

//New create new counter metric
func (c *Counter) New(name string, duration time.Duration) {
	counter := metric{duration: duration}
	c.metric[name] = counter
}

//Add add value to counter
func (c *Counter) Add(name string, value float32) {
	point := point{value: value, timestamp: time.Now()}
	//this ugly section is because we cannot assign to properties of a
	//struct within a map, so have to create the entire struct again
	counter := cull(c.metric[name].points, c.metric[name].duration)
	points := metric{append(counter, point), c.metric[name].duration}

	c.metric[name] = points
}

//Get get counter value
func (c *Counter) Get(name string) string {
	return ""
}

//GetAll return all counters
func (c *Counter) GetAll() map[string]float32 {
	output := make(map[string]float32)
	for key, value := range c.metric {
		var sum float32
		if len(value.points) > 0 {
			for i := range value.points {
				sum = sum + value.points[i].value
			}
		}
		output[key] = sum
	}
	return output
}
