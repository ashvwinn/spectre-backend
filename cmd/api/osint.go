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
		case errors.Is(osint.InvalidIPRepresentation, err):
			app.notFoundResponse(c) // TODO: change to failedValidationResponse after validations are implemented
		case errors.Is(osint.InvalidorEmptyResponse, err):
			app.notFoundResponse(c)
		case errors.Is(osint.FetchInfoError, err):
			app.serverErrorResponse(c, err)
		case errors.Is(osint.FailedParsingIntoJSON, err):
			app.serverErrorResponse(c, err)
		default:
			app.serverErrorResponse(c, err)
		}
		return
	}

	c.IndentedJSON(http.StatusOK, envelope{"iplookup": iplookupResult})
}
