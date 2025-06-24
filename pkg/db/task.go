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

func GetTaskList(filter *entities.Filter) ([]entities.Task, error) {

	rows, err := db.Query(constants.QueryGetTaskList, constants.SortASC, filter.Limit, filter.Offset)
	if err != nil {
		err = fmt.Errorf("failed to db.GetTaskList: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			err = fmt.Errorf("failed to Scan rows in Task entity: %s", err.Error())
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
