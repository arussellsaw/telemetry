package telemetry

import (
	"fmt"
	"sync"
	"time"
)

//Current current value metric
type Current struct {
	metric map[string]float32
	lock   sync.Mutex
}

//New new current metric
func (c *Current) New(name string, _ time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.metric[name] = 0
}

//Add add value to existing metric
func (c *Current) Add(name string, value float32) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.metric[name] = value
}

//Get return total
func (c *Current) Get(name string) string {
	return fmt.Sprintf("%s %v", name, c.metric[name])
}

//GetAll get all totals
func (c *Current) GetAll() map[string]float32 {
	output := make(map[string]float32)
	for key, value := range c.metric {
		output[key] = value
	}
	return output
}
