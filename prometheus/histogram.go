package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// IHistogram is a contract
type IHistogram interface {
	Observe(float64)
}

// Histogram represents a Histogram metric type
type Histogram struct {
	Histogram prometheus.Histogram
}

// NewHistogram creates an histogram metric without label
func NewHistogram(ns, name string) IHistogram {
	h := promauto.NewHistogram(prometheus.HistogramOpts{Namespace: ns, Name: name})
	return &Histogram{h}
}

// NewHistogramWithLabel creates an histogram metric with label
func NewHistogramWithLabel(ns, name string, labels map[string]string) IHistogram {
	h := promauto.NewHistogram(prometheus.HistogramOpts{Namespace: ns, Name: name, ConstLabels: labels})
	return &Histogram{h}
}

// Observe sets an observation point
func (h *Histogram) Observe(v float64) {
	h.Histogram.Observe(v)
}

// IHistogramVec represents a contract
type IHistogramVec interface {
	Observe(map[string]string, float64)
}

// HistogramVec is a wrapper
type HistogramVec struct {
	HistogramVec *prometheus.HistogramVec
}

// NewHistogramVecWithLabel constructs a HistogramVec
func NewHistogramVecWithLabel(ns, name string, labels []string, constLabels map[string]string) IHistogramVec {
	c := prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: ns, Name: name, ConstLabels: constLabels}, labels)
	prometheus.MustRegister(c)
	return &HistogramVec{c}
}

// Observe sets an observation instance
func (hv HistogramVec) Observe(labels map[string]string, value float64) {
	hv.HistogramVec.With(prometheus.Labels(labels)).Observe(value)
}
