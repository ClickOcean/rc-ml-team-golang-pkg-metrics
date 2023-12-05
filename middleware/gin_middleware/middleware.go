package ginmiddleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Metrics interface {
	ObserveIncomingRequests(statusCode int, method, path string)
	ObserveRequestsDuration(duration float64, statusCode int, method, path string)
}

func MetricsMiddleware(metrics Metrics) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		start := time.Now()
		ctx.Next()
		duration := start.Sub(start).Milliseconds()

		status, method, path := ctx.Writer.Status(), ctx.Request.Method, ctx.FullPath()

		metrics.ObserveIncomingRequests(status, method, path)
		metrics.ObserveRequestsDuration(float64(duration), status, method, path)
	}
}
