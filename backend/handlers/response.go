// Package handlers - Reponse models
package handlers

// Response - composed response object, representing meta stuff e.g. success and items count
type Response struct {
	Success bool           `json:"success"`
	Items   []ResponseItem `json:"items"`
	Count   int            `json:"count,omitempty"`
}

// ResponseItem - single response object
type ResponseItem map[string]any
