package db

import (
	"errors"
	"fmt"
	"strconv"

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

func GetTaskById(id string) (*entities.Task, error) {
	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err = errors.New("failed to parse task id")
		return nil, err
	}

	var task entities.Task
	err = db.QueryRow(constants.QueryGetTaskById, taskId).Scan(&taskId, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		err = fmt.Errorf("failed to db.GetTaskBy: %s", err.Error())
		return nil, err
	}

	task.ID = strconv.Itoa(int(taskId))

	return &task, nil
}

func UpdateTask(task *entities.Task) error {
	taskId, err := strconv.ParseInt(task.ID, 10, 64)
	if err != nil {
		err = errors.New("failed to parse task id")
		return err
	}

	res, err := db.Exec(constants.QueryUpdateTask, taskId, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return err
	}
	// метод RowsAffected() возвращает количество записей к которым
	// был применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}
