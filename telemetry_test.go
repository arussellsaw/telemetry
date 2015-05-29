package telemetry_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arussellsaw/telemetry"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	Convey("We should be able to fetch all metrics via the http handler", t, func() {
		expected := `{
  "averages": {},
  "counters": {},
  "currents": {},
  "totals": {}
}`
		handler := telemetry.HTTPMetrics{
			Metrics: map[string]telemetry.MetricInterface{
				"averages": &telemetry.Average{},
				"counters": &telemetry.Counter{},
				"totals":   &telemetry.Total{},
				"currents": &telemetry.Current{},
			},
		}
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "http://127.0.0.1/", nil)
		assert.Nil(t, err)

		handler.ServeHTTP(recorder, req)

		So(recorder.Body.String(), ShouldEqual, expected)
	})
}

//Counters

func TestCounter(t *testing.T) {
	Convey("Test Counters", t, func() {
		Convey("New Counters should exist and be zero value", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Counter.New("test.counter", (60 * time.Second))
			So(tel.Counter.Get("test.counter"), ShouldEqual, "test.counter 0")
		})
		Convey("Counters should equal the total of values added", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Counter.New("test.counter", (60 * time.Second))
			tel.Counter.Add("test.counter", 10)
			So(tel.Counter.Get("test.counter"), ShouldEqual, "test.counter 10")
		})
		Convey("Values in counters should expire after their period has elapsed", func() {
			tel, _ := telemetry.New((1 * time.Second))
			tel.Counter.New("test.counter", (2 * time.Second))
			tel.Counter.Add("test.counter", 10)
			So(tel.Counter.Get("test.counter"), ShouldEqual, "test.counter 10")
			time.Sleep(3 * time.Second)
			So(tel.Counter.Get("test.counter"), ShouldEqual, "test.counter 0")
		})
		Convey("GetAll should return current values for all counters in a map", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Counter.New("test.counter", (60 * time.Second))
			tel.Counter.New("test.counter2", (60 * time.Second))
			tel.Counter.Add("test.counter", 10)
			tel.Counter.Add("test.counter2", 20)
			var expected = map[string]float32{
				"test.counter":  10,
				"test.counter2": 20,
			}
			assert.Equal(t, expected, tel.Counter.GetAll())
		})
	})
}

//Averages

func TestAverage(t *testing.T) {
	Convey("Test Average metrics", t, func() {
		Convey("New Averages should exist and be zero", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Average.New("test.average", (60 * time.Second))
			So(tel.Average.Get("test.average"), ShouldEqual, "test.average 0")
		})
		Convey("Averages should return the mean of all values within the time period", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Average.New("test.average", (60 * time.Second))
			tel.Average.Add("test.average", 10)
			tel.Average.Add("test.average", 20)
			tel.Average.Add("test.average", 30)
			So(tel.Average.Get("test.average"), ShouldEqual, "test.average 20")
		})
		Convey("GetAll should return means for all unique average keys", func() {
			tel, _ := telemetry.New((5 * time.Second))
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
		})
	})
}

//Totals

func TestTotal(t *testing.T) {
	Convey("Test Total metrics", t, func() {
		Convey("New Totals should exist and be zero", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Total.New("test.total", (60 * time.Second))
			So(tel.Total.Get("test.total"), ShouldEqual, "test.total 0")
		})
		Convey("Totals should return the total of all values added in the applications lifetime", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Total.New("test.total", (60 * time.Second))
			tel.Total.Add("test.total", 10)
			tel.Total.Add("test.total", 20)
			tel.Total.Add("test.total", 30)
			So(tel.Total.Get("test.total"), ShouldEqual, "test.total 60")
		})
		Convey("GetAll should return totals of all unique total keys", func() {
			tel, _ := telemetry.New((5 * time.Second))
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
		})
	})
}

//Currents

func TestCurrent(t *testing.T) {
	Convey("Test Current metrics", t, func() {
		Convey("New Currents should exist and be zero", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Current.New("test.current", (60 * time.Second))
			So(tel.Current.Get("test.current"), ShouldEqual, "test.current 0")
		})
		Convey("Currents should return the most recent value added to them", func() {
			tel, _ := telemetry.New((5 * time.Second))
			tel.Current.New("test.current", (60 * time.Second))
			tel.Current.Add("test.current", 10)
			tel.Current.Add("test.current", 20)
			So(tel.Current.Get("test.current"), ShouldEqual, "test.current 20")
		})
		Convey("GetAll should return most recently added value for all unique current keys", func() {
			tel, _ := telemetry.New((5 * time.Second))
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
		})
	})
}
