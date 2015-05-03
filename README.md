# telemetry
metric reporting for Go applications

sample usage:

```
package main

import(
    "github.com/arussellsaw/telemetry"
    "time"
    )

func main() {
    var telemetry = new(telemetry.Telemetry)
    telemetry.Initialize()

    /*
    The time.Duration() parameter is used to cull metrics older than the duration
    this provides the ability to provide x per-minute stats, cull is run on append
    methods, also a scheduled cull is run every 5s (configureable in future)
    */

    telemetry.Average.New("example.avg", (60 * time.Second))
    telemetry.Counter.New("example.counter", (60 * time.Second))
    telemetry.Total.New("example.total", 0 * time.Second) //duration parameter is useless, but is needed to conform to interface


    telemetry.Average.Add("example.avg", float32(10))
    telemetry.Average.Add("example.avg", float32(20))
    telemetry.Average.Add("example.avg", float32(30))

    telemetry.Counter.Add("example.counter", float32(10))
    telemetry.Counter.Add("example.counter", float32(20))
    telemetry.Counter.Add("example.counter", float32(30))

    telemetry.Total.Add("example.total", float32(10))
    telemetry.Total.Add("example.total", float32(20))
    telemetry.Total.Add("example.total", float32(30))
}

```

to view metrics:

`curl localhost:9000`  


output:  


```
{
    "averages" {
        example.avg 20
    },
    "counters" {
        example.counter 60
    },
    "totals" {
        example.total 60
    }
}
```

the same command 61s later

```
{
    "averages" {
        example.avg 0
    },
    "counters" {
        example.counter 0
    },
    "totals" {
        example.total 60
    }
}
```
