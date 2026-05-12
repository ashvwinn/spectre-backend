package main

import (
	"errors"
	"net/http"

	"github.com/ashvwinn/spectre-backend/internal/osint"
	"github.com/gin-gonic/gin"
)

func (app *application) iplookupHandler(c *gin.Context) {
	addr := c.Param("ip")

	iplookupResult, err := osint.IPLookupRun(addr)
	if err != nil {
		switch {
		case errors.Is(err, osint.InvalidIPRepresentation):
			app.notFoundResponse(c) // TODO: change to failedValidationResponse after validations are implemented
		default:
			app.serverErrorResponse(c, err)
		}
		return
	}

	if iplookupResult.Status == "fail" {
		app.notFoundResponse(c) // TODO: change to failedValidationResponse after validations are implemented
	}

	c.IndentedJSON(http.StatusOK, envelope{"iplookup": iplookupResult})
}
