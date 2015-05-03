package telemetry

import (
	"fmt"
	"time"
)

//Total a total sum of a metric over the lifetime of the process
type Total struct {
	metric map[string]float32
}

//New new total sum metric
func (t *Total) New(name string, _ time.Duration) {
	t.metric[name] = 0
}

//Add add value to existing metric
func (t *Total) Add(name string, value float32) {
	t.metric[name] = (t.metric[name] + value)
}

//Get return total
func (t *Total) Get(name string) string {
	return ""
}

//GetAll get all totals
func (t *Total) GetAll() string {
	var output string
	for key, value := range t.metric {
		output += fmt.Sprintf("%s %v \n", key, value)
	}
	return output
}
