package builtin

import (
	"runtime"
	"runtime/debug"
	"time"

	"github.com/arussellsaw/telemetry"
)

//Runtime builtin runtime telemetry
type Runtime struct {
	Telemetry *telemetry.Telemetry
	Prefix    string
}

//Record start collecting stats about the Go runtime
func (r *Runtime) Record() {
	duration := (60 * time.Second)
	r.Telemetry.Average.New(r.Prefix+".runtime.mem.alloc", duration)
	r.Telemetry.Average.New(r.Prefix+".runtime.mem.heap.alloc", duration)
	r.Telemetry.Average.New(r.Prefix+".runtime.mem.heap.used", duration)
	r.Telemetry.Current.New(r.Prefix+".runtime.gc.count", duration)
	go r.Poll()
}

//Poll update runtime metrics
func (r *Runtime) Poll() {
	var mem runtime.MemStats
	var gc debug.GCStats
	for {
		runtime.ReadMemStats(&mem)
		debug.ReadGCStats(&gc)
		r.Telemetry.Average.Add(r.Prefix+".runtime.mem.alloc", float32(mem.Alloc))
		r.Telemetry.Average.Add(r.Prefix+".runtime.mem.heap.alloc", float32(mem.HeapAlloc))
		r.Telemetry.Average.Add(r.Prefix+".runtime.mem.heap.used", float32(mem.HeapInuse))
		r.Telemetry.Current.Add(r.Prefix+".runtime.gc.count", float32(gc.NumGC))
		time.Sleep((30 * time.Second))
	}
}
