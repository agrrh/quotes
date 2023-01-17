// Package handler - General
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetIndex - GET /
func GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// GetHealthz - GET /healthz
func GetHealthz(c echo.Context) error {
	var item ResponseItem
	var items []ResponseItem
	var resp Response

	item = ResponseItem{"dummy": true}
	items = append(items, item)

	checkRedis := checkRedis()
	item = ResponseItem{"redis": checkRedis}
	items = append(items, item)

	resp.Items = items

	// Calculate success

	successList := make([]bool, 1, len(items))
	for _, value := range successList {
		successList = append(successList, value)
	}

	resp.Success = all(successList)

	return c.JSON(http.StatusOK, resp)
}
