// Package handler -
package handler

import (
	"gorm.io/gorm"
)

// Handler - session container
type Handler struct {
	DB *gorm.DB
}
