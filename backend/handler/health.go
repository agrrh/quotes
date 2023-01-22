// Package handler -
package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/agrrh/quotes/backend/model"
)

type healthCheck struct {
	Component string `json:"component"`
	Descr     string `json:"description"`
	Result    bool   `json:"result"`
}

func newHealthCheck(name string) healthCheck {
	check := healthCheck{}
	check.Component = name
	check.Result = true
	check.Descr = "passed, no additional info"

	return check
}

// GetHealthz - health check
func (h *Handler) GetHealthz(c echo.Context) (err error) {
	resp := model.MakeResponse()

	// Dummy

	dummyCheck := newHealthCheck("dummy")
	dummyCheck.Descr = "https://www.youtube.com/watch?v=ziE7YkOiOBU"

	resp.Items = append(resp.Items, dummyCheck)

	// DB

	dbCheck := newHealthCheck("db")

	sqlDB, err := h.DB.DB()

	err = sqlDB.Ping()
	if err != nil {
		dbCheck.Result = false
		dbCheck.Descr = fmt.Sprintf("Error accessing DB: %s", err)
	}

	resp.Items = append(resp.Items, dbCheck)

	// Summarize

	resp.Success = dummyCheck.Result && dbCheck.Result

	if !resp.Success {
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
