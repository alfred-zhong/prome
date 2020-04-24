package prome

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *Client) AddCounter(opts prometheus.CounterOpts) prometheus.Counter {
	if opts.Name == "" {
		opts.Name = c.ServiceName
	}
	if opts.ConstLabels == nil {
		opts.ConstLabels = make(prometheus.Labels)
	}
	for nn, vv := range c.ConstLabels {
		if _, ok := opts.ConstLabels[nn]; !ok {
			opts.ConstLabels[nn] = vv
		}
	}

	counter := prometheus.NewCounter(opts)
	c.collectors = append(c.collectors, counter)
	return counter
}

func (c *Client) AddCounterVec(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	if opts.Name == "" {
		opts.Name = c.ServiceName
	}
	if opts.ConstLabels == nil {
		opts.ConstLabels = make(prometheus.Labels)
	}
	for nn, vv := range c.ConstLabels {
		if _, ok := opts.ConstLabels[nn]; !ok {
			opts.ConstLabels[nn] = vv
		}
	}

	counterVec := prometheus.NewCounterVec(opts, labelNames)
	c.collectors = append(c.collectors, counterVec)
	return counterVec
}
