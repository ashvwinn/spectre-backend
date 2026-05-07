package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthHandler(c *gin.Context) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	c.IndentedJSON(http.StatusOK, data)
}
