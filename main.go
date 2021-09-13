package main

import (
	"todoapp/todo-server/conn"

	"github.com/fvbock/endless"
)

func main() {
	conn.ConnectDb()
	router := setupRouter()
	err := endless.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
