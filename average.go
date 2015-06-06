package telemetry

import (
	"time"
)

//Average - A running average of values provided to it, over a time period
type Average struct {
	Name     string
	value    float64
	points   map[time.Time]float64
	duration time.Duration
}

//NewAverage - Create new Average metric with a duration for averaging points over
func NewAverage(tel *Telemetry, name string, duration time.Duration) Average {
	avg := Average{
		Name:     name,
		value:    0,
		points:   make(map[time.Time]float64),
		duration: duration,
	}
	tel.registry[name] = &avg
	return avg
}

//Add - Add a value to the metric
func (a *Average) Add(tel *Telemetry, value float64) error {
	tel.registry[a.Name].(*Average).points[time.Now()] = value
	a.Maintain()
	return nil
}

//Get - Fetch the metric, performing the average (mean)
func (a *Average) Get(tel *Telemetry) float64 {
	a.Maintain()
	return tel.registry[a.Name].(*Average).value
}

//GetName - get the name of the metric
func (a *Average) GetName() string {
	return a.Name
}

//Maintain - maintain metric value
func (a *Average) Maintain() {
	points := make(map[time.Time]float64)
	for pointTime, point := range a.points {
		if time.Since(pointTime) < a.duration {
			points[pointTime] = point
		}
	}
	a.points = points
	if len(a.points) > 0 {
		var avg float64
		for _, point := range a.points {
			avg = avg + point
		}
		avg = avg / float64(len(a.points))
		a.value = avg
	} else {
		a.value = 0
	}
}
