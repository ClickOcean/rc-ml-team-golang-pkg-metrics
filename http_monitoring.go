package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type label string

const (
	statusCodeLabel label = "statusCode"
	methodLabel     label = "method"
	pathLabel       label = "path"
	appNameLabel    label = "appName"
)

type HTTPMonitoring struct {
	*incomingRequestCounter
	*requestDurationHistogram
}

func NewHTTPMonitoring(appName string) *HTTPMonitoring {
	return &HTTPMonitoring{
		newIncomingRequestCounter(appName),
		newRequestDurationHistogram(appName),
	}
}

type incomingRequestCounter struct {
	appName string
	counter *prometheus.CounterVec
}

func newIncomingRequestCounter(appName string) *incomingRequestCounter {
	counterVec := NewCounterVec(
		"rc_ml_http_requests_total",
		"Total number of incoming requests",
		[]string{string(statusCodeLabel), string(methodLabel), string(pathLabel), string(appNameLabel)},
	)

	return &incomingRequestCounter{
		appName: appName,
		counter: counterVec,
	}
}

func (c incomingRequestCounter) ObserveIncomingRequests(
	statusCode int,
	method, path string,
) {
	c.counter.WithLabelValues(strconv.Itoa(statusCode), method, path, c.appName).Inc()
}

type requestDurationHistogram struct {
	appName string
	hist    *prometheus.HistogramVec
}

func newRequestDurationHistogram(appName string) *requestDurationHistogram {
	histogramVec := NewHistogramVec(
		"rc_ml_http_requests_duration",
		"HTTP requests duration histogram",
		[]string{string(statusCodeLabel), string(methodLabel), string(pathLabel), string(appNameLabel)},
	)

	return &requestDurationHistogram{
		appName: appName,
		hist:    histogramVec,
	}
}

func (h requestDurationHistogram) ObserveRequestsDuration(
	duration float64,
	statusCode int,
	method, path string,
) {
	h.hist.WithLabelValues(strconv.Itoa(statusCode), method, path, h.appName).Observe(duration)
}
