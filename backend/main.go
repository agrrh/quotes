// Package main - Quotes API
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	h "github.com/agrrh/quotes/backend/handler"

	"github.com/go-playground/validator"
)

// CustomValidator - ???
type CustomValidator struct {
  validator *validator.Validate
}

// Validate - ???
func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally, you could return the error to give each route more control over the status code
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
	// Echo instance
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.RequestID())

	// TODO: Add CORS
	// 	https://echo.labstack.com/middleware/cors/
	// TODO: Add metrics:
	// 	https://echo.labstack.com/middleware/prometheus/

	// Routes

	e.GET("/", h.GetIndex)
	e.GET("/healthz", h.GetHealthz)

	e.GET("/quotes", h.GetQuotesList)
	e.POST("/quotes", h.CreateQuote)
	e.GET("/quotes/:qid", h.GetQuote)
	e.PUT("/quotes/:qid", h.UpdateQuote)
	e.DELETE("/quotes/:qid", h.DeleteQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
