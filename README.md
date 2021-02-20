# prome [Deprecated]

![Go](https://github.com/alfred-zhong/prome/workflows/Go/badge.svg?branch=master) [![GoDoc](https://godoc.org/github.com/alfred-zhong/prome?status.svg)](https://pkg.go.dev/github.com/alfred-zhong/prome) [![Go Report Card](https://goreportcard.com/badge/github.com/alfred-zhong/prome)](https://goreportcard.com/report/github.com/alfred-zhong/prome)

**Deprecated. This repository has been merged into [goutil](https://github.com/alfred-zhong/goutil).**

Prometheus Client for easy use.

## Why

I wrote this library so i can easily use it in my Go projects.

## Features

* Easy easy.
* Provide basic Go runtime metrics.
* Support gin middleware for providing API request counts and duration metrics.
* Easy for developer to add custom metrics.

## Usage

The simplest usage:

```go
func main() {
	client := prome.NewClient("test", "/foo")
	if err := client.ListenAndServe(":9000"); err != nil {
		panic(err)
	}
}
```

Just include the code above already provides you the Go runtime metrics info. You can run `curl "http://127.0.0.1:9000/foo"` and get the prometheus data.

```
11:57 tmp ➜  curl "http://127.0.0.1:9000/foo"
# HELP test_runtime_gomaxprocs Runtime GOMAXPROCS
# TYPE test_runtime_gomaxprocs gauge
test_runtime_gomaxprocs 4
# HELP test_runtime_memstats_alloc Runtime memstats alloc (Unit: byte)
# TYPE test_runtime_memstats_alloc gauge
test_runtime_memstats_alloc 1.346416e+06
# HELP test_runtime_memstats_last_gc Runtime memstats last GC pause (Unit: ns)
# TYPE test_runtime_memstats_last_gc gauge
test_runtime_memstats_last_gc 0
# HELP test_runtime_memstats_sys Runtime memstats sys (Unit: byte)
# TYPE test_runtime_memstats_sys gauge
test_runtime_memstats_sys 7.2827136e+07
# HELP test_runtime_num_cpu Runtime number of CPUs
# TYPE test_runtime_num_cpu gauge
test_runtime_num_cpu 4
# HELP test_runtime_num_goroutine Runtime number of goroutines
# TYPE test_runtime_num_goroutine gauge
test_runtime_num_goroutine 3
# HELP test_runtime_os_threads Runtime number of OS threads
# TYPE test_runtime_os_threads gauge
test_runtime_os_threads 7
```

## Examples

You can get more examples in the `examples` directory to find more advenced usage.
