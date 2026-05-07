package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func slogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.RequestURI()

		c.Next()

		logger.Info("request",
			slog.String("ip", c.ClientIP()),
			slog.String("method", c.Request.Method),
			slog.Int("status", c.Writer.Status()),
			slog.String("uri", uri),
		)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error("request error", slog.String("error", err.Error()))
			}
		}
	}
}
