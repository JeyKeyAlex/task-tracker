package api

import (
	"errors"
	"github.com/JKasus/go_final_project/pkg/db"
	"github.com/JKasus/go_final_project/pkg/entities"
	"github.com/JKasus/go_final_project/pkg/internal"
	"net/http"
	"time"
)

func CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task *entities.Task
	var err error
	var id string
	now := time.Now()

	if idParam := r.URL.Query().Get("id"); idParam != "" {
		id = idParam
		task, err = db.GetTaskById(id)
		if err != nil {
			writeJSON(w, err)
			return
		}
	} else {
		err = errors.New("id param is required")
		writeJSON(w, err)
		return
	}

	if task.Repeat == "" || &task.Repeat == nil {
		err = db.DeleteTask(id)
		writeJSON(w, entities.EmptyResponse{})
		if err != nil {
			writeJSON(w, err)
			return
		}
	} else {
		task.Date, err = internal.NextDate(now, task.Date, task.Repeat)
		err = db.UpdateTask(task)
		if err != nil {
			writeJSON(w, err)
			return
		}
		writeJSON(w, entities.EmptyResponse{})
	}
}
