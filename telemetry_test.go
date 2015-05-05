package telemetry_test

import (
	"testing"
	"time"

	"github.com/arussellsaw/telemetry"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = telemetry.New(":9000", (5 * time.Second))
}

//Counters

func TestCounterNew(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Counter.New("test.counter", (60 * time.Second))
	assert.Equal(t, "test.counter 0", tel.Counter.Get("test.counter"))
}

func TestCounterAdd(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Counter.New("test.counter", (60 * time.Second))
	tel.Counter.Add("test.counter", 10)
	assert.Equal(t, "test.counter 10", tel.Counter.Get("test.counter"))
}

func TestCounterGetAll(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Counter.New("test.counter", (60 * time.Second))
	tel.Counter.New("test.counter2", (60 * time.Second))
	tel.Counter.Add("test.counter", 10)
	tel.Counter.Add("test.counter2", 20)
	var expected = map[string]float32{
		"test.counter":  10,
		"test.counter2": 20,
	}
	assert.Equal(t, expected, tel.Counter.GetAll())
}

//Averages

func TestAverageNew(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Average.New("test.average", (60 * time.Second))
	assert.Equal(t, "test.average 0", tel.Average.Get("test.average"))
}

func TestAverageAdd(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Average.New("test.average", (60 * time.Second))
	tel.Average.Add("test.average", 10)
	tel.Average.Add("test.average", 20)
	tel.Average.Add("test.average", 30)
	assert.Equal(t, "test.average 20", tel.Average.Get("test.average"))
}

func TestAverageGetAll(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Average.New("test.average", (60 * time.Second))
	tel.Average.Add("test.average", 10)
	tel.Average.Add("test.average", 20)

	tel.Average.New("test.average2", (60 * time.Second))
	tel.Average.Add("test.average2", 20)
	tel.Average.Add("test.average2", 30)

	var expected = map[string]float32{
		"test.average":  15,
		"test.average2": 25,
	}
	assert.Equal(t, expected, tel.Average.GetAll())
}

//Totals

func TestTotalNew(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Total.New("test.total", (60 * time.Second))
	assert.Equal(t, "test.total 0", tel.Total.Get("test.total"))
}

func TestTotalAdd(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Total.New("test.total", (60 * time.Second))
	tel.Total.Add("test.total", 10)
	tel.Total.Add("test.total", 20)
	tel.Total.Add("test.total", 30)
	assert.Equal(t, "test.total 60", tel.Total.Get("test.total"))
}

func TestTotalGetAll(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Total.New("test.total", (60 * time.Second))
	tel.Total.Add("test.total", 10)
	tel.Total.Add("test.total", 20)

	tel.Total.New("test.total2", (60 * time.Second))
	tel.Total.Add("test.total2", 20)
	tel.Total.Add("test.total2", 30)

	var expected = map[string]float32{
		"test.total":  30,
		"test.total2": 50,
	}
	assert.Equal(t, expected, tel.Total.GetAll())
}

//Currents

func TestCurrentlNew(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Current.New("test.current", (60 * time.Second))
	assert.Equal(t, "test.current 0", tel.Current.Get("test.current"))
}

func TestCurrentAdd(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Current.New("test.current", (60 * time.Second))
	tel.Current.Add("test.current", 10)
	tel.Current.Add("test.current", 20)
	assert.Equal(t, "test.current 20", tel.Current.Get("test.current"))
}

func TestCurrentGetAll(t *testing.T) {
	tel := telemetry.New(":9000", (5 * time.Second))
	tel.Current.New("test.current", (60 * time.Second))
	tel.Current.Add("test.current", 10)
	tel.Current.Add("test.current", 20)

	tel.Current.New("test.current2", (60 * time.Second))
	tel.Current.Add("test.current2", 20)
	tel.Current.Add("test.current2", 10)

	var expected = map[string]float32{
		"test.current":  20,
		"test.current2": 10,
	}
	assert.Equal(t, expected, tel.Current.GetAll())
}
