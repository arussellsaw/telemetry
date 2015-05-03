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
