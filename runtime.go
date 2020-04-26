package prome

import (
	"fmt"
	"runtime"
	"runtime/pprof"

	"github.com/prometheus/client_golang/prometheus"
)

// Below represents the id of log which should be printed so some runtime
// information or abnornal situations can be captured.
const (
	logIDMemstatsAllocGauge         = "runtime_memstats_alloc"
	logIDMemstatsSysGauge           = "runtime_memstats_sys"
	logIDMemstatsLastGCPauseNSGauge = "runtime_memstats_last_gc"
	logIDRuntimeNumGoroutineGauge   = "runtime_num_goroutine"
	logIDOSThreadsGauge             = "runtime_os_threads"
	logIDRuntimeGOMaxProcsGauge     = "runtime_gomaxprocs"
	logIDRuntimeNumCPUGauge         = "runtime_num_cpu"
)

var (
	memstatsAllocGauge         prometheus.Gauge
	memstatsSysGauge           prometheus.Gauge
	memstatsLastGCPauseNSGauge prometheus.Gauge
	runtimeNumGoroutineGauge   prometheus.Gauge
	osThreadsGauge             prometheus.Gauge
	runtimeGOMaxProcsGauge     prometheus.Gauge
	runtimeNumCPUGauge         prometheus.Gauge
)

func registerRuntime(
	name string,
	collectors *[]prometheus.Collector, defaultLabels prometheus.Labels,
) {
	memstatsAllocGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsAllocGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime memstats alloc (Unit: byte)",
	})
	*collectors = append(*collectors, memstatsAllocGauge)

	memstatsSysGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsSysGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime memstats sys (Unit: byte)",
	})
	*collectors = append(*collectors, memstatsSysGauge)

	memstatsLastGCPauseNSGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsLastGCPauseNSGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime memstats last GC pause (Unit: ns)",
	})
	*collectors = append(*collectors, memstatsLastGCPauseNSGauge)

	runtimeNumGoroutineGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeNumGoroutineGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime number of goroutines",
	})
	*collectors = append(*collectors, runtimeNumGoroutineGauge)

	osThreadsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDOSThreadsGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime number of OS threads",
	})
	*collectors = append(*collectors, osThreadsGauge)

	runtimeGOMaxProcsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeGOMaxProcsGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime GOMAXPROCS",
	})
	*collectors = append(*collectors, runtimeGOMaxProcsGauge)

	runtimeNumCPUGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeNumCPUGauge),
		ConstLabels: defaultLabels,
		Help:        "Runtime number of CPUs",
	})
	*collectors = append(*collectors, runtimeNumCPUGauge)
}

func updateRuntimeGuage() {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	memstatsAllocGauge.Set(float64(memstats.Alloc))
	memstatsSysGauge.Set(float64(memstats.Sys))
	memstatsLastGCPauseNSGauge.Set(float64(memstats.PauseNs[(memstats.NumGC+255)%256]))

	runtimeNumGoroutineGauge.Set(float64(runtime.NumGoroutine()))
	osThreadsGauge.Set(float64(pprof.Lookup("threadcreate").Count()))
	runtimeGOMaxProcsGauge.Set(float64(runtime.GOMAXPROCS(0)))
	runtimeNumCPUGauge.Set(float64(runtime.NumCPU()))
}
