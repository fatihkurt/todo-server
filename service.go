package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"todoapp/todo-server/conn"
)

type Task struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type AddTaskRequest struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
}

func addTask(c *gin.Context) (Task, error) {

	rawData, err := c.GetRawData()

	if err != nil {
		return Task{}, err
	}

	var addTask AddTaskRequest
	err = json.Unmarshal(rawData, &addTask)

	if err != nil {
		return Task{}, err
	}

	res, err := conn.Db.Exec("INSERT INTO task(user_id, name) VALUES (?, ?)", addTask.UserId, addTask.Name)

	if err != nil {
		return Task{}, err
	}

	lastInsertedId, err := res.LastInsertId()

	if err != nil {
		return Task{}, err
	}

	var task Task

	row := conn.Db.QueryRow("SELECT id,user_id,name FROM task WHERE id=?", lastInsertedId)

	scanErr := row.Scan(&task.Id, &task.UserId, &task.Name)

	if scanErr != nil {
		return Task{}, err
	}

	return task, nil
}

func listTasks(c *gin.Context) ([]Task, error) {

	userIdParam := c.Query("userId")

	if userIdParam == "" {
		return nil, errors.New("userId param not specified." + userIdParam)
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return nil, err
	}

	rows, err := conn.Db.Query("SELECT id,user_id,name FROM task WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.UserId, &task.Name); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
