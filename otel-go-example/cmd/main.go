package main

import (
	"os"

	"github.com/nsalexamy/otel-examples/otel-go-example/internal/example"
	"github.com/nsalexamy/otel-examples/otel-go-example/internal/logger"

	"context"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Init()

	// SERVER_ADDR is the address the server listens on. from environment variable or default to ":8080"
	var serverAddr = os.Getenv("SERVER_ADDR")

	if serverAddr == "" {
		serverAddr = ":8080"
	}

	logger.Info(context.Background(), "Starting server", map[string]interface{}{
		"server_addr": serverAddr,
	})

	r := gin.Default()

	r.Use(logger.Middleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ready",
		})
	})

	exampleGroup := r.Group("/")
	example.RegisterHandlers(exampleGroup)

	r.Run(serverAddr)
}
