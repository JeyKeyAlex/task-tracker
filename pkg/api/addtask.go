package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/JKasus/go_final_project/pkg/constants"
	"github.com/JKasus/go_final_project/pkg/db"
)

func writeJSONError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	resp := map[string]string{"error": msg}
	json.NewEncoder(w).Encode(resp)
}

func writeJSONSuccess(w http.ResponseWriter, id string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	resp := map[string]string{"id": id}
	json.NewEncoder(w).Encode(resp)
}

func writeJSON(w http.ResponseWriter, msg any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	resp := make(map[string]string)
	switch v := msg.(type) {
	case string:
		resp["error"] = v
	case int64:
		resp["id"] = strconv.FormatInt(v, 10)
	}
	return json.NewEncoder(w).Encode(resp)
}

func checkDate(task *db.Task) error {
	now := time.Now()
	var next string
	if task.Date == "" {
		task.Date = now.Format(constants.DateFormat)
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		err = errors.New("Invalid date: " + task.Date)
		return err
	}
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			err = errors.New("Invalid repeat value: " + task.Repeat)
			return err
		}
	}
	if afterNow(now, t) {
		if len(task.Repeat) == 0 || now.Format(constants.DateFormat) == t.Format(constants.DateFormat) {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}
	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		err = errors.New("Error reading body: " + err.Error())
		writeJSONError(w, err.Error())
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		err = errors.New("Error unmarshalling body: " + err.Error())
		writeJSONError(w, err.Error())
		return
	}

	if task.Title == "" {
		err = errors.New("title is required")
		writeJSONError(w, err.Error())
		return
	}

	if err = checkDate(&task); err != nil {
		err = errors.New("checkDate failed: " + err.Error())
		writeJSONError(w, err.Error())
		return
	}

	taskId, err := db.AddTask(&task)
	if err != nil {
		err = errors.New("Error adding task: " + err.Error())
		writeJSONError(w, err.Error())
		return
	}

	writeJSONSuccess(w, strconv.FormatInt(taskId, 10))
}
