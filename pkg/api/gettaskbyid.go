package api

import (
	"errors"
	"net/http"

	"github.com/JKasus/go_final_project/pkg/db"
	"github.com/JKasus/go_final_project/pkg/entities"
)

func getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {

	var task entities.Task

	if idParam := r.URL.Query().Get("id"); idParam != "" {
		task.ID = idParam
	} else {
		err := errors.New("id param is required")
		writeJSON(w, err.Error())
	}

	task, err := db.GetTaskById(task.ID)
	if err != nil {
		writeJSON(w, err.Error())
	}

	writeJSON(w, task)
}
