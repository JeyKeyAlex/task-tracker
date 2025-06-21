package db

import "github.com/JKasus/go_final_project/pkg/constants"

type Task struct {
	ID      string  `json:"id"`
	Date    string  `json:"date"`
	Title   string  `json:"title"`
	Comment *string `json:"comment"`
	Repeat  string  `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	res, err := db.Exec(constants.AddTaskQuery, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}
