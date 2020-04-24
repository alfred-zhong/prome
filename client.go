package prome

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Client represents the client for prometheus server to pull data from.
type Client struct {
	ServiceName string
	Path        string

	collectors    []prometheus.Collector
	defaultLabels prometheus.Labels
}

// ListenAndServe listen on the addr and provide access for prometheus server to
// pull data.
func (c *Client) ListenAndServe(addr string) error {
	reg := prometheus.NewRegistry()

	// Register different collectors.
	reg.MustRegister(c.collectors...)

	http.Handle(
		c.Path,
		decorateHandler(
			promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		),
	)
	return http.ListenAndServe(addr, nil)
}

type decoratedHandler struct {
	h http.Handler
}

func (d *decoratedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for _, f := range constructs {
		f()
	}
	d.h.ServeHTTP(rw, r)
}

func decorateHandler(h http.Handler) http.Handler {
	return &decoratedHandler{h}
}

func init() {
	constructs = append(constructs, updateRuntimeGuage)
}

var constructs []func()

// UseConstruct 给添加构造方法。这些方法会在 prometheus 访问服务接口时并在返回结果前被调用，
// 通常用来更新监控指标。
func UseConstruct(f func()) {
	constructs = append(constructs, f)
}

// NewClient creates and returns a new Client instance.
func NewClient(serviceName string, path string) *Client {
	c := &Client{ServiceName: serviceName, Path: path}
	registerRuntime(c.ServiceName, &c.collectors, c.defaultLabels)

	return c
}
