package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JKasus/go_final_project/pkg/db"
	"github.com/JKasus/go_final_project/pkg/entities"
	"github.com/JKasus/go_final_project/pkg/internal"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		err = errors.New("Error reading body: " + err.Error())
		writeJSON(w, err.Error())
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		err = errors.New("Error unmarshalling body: " + err.Error())
		writeJSON(w, err.Error())
		return
	}

	if task.Title == "" {
		err = errors.New("title is required")
		writeJSON(w, err.Error())
		return
	}

	if err = internal.CheckDate(&task); err != nil {
		err = errors.New("checkDate failed: " + err.Error())
		writeJSON(w, err.Error())
		return
	}

	taskId, err := db.AddTask(&task)
	if err != nil {
		err = errors.New("Error adding task: " + err.Error())
		writeJSON(w, err.Error())
		return
	}

	writeJSON(w, taskId)
}
