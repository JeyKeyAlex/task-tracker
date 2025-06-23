package api

import (
	"errors"
	"github.com/JKasus/go_final_project/pkg/internal"
	"net/http"
	"strconv"

	"github.com/JKasus/go_final_project/pkg/db"
	"github.com/JKasus/go_final_project/pkg/entities"
)

func getTaskListHandler(w http.ResponseWriter, r *http.Request) {

	var filter entities.Filter

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		limit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			err = errors.New("failed to parse 'limit' parameter")
			internal.WriteJSON(w, err.Error())
			return
		}
		filter.Limit = limit
	}
	if offsetParam := r.URL.Query().Get("offset"); offsetParam != "" {
		offset, err := strconv.ParseInt(offsetParam, 10, 64)
		if err != nil {
			err = errors.New("failed to parse 'offset' parameter")
			internal.WriteJSON(w, err.Error())
			return
		}
		filter.Offset = offset
	}

	taskList, err := db.GetTaskList(filter.Limit, filter.Offset)
	if err != nil {
		internal.WriteJSON(w, err.Error())
		return
	}

	if taskList == nil {
		internal.WriteJSON(w, []entities.Task{})
	} else {
		internal.WriteJSON(w, taskList)
	}
}
