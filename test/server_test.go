package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "todoapp/todo-server/db"
	"todoapp/todo-server/model"
	"todoapp/todo-server/server"
)

func TestTaskListRoute(t *testing.T) {
	ConnectDb()
	router := server.SetupRouter()

	w := performRequest(router, "GET", "/api/task?userId=1", nil)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddTaskRoute(t *testing.T) {
	ConnectDb()
	clearTable(Db)
	router := server.SetupRouter()

	task := model.AddTaskRequest{
		Name:   "Test task",
		UserId: 1,
	}

	var addingTask []byte
	addingTask, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	w := performRequest(router, "POST", "/api/task", addingTask)

	// expected data
	response := gin.H{
		"message": model.Task{
			Id:     1,
			UserId: 1,
			Name:   "Test task",
		},
	}
	resp, _ := json.Marshal(response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(resp), w.Body.String())
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	var req *http.Request
	if body != nil {
		req, _ = http.NewRequest(method, path, strings.NewReader(string(body)))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func clearTable(db *sql.DB) {
	_, err := Db.Exec("TRUNCATE table task")

	if err != nil {
		fmt.Println(err)
	}
}
