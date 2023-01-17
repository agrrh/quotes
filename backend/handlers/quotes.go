// Package handlers - Main
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandlerGetQuotesList - GET /quotes
func HandlerGetQuotesList(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// HandlerCreateQuote - POST /quotes
func HandlerCreateQuote(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

// HandlerGetQuote - GET /quotes/:id
func HandlerGetQuote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

// HandlerUpdateQuote - PUT /quotes/:id
func HandlerUpdateQuote(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

// HandlerDeleteQuote - DELETE /quotes/:id
func HandlerDeleteQuote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
