package telemetry

import (
	"fmt"
	"time"
)

//Average running average over a time.Duration
type Average struct {
	metric map[string]metric
}

//New create new averaged metric
func (a *Average) New(name string, duration time.Duration) {
	average := metric{duration: duration}
	a.metric[name] = average
}

//Add add value to averaged metric
func (a *Average) Add(name string, value float32) {
	point := point{value: value, timestamp: time.Now()}
	//this ugly section is because we cannot assign to properties of a
	//struct within a map, so have to create the entire struct again
	averages := cull(a.metric[name].points, a.metric[name].duration)
	points := metric{append(averages, point), a.metric[name].duration}

	a.metric[name] = points
}

//Get return average metric
func (a *Average) Get(name string) string {
	return ""
}

//GetAll return all average metrics
func (a *Average) GetAll() string {
	var output string
	for key, value := range a.metric {
		var avg float32
		if len(value.points) > 0 {
			for i := range value.points {
				avg = avg + value.points[i].value
			}
			avg = avg / float32(len(value.points))
		}
		output += fmt.Sprintf("%s %v \n", key, avg)
	}
	return output
}
