package telemetry_test

import (
	"testing"
	"time"

	"github.com/arussellsaw/telemetry"
	. "github.com/smartystreets/goconvey/convey"
)

//Counters

func TestCounter(t *testing.T) {
	Convey("Test Counters", t, func() {
		Convey("New Counters should exist and be zero value", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			counter := telemetry.NewCounter(tel, "test.counter", (60 * time.Second))
			So(counter.Get(tel), ShouldEqual, 0)
		})
		Convey("Counters should equal the total of values added", func() {
			tel := telemetry.New("test", (1 * time.Second))
			counter := telemetry.NewCounter(tel, "test.counter", (60 * time.Second))
			counter.Add(tel, float64(3))
			time.Sleep(1 * time.Second)
			So(counter.Get(tel), ShouldEqual, 3)
		})
		Convey("Values in counters should expire after their period has elapsed", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			counter := telemetry.NewCounter(tel, "test.counter", (200 * time.Millisecond))
			counter.Add(tel, float64(3))
			time.Sleep(100 * time.Millisecond)
			So(counter.Get(tel), ShouldEqual, 3)
			time.Sleep(200 * time.Millisecond)
			So(counter.Get(tel), ShouldEqual, 0)
		})
	})
}

//Averages

func TestAverage(t *testing.T) {
	Convey("Test Averages", t, func() {
		Convey("New Averagess should exist and be zero value", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			avg := telemetry.NewAverage(tel, "test.avg", (60 * time.Second))
			avg.Maintain()
			So(avg.Get(tel), ShouldEqual, 0)
		})
		Convey("Averages should equal the total of values added", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			avg := telemetry.NewAverage(tel, "test.avg", (60 * time.Second))
			avg.Add(tel, float64(3))
			time.Sleep(1 * time.Second)
			So(avg.Get(tel), ShouldEqual, 3)
		})
		Convey("Values in averages should expire after their period has elapsed", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			avg := telemetry.NewAverage(tel, "test.avg", (200 * time.Millisecond))
			avg.Add(tel, float64(3))
			time.Sleep(100 * time.Millisecond)
			So(avg.Get(tel), ShouldEqual, 3)
			time.Sleep(200 * time.Millisecond)
			So(avg.Get(tel), ShouldEqual, 0)
		})
	})
}

//Currents

func TestCurrent(t *testing.T) {
	Convey("Test Currents", t, func() {
		Convey("New Currents should exist and be zero value", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			cur := telemetry.NewCurrent(tel, "test.cur", (60 * time.Second))
			So(cur.Get(tel), ShouldEqual, 0)
		})
		Convey("Current's value should be the latest one passed to it", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			cur := telemetry.NewCurrent(tel, "test.cur", (60 * time.Second))
			cur.Add(tel, float64(10))
			cur.Add(tel, float64(3))
			So(cur.Get(tel), ShouldEqual, 3)
		})
	})
}

//Totals

func TestTotal(t *testing.T) {
	Convey("Test Totals", t, func() {
		Convey("New Totals should exist and be zero value", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			tot := telemetry.NewTotal(tel, "test.total", (60 * time.Second))
			So(tot.Get(tel), ShouldEqual, 0)
		})
		Convey("Total's value should be the sum of all values passed to it", func() {
			tel := telemetry.New("test", (100 * time.Millisecond))
			tot := telemetry.NewTotal(tel, "test.total", (60 * time.Second))
			tot.Add(tel, float64(10))
			tot.Add(tel, float64(3))
			So(tot.Get(tel), ShouldEqual, 13)
		})
	})
}
