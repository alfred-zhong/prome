package main

import (
	"math/rand"
	"time"

	"github.com/alfred-zhong/prome"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	client := prome.NewClient("test", "/foo")
	client.ConstLabels = prometheus.Labels{
		"env": "test",
	}
	client.EnableRuntime = false

	counter := client.AddCounter(prometheus.CounterOpts{
		Name: "test_counter",
	})
	counterVec := client.AddCounterVec(prometheus.CounterOpts{
		Name: "test_counterVec",
	}, []string{"name"})

	gauge := client.AddGauge(prometheus.GaugeOpts{
		Name: "test_gauge",
	})
	gaugeVec := client.AddGaugeVec(prometheus.GaugeOpts{
		Name: "test_gaugeVec",
	}, []string{"name"})

	histogram := client.AddHistogram(prometheus.HistogramOpts{
		Name:    "test_histogram",
		Buckets: []float64{10, 20, 30, 40, 50},
	})
	histogramVec := client.AddHistogramVec(prometheus.HistogramOpts{
		Name:    "test_histogramVec",
		Buckets: []float64{10, 20, 30, 40, 50},
	}, []string{"name"})

	summary := client.AddSummary(prometheus.SummaryOpts{
		Name:       "test_summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
	summaryVec := client.AddSummaryVec(prometheus.SummaryOpts{
		Name:       "test_summaryVec",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"name"})

	go func() {
		for {
			counter.Inc()
			counterVec.WithLabelValues("hysteria").Inc()

			gauge.Inc()
			gaugeVec.WithLabelValues("roxanne").Inc()

			histogram.Observe(rand.Float64() * 100)
			histogramVec.WithLabelValues("cassandra").Observe(rand.Float64() * 100)

			summary.Observe(rand.Float64() * 100)
			summaryVec.WithLabelValues("hysteria").Observe(rand.Float64() * 30)
			summaryVec.WithLabelValues("roxanne").Observe(rand.Float64() * 50)
			summaryVec.WithLabelValues("riful").Observe(rand.Float64() * 100)

			time.Sleep(time.Second)
		}
	}()

	if err := client.ListenAndServe(":9000"); err != nil {
		panic(err)
	}
}
