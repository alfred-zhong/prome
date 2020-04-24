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
	logIDMemstatsAllocGauge         = "999901"
	logIDMemstatsSysGauge           = "999902"
	logIDMemstatsLastGCPauseNSGauge = "999903"
	logIDRuntimeNumGoroutineGauge   = "999904"
	logIDOSThreadsGauge             = "999905"
	logIDRuntimeGOMaxProcsGauge     = "999906"
	logIDRuntimeNumCPUGauge         = "999907"
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

func attachRuntimeLabel(
	labels prometheus.Labels,
	errorID string, desc string,
) prometheus.Labels {
	if labels == nil {
		labels = make(prometheus.Labels)
	}

	labels["errorId"] = errorID
	labels["desc"] = desc
	return labels
}

func registerRuntime(
	name string,
	collectors *[]prometheus.Collector, defaultLabels prometheus.Labels,
) {
	memstatsAllocGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsAllocGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDMemstatsAllocGauge, "Memstats alloc"),
	})
	*collectors = append(*collectors, memstatsAllocGauge)

	memstatsSysGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsSysGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDMemstatsSysGauge, "Memstats sys"),
	})
	*collectors = append(*collectors, memstatsSysGauge)

	memstatsLastGCPauseNSGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDMemstatsLastGCPauseNSGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDMemstatsLastGCPauseNSGauge, "Memstats GC pause"),
	})
	*collectors = append(*collectors, memstatsLastGCPauseNSGauge)

	runtimeNumGoroutineGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeNumGoroutineGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDRuntimeNumGoroutineGauge, "Mumber of goroutines"),
	})
	*collectors = append(*collectors, runtimeNumGoroutineGauge)

	osThreadsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDOSThreadsGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDOSThreadsGauge, "OS threads"),
	})
	*collectors = append(*collectors, osThreadsGauge)

	runtimeGOMaxProcsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeGOMaxProcsGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDRuntimeGOMaxProcsGauge, "GOMAXPROCS"),
	})
	*collectors = append(*collectors, runtimeGOMaxProcsGauge)

	runtimeNumCPUGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        fmt.Sprintf("%s_%s", name, logIDRuntimeNumCPUGauge),
		ConstLabels: attachRuntimeLabel(defaultLabels, logIDRuntimeNumCPUGauge, "Number of CPUs"),
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
