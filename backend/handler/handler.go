// Package handler -
package handler

import (
	"gopkg.in/mgo.v2"
)

// Handler - session container
type Handler struct {
	DB *mgo.Session
}

const (
	// DBName - name of database
	DBName = "quotes"

	// TableNameQuotes - name of quotes table
	TableNameQuotes = "quotes"
)
