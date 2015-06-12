# telemetry

[![build-status](https://travis-ci.org/arussellsaw/telemetry.svg?branch=master)](https://travis-ci.org/arussellsaw/telemetry) [![code-coverage](http://gocover.io/_badge/github.com/arussellsaw/telemetry)](http://gocover.io/github.com/arussellsaw/telemetry)
[![go-doc](https://godoc.org/github.com/arussellsaw/telemetry?status.svg)](https://godoc.org/github.com/arussellsaw/telemetry)

metric reporting for Go applications

sample usage:

```go
package main

import(
    "github.com/arussellsaw/telemetry"
    "github.com/arussellsaw/telemetry/reporters"
    "time"
    "net/http"
    )

func main() {
    //New telemetry object (prefix, maintainance schedule)
    tel := telemetry.New("test", 5 * time.Second)
    avg := telemetry.NewAverage(tel, "average", 60 * time.Second)

    //Register influxdb reporter
    influx := reporters.InfluxReporter{
        Host: "192.168.1.100:8086",
        Interval: 60 * time.Second,
        Tel: tel,
        Database: "telemetry"
    }
    influx.Report() //trigger reporting loop

    //Create http handler for json metrics
    telemetryHandler := reporters.TelemetryHandler{
        Tel: tel,
    }
    http.HandleFunc("/metrics", telemetryHandler.ServeHTTP)
    http.ListenAndServe(":8080", nil)

    start = time.Now()
    somethingYouWantToTime()
    avg.Add(tel, float64(time.Since(start).Nanoseconds()))
}

```
