package logger

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/propagation"
)

// Global propagator for W3C Trace Context (traceparent, tracestate)
var Propagator = propagation.TraceContext{}

// ExtractContext extracts any incoming OpenTelemetry trace context
// (e.g., from OBI, upstream services, or gateways) from HTTP headers.
func ExtractContext(r *http.Request) context.Context {

	return Propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract context from request headers (from OBI or upstream)
		ctx := ExtractContext(c.Request)

		// Replace the request context so downstream handlers use it
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
