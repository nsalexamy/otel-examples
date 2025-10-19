package example

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsalexamy/otel-examples/otel-go-example/internal/logger"
)

func Sleep(c *gin.Context) {
	// get duration from query parameter, default to 1s
	durationStr := c.DefaultQuery("duration", "1s")
	logger.Info(c.Request.Context(), "Sleep handler called", map[string]interface{}{
		"duration": durationStr,
	})

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		logger.Error(c.Request.Context(), "Failed to parse duration", err)
		c.JSON(400, gin.H{"error": "Invalid duration"})
		return
	}

	// simulate work
	time.Sleep(duration)
	logger.Info(c.Request.Context(), "Finished sleeping", map[string]interface{}{
		"duration": durationStr,
	})

	c.JSON(200, gin.H{"message": "Slept for " + duration.String()})
}
