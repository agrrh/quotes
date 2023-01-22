// Package model - Quote
package model

import (
	"time"

	"gorm.io/gorm"
)

// Quote - quote object
type Quote struct {
	ID         uint           `gorm:"primaryKey" json:"id,omitempty"`
	Author     string         `gorm:"not null" json:"author" valid:"required,length(3|64)"`
	Content    string         `gorm:"not null" json:"content" valid:"required,length(8|1024)"`
	AddedAt    time.Time      `gorm:"index" json:"added_at"`
	ApprovedAt time.Time      `gorm:"index" json:"approved_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TODO: Validation
