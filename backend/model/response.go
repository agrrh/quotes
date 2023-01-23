// Package model - Response
package model

// Response - response object
type Response struct {
	Message string        `json:"message,omitempty"`
	Success bool          `json:"success"`
	Count   int           `json:"count,omitempty"`
	Items   []interface{} `json:"items,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
}

// MakeResponse - return Response with default values
func MakeResponse() Response {
	resp := Response{}
	resp.Success = true
	return resp
}
