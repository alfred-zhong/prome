package main

import (
	"time"

	"github.com/alfred-zhong/prome"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	client := prome.NewClient("test", "/foo")
	client.ConstLabels = prometheus.Labels{
		"env": "test",
		"foo": "bar",
	}
	client.EnableRuntime = false
	counter := client.AddCounter(prometheus.CounterOpts{
		Name: "test_counter",
	})
	counterVec := client.AddCounterVec(prometheus.CounterOpts{
		Name: "test_counterVec",
	}, []string{"name"})
	go func() {
		for {
			counter.Inc()
			counterVec.WithLabelValues("jack").Inc()

			time.Sleep(time.Second)
		}
	}()

	if err := client.ListenAndServe(":9000"); err != nil {
		panic(err)
	}
}
