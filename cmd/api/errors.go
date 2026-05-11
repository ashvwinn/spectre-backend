package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = debug.Stack()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "stack", string(trace))
}

func (app *application) errorResponse(c *gin.Context, status int, message any, err error) {
	env := envelope{"error": message, "logs": err}
	c.IndentedJSON(status, env)
}

func (app *application) serverErrorResponse(c *gin.Context, err error) {
	app.logError(c.Request, err)
	message := "the server encounterd a problem and could not process your request"
	app.errorResponse(c, http.StatusInternalServerError, message, err)
}

func (app *application) notFoundResponse(c *gin.Context) {
	message := "the requested resource could not be found"
	app.errorResponse(c, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowedResponse(c *gin.Context) {
	message := fmt.Sprintf("the %s method is not supported for this resource", c.Request.Method)
	app.errorResponse(c, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) badRequestResponse(c *gin.Context, err error) {
	app.errorResponse(c, http.StatusBadRequest, err.Error(), nil)
}
