package telemetry

import (
	"fmt"
	"sync"
	"time"
)

//Total a total sum of a metric over the lifetime of the process
type Total struct {
	metric map[string]float32
	lock   sync.Mutex
}

//New new total sum metric
func (t *Total) New(name string, _ time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.metric[name] = 0
}

//Add add value to existing metric
func (t *Total) Add(name string, value float32) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.metric[name] = (t.metric[name] + value)
}

//Get return total
func (t *Total) Get(name string) string {
	return fmt.Sprintf("%s %v", name, t.metric[name])
}

//GetAll get all totals
func (t *Total) GetAll() map[string]float32 {
	output := make(map[string]float32)
	for key, value := range t.metric {
		output[key] = value
	}
	return output
}
