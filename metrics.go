package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func NewCounterVec(name, desc string, labels []string) *prometheus.CounterVec {
	counterVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: desc,
		},
		labels,
	)

	prometheus.MustRegister(
		counterVec,
	)

	return counterVec
}

func NewHistogramVec(name, desc string, labels []string) *prometheus.HistogramVec {
	histogramVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name,
			Help: desc,
		},
		labels,
	)

	prometheus.MustRegister(
		histogramVec,
	)

	return histogramVec
}
