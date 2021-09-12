package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/list", func(c *gin.Context) {
		var response = listTasks()
		c.JSON(200, response)
	})

	router.GET("/task", func(c *gin.Context) {
		var response = addTask(c)
		c.JSON(200, response)
	})
	return router
}
