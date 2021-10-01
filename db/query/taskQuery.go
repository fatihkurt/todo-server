package query

import (
	"todoapp/todo-server/db"
	"todoapp/todo-server/model"
)

func AddTask(task model.AddTaskRequest) (lastInsertedId int, err error) {
	res, err := db.Db.Exec("INSERT INTO task(user_id, name) VALUES (?, ?)", task.UserId, task.Name)

	if err != nil {
		return 0, err
	}

	lastInsertedId64, err := res.LastInsertId()

	return int(lastInsertedId64), err
}

func GetTaskById(id int) (model.Task, error) {
	var task model.Task

	row := db.Db.QueryRow("SELECT id,user_id,name FROM task WHERE id=?", id)

	scanErr := row.Scan(&task.Id, &task.UserId, &task.Name)

	return task, scanErr
}

func GetTasksByUserId(userId int) ([]model.Task, error) {
	rows, err := db.Db.Query("SELECT id,user_id,name FROM task WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]model.Task, 0)
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.UserId, &task.Name); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
