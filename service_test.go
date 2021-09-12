package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskListRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/list", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestAddTaskRoute(t *testing.T) {
	task := Task{
		Name: "Test task",
		Done: false,
	}

	var addedTask []byte
	addedTask, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	router := setupRouter()

	w := httptest.NewRecorder()
	var body = strings.NewReader(string(addedTask))

	req, _ := http.NewRequest("POST", "/task", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(addedTask), w.Body.String())
}
