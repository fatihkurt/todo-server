package model

type Task struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type AddTaskRequest struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
}
