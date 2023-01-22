// Package main - Quotes API
package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/agrrh/quotes/backend/handler"
	"github.com/agrrh/quotes/backend/model"
)

const (
	dsnEnvVar  = "APP_DB_DSN"
	dsnDefault = "quotes:password@tcp(127.0.0.1:3306)/quotes?charset=utf8mb4&parseTime=True&loc=Local"
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
	dsn := os.Getenv(dsnEnvVar)
	if dsn == "" {
		dsn = dsnDefault
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// TODO: Validations

	db.AutoMigrate(&model.Quote{})
	// db.AutoMigrate(&model.Quote{}, &Product{}, &Order{})

	// Initialize handler

	h := &handler.Handler{DB: db}

	// Routes

	e.GET("/", h.GetIndex)
	e.GET("/healthz", h.GetHealthz)

	e.POST("/quotes", h.CreateQuote)
	e.GET("/quotes", h.FetchQuotes)
	e.GET("/quotes/:id", h.FetchQuote)
	e.PUT("/quotes/:id/approve", h.ApproveQuote)
	e.PUT("/quotes/:id/deny", h.DenyQuote)
	e.DELETE("/quotes/:id", h.DeleteQuote)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
