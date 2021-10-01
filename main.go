package main

import (
	"todoapp/todo-server/db"
	"todoapp/todo-server/server"

	"github.com/fvbock/endless"
)

func main() {
	db.ConnectDb()
	router := server.SetupRouter()
	err := endless.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}
}
