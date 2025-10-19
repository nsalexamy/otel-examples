package example

import (
	"github.com/gin-gonic/gin"
	"github.com/nsalexamy/otel-examples/otel-go-example/internal/logger"
)

func RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		logger.Info(c.Request.Context(), "Hello handler called", nil)

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/logging", Logging)
	router.GET("/sleep", Sleep)
}
