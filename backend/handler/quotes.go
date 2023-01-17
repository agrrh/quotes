// Package handler - Main
package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"time"
	"math"

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
// FIXME: How to properly paginate? Redis KEYS does not actually sort keys.
// FIXME: Simplify this method
func GetQuotesList(c echo.Context) error {
	var offset, limit int
	var qids []string
	var quotes []Quote

	offset, _ = strconv.Atoi(c.QueryParam("offset"))
	if offset == 0 {
		offset = 0
	}
	limit, _ = strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 20
	}

	qids, err := CRUD.List("quote:*")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't list quotes")
	}

	if offset > len(qids) {
		offset = int(math.Max(0, float64(len(qids) - limit)))
	}

	if offset + limit > len(qids) {
		limit = len(qids) - offset
	}

	qids = qids[offset : offset + limit]

	for _, qid := range qids {
		data, err := CRUD.Read(qid)

		q := &Quote{}

		err = json.Unmarshal([]byte(data), &q)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Can't parse quote data")
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Can't get listed quote")
		}

		quotes = append(quotes, *q)
	}

	// TODO: Return response struct
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

	// TODO: Return response struct
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

	// TODO: Return response struct
	return c.JSON(http.StatusOK, "Quote deleted")
}
