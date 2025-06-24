package api

import (
	"encoding/json"
	"github.com/JKasus/go_final_project/pkg/entities"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, msg any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var resp any
	var statusCode int

	switch v := msg.(type) {
	case string:
		statusCode = http.StatusBadRequest
		resp = map[string]interface{}{"error": v}
	case int64:
		statusCode = http.StatusOK
		resp = map[string]interface{}{"id": strconv.FormatInt(v, 10)}
	case []entities.Task:
		statusCode = http.StatusOK
		resp = map[string]interface{}{"tasks": v}
	case *entities.Task:
		statusCode = http.StatusOK
		resp = v
	default:
		statusCode = http.StatusBadRequest
		resp = map[string]interface{}{"error": "invalid type of message"}
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
