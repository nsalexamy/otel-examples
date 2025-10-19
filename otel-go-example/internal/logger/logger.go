package logger

import (
	"context"
	//"log"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

var _logger = logrus.New()

func Init() {
	_logger.SetFormatter(&logrus.JSONFormatter{})
	_logger.SetLevel(logrus.InfoLevel)
}

// Info logs a message with trace_id and span_id from context
func Info(ctx context.Context, msg string, fields ...logrus.Fields) {
	entry := _logger.WithFields(extractTraceFields(ctx))
	if len(fields) > 0 {
		for k, v := range fields[0] {
			entry = entry.WithField(k, v)
		}
	}
	entry.Info(msg)
}

// Error logs an error message with trace context
func Error(ctx context.Context, msg string, err error) {
	_logger.WithFields(extractTraceFields(ctx)).WithError(err).Error(msg)
}

// extractTraceFields pulls trace_id/span_id from context
func extractTraceFields(ctx context.Context) logrus.Fields {
	sc := trace.SpanContextFromContext(ctx)

	//log.Printf("## Extracted SpanContext: %+v", sc)
	if !sc.IsValid() {
		//log.Printf("## Invalid SpanContext")
		return logrus.Fields{}
	}
	return logrus.Fields{
		"trace_id": sc.TraceID().String(),
		"span_id":  sc.SpanID().String(),
	}
}
