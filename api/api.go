package api

import (
	"encoding/json"
	"net/http"
)

type TodoParams struct {
	Username string
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	Todos []Todo
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Error", http.StatusInternalServerError)
	}
)
