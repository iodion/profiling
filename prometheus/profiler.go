package metric

// Profiler represents a metric collector
type Profiler struct {
	namespace     string
	counters      map[string]ICounter
	gauges        map[string]IGauge
	histograms    map[string]IHistogram
	histogramVecs map[string]IHistogramVec
}

// NewProfiler creates a new profiler
func NewProfiler(ns string) *Profiler {
	return &Profiler{
		namespace:     ns,
		counters:      map[string]ICounter{},
		gauges:        map[string]IGauge{},
		histograms:    map[string]IHistogram{},
		histogramVecs: map[string]IHistogramVec{},
	}
}

// CounterMetric creates counter metric type
func (p *Profiler) CounterMetric(name string) ICounter {
	p.counters[name] = NewCounter(p.namespace, name)
	return p.counters[name]
}

// CounterMetricWithLabel creates counter metric type
func (p *Profiler) CounterMetricWithLabel(name string, labels map[string]string) ICounter {
	p.counters[name] = NewCounterWithLabel(p.namespace, name, labels)
	return p.counters[name]
}

// GaugeMetric creates gauge metric type
func (p *Profiler) GaugeMetric(name string) IGauge {
	p.gauges[name] = NewGauge(p.namespace, name)
	return p.gauges[name]
}

// GaugeMetricWithLabel creates gauge metric type
func (p *Profiler) GaugeMetricWithLabel(name string, labels map[string]string) IGauge {
	p.gauges[name] = NewGaugeWithLabel(p.namespace, name, labels)
	return p.gauges[name]
}

// HistogramMetric creates histogram metric type
func (p *Profiler) HistogramMetric(name string) IHistogram {
	p.histograms[name] = NewHistogram(p.namespace, name)
	return p.histograms[name]
}

// HistogramMetricWithLabel creates histogram metric type
func (p *Profiler) HistogramMetricWithLabel(name string, labels map[string]string) IHistogram {
	p.histograms[name] = NewHistogramWithLabel(p.namespace, name, labels)
	return p.histograms[name]
}

// HistogramVecWithLabel creates histogramvec object
func (p *Profiler) HistogramVecWithLabel(name string, labels []string, constLabels map[string]string) IHistogramVec {
	p.histogramVecs[name] = NewHistogramVecWithLabel(p.namespace, name, labels, constLabels)
	return p.histogramVecs[name]
}

// GetCounterMetric returns an existing counter metric
func (p *Profiler) GetCounterMetric(name string) (counter ICounter) {
	if v, ok := p.counters[name]; ok {
		counter = v
	}
	return
}

// GetGaugeMetric returns an existing gauge metric
func (p *Profiler) GetGaugeMetric(name string) (gauge IGauge) {
	if v, ok := p.gauges[name]; ok {
		gauge = v
	}
	return
}

// GetHistogramMetric returns an existing histogram metric
func (p *Profiler) GetHistogramMetric(name string) (his IHistogram) {
	if v, ok := p.histograms[name]; ok {
		his = v
	}
	return
}

// GetHistogramVec returns an existing histogramvec object
func (p *Profiler) GetHistogramVec(name string) (hv IHistogramVec) {
	if v, ok := p.histogramVecs[name]; ok {
		hv = v
	}
	return
}
