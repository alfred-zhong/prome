package prome

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *Client) AddCounter(opts prometheus.CounterOpts) prometheus.Counter {
	opts = prometheus.CounterOpts(
		formatOpts(prometheus.Opts(opts), c.ConstLabels, c.ServiceName),
	)

	counter := prometheus.NewCounter(opts)
	c.collectors = append(c.collectors, counter)
	return counter
}

func (c *Client) AddCounterVec(
	opts prometheus.CounterOpts, labelNames []string,
) *prometheus.CounterVec {
	opts = prometheus.CounterOpts(
		formatOpts(prometheus.Opts(opts), c.ConstLabels, c.ServiceName),
	)

	counterVec := prometheus.NewCounterVec(opts, labelNames)
	c.collectors = append(c.collectors, counterVec)
	return counterVec
}

func (c *Client) AddGauge(opts prometheus.GaugeOpts) prometheus.Gauge {
	opts = prometheus.GaugeOpts(
		formatOpts(prometheus.Opts(opts), c.ConstLabels, c.ServiceName),
	)

	gauge := prometheus.NewGauge(opts)
	c.collectors = append(c.collectors, gauge)
	return gauge
}

func (c *Client) AddGaugeVec(
	opts prometheus.GaugeOpts, labelNames []string,
) *prometheus.GaugeVec {
	opts = prometheus.GaugeOpts(
		formatOpts(prometheus.Opts(opts), c.ConstLabels, c.ServiceName),
	)

	gaugeVec := prometheus.NewGaugeVec(opts, labelNames)
	c.collectors = append(c.collectors, gaugeVec)
	return gaugeVec
}

func (c *Client) AddHistogram(opts prometheus.HistogramOpts) prometheus.Histogram {
	opts = formatHistogramOpts(opts, c.ConstLabels, c.ServiceName)

	histogram := prometheus.NewHistogram(opts)
	c.collectors = append(c.collectors, histogram)
	return histogram
}

func (c *Client) AddHistogramVec(
	opts prometheus.HistogramOpts, labelNames []string,
) *prometheus.HistogramVec {
	opts = formatHistogramOpts(opts, c.ConstLabels, c.ServiceName)

	histogramVec := prometheus.NewHistogramVec(opts, labelNames)
	c.collectors = append(c.collectors, histogramVec)
	return histogramVec
}

func (c *Client) AddSummary(opts prometheus.SummaryOpts) prometheus.Summary {
	opts = formatSummaryOpts(opts, c.ConstLabels, c.ServiceName)

	summary := prometheus.NewSummary(opts)
	c.collectors = append(c.collectors, summary)
	return summary
}

func (c *Client) AddSummaryVec(
	opts prometheus.SummaryOpts, labelNames []string,
) *prometheus.SummaryVec {
	opts = formatSummaryOpts(opts, c.ConstLabels, c.ServiceName)

	summaryVec := prometheus.NewSummaryVec(opts, labelNames)
	c.collectors = append(c.collectors, summaryVec)
	return summaryVec
}

func formatOpts(
	opts prometheus.Opts,
	defaultLabels prometheus.Labels, defaultName string,
) prometheus.Opts {
	if opts.Name == "" {
		opts.Name = defaultName
	}
	if opts.ConstLabels == nil {
		opts.ConstLabels = make(prometheus.Labels)
	}
	for nn, vv := range defaultLabels {
		if _, ok := opts.ConstLabels[nn]; !ok {
			opts.ConstLabels[nn] = vv
		}
	}
	return opts
}

func formatHistogramOpts(
	opts prometheus.HistogramOpts,
	defaultLabels prometheus.Labels, defaultName string,
) prometheus.HistogramOpts {
	if opts.Name == "" {
		opts.Name = defaultName
	}
	if opts.ConstLabels == nil {
		opts.ConstLabels = make(prometheus.Labels)
	}
	for nn, vv := range defaultLabels {
		if _, ok := opts.ConstLabels[nn]; !ok {
			opts.ConstLabels[nn] = vv
		}
	}
	return opts
}

func formatSummaryOpts(
	opts prometheus.SummaryOpts,
	defaultLabels prometheus.Labels, defaultName string,
) prometheus.SummaryOpts {
	if opts.Name == "" {
		opts.Name = defaultName
	}
	if opts.ConstLabels == nil {
		opts.ConstLabels = make(prometheus.Labels)
	}
	for nn, vv := range defaultLabels {
		if _, ok := opts.ConstLabels[nn]; !ok {
			opts.ConstLabels[nn] = vv
		}
	}
	return opts
}
