package service

import (
	"todoapp/todo-server/db/query"
	"todoapp/todo-server/model"
)

func AddTask(task model.AddTaskRequest) (model.Task, error) {

	lastInsertedId, err := query.AddTask(task)

	if err != nil {
		return model.Task{}, err
	}

	return query.GetTaskById(lastInsertedId)
}

func ListTasks(userId int) ([]model.Task, error) {

	return query.GetTasksByUserId(userId)
}
