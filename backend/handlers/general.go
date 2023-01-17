// Package handlers - General
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandlerGetIndex - GET /
func HandlerGetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// HandlerGetHealthz - GET /healthz
func HandlerGetHealthz(c echo.Context) error {
	var items []ResponseItem
	var resp Response

	item := ResponseItem{"foo": "bar", "baz": 42, "zap": true}
	items = append(items, item)

	resp.Success = true
	resp.Items = items

	return c.JSON(http.StatusOK, resp)
}
