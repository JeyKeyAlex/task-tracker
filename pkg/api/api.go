package api

import (
	"github.com/go-chi/chi/v5"
)

func Init(r chi.Router) {
	r.Get("/api/nextdate", getNextDayHandler)
	r.Post("/api/task", addTaskHandler)
	r.Get("/api/tasks", getTaskListHandler)
}
