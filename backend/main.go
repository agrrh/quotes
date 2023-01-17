// Package main - Quotes API
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	q "github.com/agrrh/quotes/backend/quotes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.RequestID())

	// TODO: Add CORS
	// 	https://echo.labstack.com/middleware/cors/
	// TODO: Add metrics:
	// 	https://echo.labstack.com/middleware/prometheus/

	// Routes

	e.GET("/", q.HandlerGetIndex)
	e.GET("/healthz", q.HandlerGetHealthz)

	e.GET("/quotes", q.HandlerGetQuotesList)
	e.POST("/quotes", q.HandlerCreateQuote)
	e.GET("/quotes/:id", q.HandlerGetQuote)
	e.PUT("/quotes/:id", q.HandlerUpdateQuote)
	e.DELETE("/quotes/:id", q.HandlerDeleteQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
