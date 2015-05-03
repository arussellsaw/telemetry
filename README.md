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

    telemetry.CreateAvg("example.avg", (60 * time.Second))
    telemetry.CreateCounter("example.counter", (60 * time.Second))

    telemetry.AppendAvg("example.avg", float32(10))
    telemetry.AppendAvg("example.avg", float32(20))
    telemetry.AppendAvg("example.avg", float32(30))

    telemetry.AppendCounter("example.counter", float32(10))
    telemetry.AppendCounter("example.counter", float32(20))
    telemetry.AppendCounter("example.counter", float32(30))
}

```

to view metrics:

`curl localhost:9000`  


output:  


```
example.avg 20
example.counter 60
```
