package example

import (
	"github.com/gin-gonic/gin"
	"github.com/nsalexamy/otel-examples/otel-go-example/internal/logger"
)

func Logging(c *gin.Context) {
	// get log level from query parameter, default to "info"
	level := c.DefaultQuery("level", "info")

	switch level {
	case "info":
		logger.Info(c.Request.Context(), "This is an info message")
	case "error":
		logger.Error(c.Request.Context(), "This is an error message", nil)
	default:
		level = "info"
		logger.Info(c.Request.Context(), "Unknown log level, defaulting to info")
	}

	c.JSON(200, gin.H{
		"message": "Logged a " + level + " message",
	})
}
