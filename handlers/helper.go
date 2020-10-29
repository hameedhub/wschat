package handlers

import (
	"encoding/json"
	"net/http"
)

// Success struct for success response
type Success struct {
	Status  int
	Code    bool
	Data    []string
	Message string
}

// Error struct for error response
type Error struct {
	Status  int
	Code    bool
	Message string
}

// SuccessResponse - message on success request response
func SuccessResponse(w http.ResponseWriter, m string, d []string) (s *Success) {
	w.WriteHeader(s.Status)
	s.Code = true
	s.Message = m
	s.Data = d
	enc := json.NewEncoder(w)
	enc.Encode(s)
	return
}

// ErrorResponse - message on error request response
func (e *Error) ErrorResponse(w http.ResponseWriter, message string) error {
	w.WriteHeader(e.Status)
	e.Code = false
	e.Message = message
	enc := json.NewEncoder(w)
	enc.Encode(e)
	return nil
}
