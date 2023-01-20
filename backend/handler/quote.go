// Package handler - Quote
package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/agrrh/quotes/backend/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

var (
	// TimeNil - empty time value
	TimeNil = time.Time{}
)

// CreateQuote - create quote object
func (h *Handler) CreateQuote(c echo.Context) (err error) {
	q := &model.Quote{
		ID:      bson.NewObjectId(),
		AddedAt: time.Now(),
	}

	resp := model.MakeResponse()

	err = c.Bind(q)
	if err != nil {
		resp.Success = false
		resp.Message = "Error parsing data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// TODO: Validation
	// https://echo.labstack.com/cookbook/twitter/

	db := h.DB.Clone()
	defer db.Close()

	// Save quote in database
	err = db.DB(DBName).C(TableNameQuotes).Insert(q)
	if err != nil {
		resp.Success = false
		resp.Message = "Error during creating Quote"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Items = append(resp.Items, q)

	return c.JSON(http.StatusCreated, resp)
}

// FetchQuotes - fetch quotes
func (h *Handler) FetchQuotes(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	resp := model.MakeResponse()

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}

	// Retrieve quotes from database
	quotes := []*model.Quote{}
	db := h.DB.Clone()

	err = db.DB(DBName).C(TableNameQuotes).
		Find(
			bson.M{"approved_at": bson.M{"$ne": TimeNil}},
		).
		Sort("added_at").
		Skip((page - 1) * limit).
		Limit(limit).
		All(&quotes)
	if err != nil {
		resp.Success = false
		resp.Message = "Error during creating Quote"
		return c.JSON(http.StatusInternalServerError, resp)
	}
	defer db.Close()

	quotesPointers := make([]interface{}, len(quotes))
	for i, v := range quotes {
		quotesPointers[i] = &v
	}

	resp.Items = quotesPointers
	resp.Count = len(quotesPointers)

	return c.JSON(http.StatusOK, resp)
}

// ApproveQuote - approve quote
func (h *Handler) ApproveQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	qSeek := bson.M{"_id": bson.ObjectIdHex(id)}

	db := h.DB.Clone()
	defer db.Close()

	q := &model.Quote{}
	err = db.DB(DBName).C(TableNameQuotes).
		Find(qSeek).
		One(&q)
	if err != nil {
		resp.Success = false
		resp.Message = "Could not find Quote"
		return c.JSON(http.StatusNotFound, resp)
	}

	q.ApprovedAt = time.Now()

	// Approve quote in database
	err = db.DB(DBName).C(TableNameQuotes).Update(qSeek, q)
	if err != nil {
		resp.Success = false
		resp.Message = "Could not approve Quote"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Items = append(resp.Items, q)

	return c.JSON(http.StatusOK, resp)
}

// DenyQuote - deny quote
func (h *Handler) DenyQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	qSeek := bson.M{"_id": bson.ObjectIdHex(id)}

	db := h.DB.Clone()
	defer db.Close()

	q := &model.Quote{}
	err = db.DB(DBName).C(TableNameQuotes).
		Find(qSeek).
		One(&q)
	if err != nil {
		resp.Success = false
		resp.Message = "Could not find Quote"
		return c.JSON(http.StatusNotFound, resp)
	}

	q.ApprovedAt = TimeNil

	// Approve quote in database
	err = db.DB(DBName).C(TableNameQuotes).Update(qSeek, q)
	if err != nil {
		resp.Success = false
		resp.Message = "Could not deny Quote"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Items = append(resp.Items, q)

	return c.JSON(http.StatusOK, resp)
}

// DeleteQuote - delete quote
func (h *Handler) DeleteQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	db := h.DB.Clone()
	defer db.Close()

	err = db.DB(DBName).C(TableNameQuotes).RemoveId(id)
	if err != nil {
		resp.Success = false
		resp.Message = "Could not remove Quote"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
