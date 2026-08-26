package models

// JSONValue represents an API value whose concrete shape varies by object set.
type JSONValue = any

// Metadata contains arbitrary key-value metadata returned by several object sets.
type Metadata map[string]any

// APIError is the common error envelope returned by NimbleOS.
type APIError struct {
	Code      string `json:"code,omitempty"`
	Text      string `json:"text,omitempty"`
	Arguments any    `json:"arguments,omitempty"`
}

// APIResponse is the common NimbleOS response envelope.
type APIResponse[T any] struct {
	Data  T         `json:"data,omitempty"`
	Error *APIError `json:"error,omitempty"`
}

// APIRequest is the common NimbleOS request envelope.
type APIRequest[T any] struct {
	Data T `json:"data,omitempty"`
}
