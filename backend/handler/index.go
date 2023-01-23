// Package handler -
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/agrrh/quotes/backend/model"
)

// GetIndex - index page
func (h *Handler) GetIndex(c echo.Context) (err error) {
	resp := model.MakeResponse()

	resp.Message = "Quotes API"

	return c.JSON(http.StatusOK, resp)
}
