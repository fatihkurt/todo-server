package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id   int
	Name string
	Done bool
}

func addTask(c *gin.Context) []byte {
	var task []byte
	// TODO add task
	task, err := json.Marshal("")

	if err != nil {
		panic("Error while json encoding")
	}

	return task
}

func listTasks() []byte {
	var tasks []byte
	// TODO get tasks
	tasks, err := json.Marshal("")

	if err != nil {
		panic("Error while json encoding")
	}

	return tasks
}
