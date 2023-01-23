// Package handler - Quote
package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/agrrh/quotes/backend/model"
)

var (
	timeNil = time.Time{}
)

// CreateQuote - create quote object
func (h *Handler) CreateQuote(c echo.Context) (err error) {
	quote := &model.Quote{
		AddedAt: time.Now(),
	}

	resp := model.MakeResponse()

	err = c.Bind(quote)
	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("Error parsing data: %s", err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	dbResult := h.DB.Create(&quote)

	if dbResult.Error != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("Error inserting quote: %s", dbResult.Error)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = quote

	return c.JSON(http.StatusCreated, resp)
}

// FetchQuotes - fetch quotes
func (h *Handler) FetchQuotes(c echo.Context) (err error) {
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	resp := model.MakeResponse()

	// Defaults
	if skip < 0 {
		skip = 0
	}
	if limit < 1 || limit > 20 {
		limit = 20
	}

	var quotes []interface{}

	// Retrieve quotes from database

	// NOTE: Default db.Find() passes slice of pointers which is not suitable for echo c.JSON() Response
	//	So here I iterate over records to form an array beforehand

	rows, err := h.DB.Model(&model.Quote{}).
		Where("approved_at <> ?", timeNil).
		Order("approved_at").
		Offset(skip).
		Limit(limit).
		Rows()
	defer rows.Close()

	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("Error fetching quotes: %s", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	for rows.Next() {
		var quote model.Quote
		h.DB.ScanRows(rows, &quote) // Row to struct
		quotes = append(quotes, quote)
	}

	resp.Items = quotes
	resp.Count = len(quotes)

	if resp.Count == 0 {
		resp.Message = "Empty quotes list"
	}

	return c.JSON(http.StatusOK, resp)
}

// FetchQuote - fetch single quote
func (h *Handler) FetchQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	var quote model.Quote

	dbResult := h.DB.
		Where("approved_at <> ?", timeNil).
		First(&quote, id)

	if dbResult.Error != nil {
		resp.Success = false

		if dbResult.Error == gorm.ErrRecordNotFound {
			resp.Message = fmt.Sprintf("Quote not found")
			return c.JSON(http.StatusNotFound, resp)
		}

		resp.Message = fmt.Sprintf("Error fetching quote: %s", dbResult.Error)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = quote

	return c.JSON(http.StatusOK, resp)
}

// ApproveQuote - approve quote
func (h *Handler) ApproveQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	var quote model.Quote

	dbResult := h.DB.
		First(&quote, id)

	if dbResult.Error != nil {
		resp.Success = false

		if dbResult.Error == gorm.ErrRecordNotFound {
			resp.Message = fmt.Sprintf("Quote not found")
			return c.JSON(http.StatusNotFound, resp)
		}

		resp.Message = fmt.Sprintf("Error approving quote: %s", dbResult.Error)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	h.DB.Model(&quote).Update("approved_at", time.Now())

	resp.Data = quote

	return c.JSON(http.StatusOK, resp)
}

// DenyQuote - deny quote
func (h *Handler) DenyQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	var quote model.Quote

	dbResult := h.DB.
		First(&quote, id)

	if dbResult.Error != nil {
		resp.Success = false

		if dbResult.Error == gorm.ErrRecordNotFound {
			resp.Message = fmt.Sprintf("Quote not found")
			return c.JSON(http.StatusNotFound, resp)
		}

		resp.Message = fmt.Sprintf("Error denying quote: %s", dbResult.Error)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	h.DB.Model(&quote).Update("approved_at", timeNil)

	resp.Data = quote

	return c.JSON(http.StatusOK, resp)
}

// DeleteQuote - delete quote
func (h *Handler) DeleteQuote(c echo.Context) (err error) {
	id := c.Param("id")

	resp := model.MakeResponse()

	var quote model.Quote

	dbResult := h.DB.
		First(&quote, id)

	if dbResult.Error != nil {
		resp.Success = false

		if dbResult.Error == gorm.ErrRecordNotFound {
			resp.Message = fmt.Sprintf("Quote not found")
			return c.JSON(http.StatusNotFound, resp)
		}

		resp.Message = fmt.Sprintf("Error deleting quote: %s", dbResult.Error)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	h.DB.Where("id = ?", id).Delete(&quote)

	resp.Data = quote

	return c.JSON(http.StatusOK, resp)
}
