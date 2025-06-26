package api

import (
	"errors"
	"github.com/JKasus/go_final_project/pkg/db"
	"github.com/JKasus/go_final_project/pkg/entities"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if idParam := r.URL.Query().Get("id"); idParam != "" {
		err = db.DeleteTask(idParam)
		if err != nil {
			writeJSON(w, err)
			return
		}
	} else {
		err = errors.New("id param is required")
		writeJSON(w, err)
		return
	}

	writeJSON(w, entities.EmptyResponse{})
}
