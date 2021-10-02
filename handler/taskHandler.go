package handler

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"todoapp/todo-server/model"
	"todoapp/todo-server/service"
)

type TaskHandler struct {
	Context *gin.Context
}

func NewTaskHandler(ctx *gin.Context) *TaskHandler {
	return &TaskHandler{
		Context: ctx,
	}
}

func (h *TaskHandler) HandleTaskList() {
	userIdParam := h.Context.Query("userId")

	if userIdParam == "" {
		service.ErrorResponse(h.Context, errors.New("userId param not specified."+userIdParam))
		return
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		service.ErrorResponse(h.Context, err)
		return
	}

	tasks, err := service.ListTasks(userId)
	if err != nil {
		service.ErrorResponse(h.Context, err)
		return
	}

	service.SendResponse(h.Context, gin.H{"message": tasks})
}

func (h *TaskHandler) HandleTaskAdd() {
	rawData, err := h.Context.GetRawData()

	if err != nil {
		service.ErrorResponse(h.Context, err)
		return
	}

	var task model.AddTaskRequest
	err = json.Unmarshal(rawData, &task)
	if err != nil {
		service.ErrorResponse(h.Context, err)
		return
	}

	if task.UserId == 0 {
		service.ErrorResponse(h.Context, errors.New("user id not provided"))
		return
	}

	if task.Name == "" {
		service.ErrorResponse(h.Context, errors.New("task not provided"))
		return
	}

	addedTask, err := service.AddTask(task)
	if err != nil {
		service.ErrorResponse(h.Context, err)
		return
	}

	service.SendResponse(h.Context, gin.H{"message": addedTask})
}
