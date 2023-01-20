// Package main - Quotes API
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/agrrh/quotes/backend/handler"

	"gopkg.in/mgo.v2"
)

func main() {
	// Echo instance
	e := echo.New()

	// TODO: Attach JWT:
	// 	https://github.com/labstack/echox/blob/master/cookbook/twitter/server.go

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// TODO: Add CORS:
	// 	https://echo.labstack.com/middleware/cors/
	// TODO: Add metrics:
	// 	https://echo.labstack.com/middleware/prometheus/

	// Database connection
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// TODO: Ensure indexes

	// Initialize handler

	h := &handler.Handler{DB: db}

	// Routes

	// e.GET("/", h.GetIndex)
	// e.GET("/healthz", h.GetHealthz)

	e.GET("/quotes", h.FetchQuotes)
	e.POST("/quotes", h.CreateQuote)
	e.PUT("/quotes/:id/approve", h.ApproveQuote)
	e.PUT("/quotes/:id/deny", h.DenyQuote)
	// e.PUT("/quotes/:id", h.UpdateQuote)
	// e.DELETE("/quotes/:id", h.DeleteQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
