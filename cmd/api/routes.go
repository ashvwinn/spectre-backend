package main

import "github.com/gin-gonic/gin"

func (app *application) registerRoutes(g *gin.Engine) {
	g.Use(slogMiddleware(app.logger))
	g.Use(gin.Recovery())

	g.GET("/health", app.healthHandler)
}
