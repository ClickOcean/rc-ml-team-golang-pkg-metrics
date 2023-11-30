package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type incomingRequestCounter struct {
	appName string
	*prometheus.CounterVec
}

func NewIncomingRequestCounter(appName string) incomingRequestCounter { //nolint: revive
	counterVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rc_ml_http_requests_total",
			Help: "Total number of incoming requests",
		},
		[]string{"statusCode", "method", "path", "appName"},
	)

	prometheus.MustRegister(
		counterVec,
	)

	return incomingRequestCounter{
		appName:    appName,
		CounterVec: counterVec,
	}
}

func (c incomingRequestCounter) ObserveIncomingRequests(
	statusCode int,
	method, path string,
) {
	c.WithLabelValues(strconv.Itoa(statusCode), method, path, c.appName).Inc()
}
