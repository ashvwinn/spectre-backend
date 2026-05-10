package main

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type envelope map[string]any

func (app *application) readIntParam(c *gin.Context, param string) (int, error) {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}
	return id, nil
}

func (app *application) readString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}
