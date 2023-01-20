// Package model - Quote
package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

// Quote - quote object
type Quote struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Author     string        `json:"author" bson:"author"`
	Content    string        `json:"content" bson:"content"`
	AddedAt    time.Time `json:"added_at" bson:"added_at"`
	ApprovedAt time.Time `json:"approved_at" bson:"approved_at"`
}
