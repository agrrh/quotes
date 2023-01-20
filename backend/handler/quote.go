// Package handler - Quote
package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/agrrh/quotes/backend/model"
	"gopkg.in/mgo.v2/bson"
)

var (
	// TimeNil - empty time value
	TimeNil = time.Time{}
)

// CreateQuote - create quote object
func (h *Handler) CreateQuote(c echo.Context) (err error) {
	q := &model.Quote{
		ID:   bson.NewObjectId(),
		AddedAt: time.Now(),
	}
	err = c.Bind(q)
	if err != nil {
		return
	}

	// TODO: Validation
	// https://echo.labstack.com/cookbook/twitter/

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()

	// Save post in database
	err = db.DB(DBName).C(TableNameQuotes).Insert(q)
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, q)
}

// FetchQuotes - fetch quotes
func (h *Handler) FetchQuotes(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}

	// Retrieve posts from database
	quotes := []*model.Quote{}
	db := h.DB.Clone()

	err = db.DB(DBName).C(TableNameQuotes).
		Find(
			bson.M{"approved_at": bson.M{"$ne": TimeNil}},
		).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&quotes)
	if err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, quotes)
}
