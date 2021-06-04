package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// IGauge is a gauge contract
type IGauge interface {
	Set(v float64)
}

// Gauge represents a gauge metric type
type Gauge struct {
	Gauge prometheus.Gauge
}

// NewGauge is a gauge constructor
func NewGauge(ns, name string) IGauge {
	g := promauto.NewGauge(prometheus.GaugeOpts{Namespace: ns, Name: name})
	return &Gauge{g}
}

// NewGaugeWithLabel is a gauge constructor with label
func NewGaugeWithLabel(ns, name string, labels map[string]string) IGauge {
	g := promauto.NewGauge(prometheus.GaugeOpts{Namespace: ns, Name: name, ConstLabels: labels})
	return &Gauge{g}
}

// Set sets a gauge value
func (g *Gauge) Set(v float64) {
	g.Gauge.Set(v)
}
