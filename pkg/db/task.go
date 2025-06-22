package db

import (
	"fmt"

	"github.com/JKasus/go_final_project/pkg/constants"
	"github.com/JKasus/go_final_project/pkg/entities"
)

func AddTask(task *entities.Task) (int64, error) {
	var id int64
	res, err := db.Exec(constants.QueryAddTask, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func GetTaskList(limit, offset int64) ([]entities.Task, error) {
	if limit == 0 {
		limit = constants.DefaultLimit
	}
	rows, err := db.Query(constants.QueryGetTaskList, constants.SortASC, limit, offset)
	if err != nil {
		err = fmt.Errorf("Failed to db.GetTaskList: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			err = fmt.Errorf("Failed to Scan rows in Task entity: %s", err.Error())
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
