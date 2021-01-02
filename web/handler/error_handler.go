package handler

import (
	"encoding/json"
	"net/http"
)

type AnuError struct {
	Error        bool   `json:"error"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func WriteNotFound(w http.ResponseWriter) {
	writeError(404, "Not found", w)
}

func WriteMethodNotAllowed(w http.ResponseWriter) {
	writeError(405, "Method not allowed", w)
}

func writeBadRequest(w http.ResponseWriter) {
	writeError(400, "Bad request", w)
}

func writeInternalServerError(w http.ResponseWriter) {
	writeError(500, "Internal server error", w)
}

func writeError(code int, message string, w http.ResponseWriter) {
	w.WriteHeader(code)
	r, _ := json.Marshal(AnuError{
		Error:        true,
		ErrorCode:    code,
		ErrorMessage: message,
	})
	_, _ = w.Write(r)
}
