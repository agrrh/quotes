// Package handler - CRUD interface
package handler

import (
	"github.com/agrrh/quotes/backend/handler/crud"
)

// CRUD - client init
// FIXME: Parametrize options
var CRUD = crud.CreateClient("localhost:6379", "", 0)
