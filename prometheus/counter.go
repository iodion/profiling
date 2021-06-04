package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// ICounter is a counter contract. Describes a counter behavior
type ICounter interface {
	Inc()
}

// Counter is a counter metric wrapper
type Counter struct {
	Counter prometheus.Counter
}

// NewCounter is a counter constructor
func NewCounter(ns, name string) ICounter {
	c := promauto.NewCounter(prometheus.CounterOpts{Namespace: ns, Name: name})
	return &Counter{c}
}

// NewCounterWithLabel is a counter constructor with label
func NewCounterWithLabel(ns, name string, labels map[string]string) ICounter {
	c := promauto.NewCounter(prometheus.CounterOpts{Namespace: ns, Name: name, ConstLabels: labels})
	return &Counter{c}
}

// Inc increments a counter
func (c *Counter) Inc() {
	c.Counter.Inc()
}
