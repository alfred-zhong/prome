package main

import (
	"github.com/alfred-zhong/prome"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	client := prome.NewClient("test", "/foo")
	client.ConstLabels = prometheus.Labels{
		"env": "test",
		"foo": "bar",
	}
	if err := client.ListenAndServe(":9000"); err != nil {
		panic(err)
	}
}
