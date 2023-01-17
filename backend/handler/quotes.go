// Package handler - Main
package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

// Quote - Single Quote
type Quote struct {
	Qid     int64    `json:"qid"`
	Author  string `json:"author" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// QuotesList - list of Quotes
type QuotesList []Quote

// GetQuotesList - GET /quotes
// TODO: How to paginate?
func GetQuotesList(c echo.Context) error {
	quotes, err := CRUD.List("*")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't list quotes")
	}
	return c.JSON(http.StatusOK, quotes)
}

// CreateQuote - POST /quotes
func CreateQuote(c echo.Context) error {
	q := &Quote{}
	if err := c.Bind(q); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	q.Qid = time.Now().UnixNano()
	qKey := fmt.Sprintf("quote:%v", q.Qid)

	qJSON, err := json.Marshal(q)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "Can't convert payload to JSON")
	}

	// TODO: Replace with response struct
	_, err = CRUD.Create(qKey, string(qJSON))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't create quote")
	}

	// TODO: Return response struct
	return c.JSON(http.StatusCreated, q)
}

// GetQuote - GET /quotes/:qid
func GetQuote(c echo.Context) error {
	qid, _ := strconv.Atoi(c.Param("qid"))

	q := Quote{}

	qKey := fmt.Sprintf("quote:%v", qid)
	data, err := CRUD.Read(qKey)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Can't get quote")
	}

	err = json.Unmarshal([]byte(data), &q)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't parse quote data")
	}

	return c.JSON(http.StatusOK, q)
}

// UpdateQuote - PUT /quotes/:qid
func UpdateQuote(c echo.Context) error {
	q := Quote{}
	if err := c.Bind(q); err != nil {
		return c.JSON(http.StatusBadRequest, "Can't parse data")
	}

	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: Generate some uuid?
	q.Qid = 1
	qKey := fmt.Sprintf("quote:%v", q.Qid)

	qJSON, err := json.Marshal(q)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "Can't convert payload to JSON")
	}

	// TODO: Replace with response struct
	_, err = CRUD.Create(qKey, string(qJSON))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't update quote")
	}

	// TODO: Return response struct
	return c.JSON(http.StatusCreated, q)
}

// DeleteQuote - DELETE /quotes/:qid
func DeleteQuote(c echo.Context) error {
	qid, _ := strconv.Atoi(c.Param("qid"))
	qKey := fmt.Sprintf("quote:%v", qid)

	_, err := CRUD.Delete(qKey)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Can't delete quote")
	}

	return c.JSON(http.StatusOK, "Quote deleted")
}
