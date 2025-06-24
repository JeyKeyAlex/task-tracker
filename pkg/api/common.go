package api

import (
	"encoding/json"
	"github.com/JKasus/go_final_project/pkg/entities"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, msg any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := make(map[string]interface{})
	switch v := msg.(type) {
	case string:
		w.WriteHeader(http.StatusBadRequest)
		resp["error"] = v
	case int64:
		w.WriteHeader(http.StatusOK)
		resp["id"] = strconv.FormatInt(v, 10)
	case []entities.Task:
		w.WriteHeader(http.StatusOK)
		resp["tasks"] = v
	default:
		w.WriteHeader(http.StatusBadRequest)
		resp["error"] = "invalid type of message"
	}
	json.NewEncoder(w).Encode(resp)
}
