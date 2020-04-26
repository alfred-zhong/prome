package prome

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// MiddlewareRequestCount returns a gin HandlerFunc which can be used as
// middleware to capture request count.
func (c *Client) MiddlewareRequestCount(metricsName string) gin.HandlerFunc {
	if metricsName == "" {
		metricsName = fmt.Sprintf("%s_request_count", c.ServiceName)
	}

	cv := c.AddCounterVec(prometheus.CounterOpts{
		Name: metricsName,
		Help: "Request count",
	}, []string{"method", "path"})

	return func(c *gin.Context) {
		cv.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()

		c.Next()
	}
}

var defaultDurationSummaryObjectives = map[float64]float64{
	0.5:  0.05,
	0.9:  0.01,
	0.95: 0.005,
	0.99: 0.001,
}

// MiddlewareRequestDuration returns a gin handler which can be used as
// middleware to capture request duration summary.
func (c *Client) MiddlewareRequestDuration(
	metricsName string, objectives map[float64]float64,
) gin.HandlerFunc {
	if metricsName == "" {
		metricsName = fmt.Sprintf("%s_request_duration", c.ServiceName)
	}
	if objectives == nil {
		objectives = defaultDurationSummaryObjectives
	}

	sv := c.AddSummaryVec(prometheus.SummaryOpts{
		Name:       metricsName,
		Objectives: objectives,
		Help:       "Request duration (Unit: ns)",
	}, []string{"method", "path"})

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		sv.WithLabelValues(
			c.Request.Method, c.Request.URL.Path,
		).Observe(float64(time.Since(startTime).Nanoseconds()))
	}
}
