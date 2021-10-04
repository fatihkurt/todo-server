package server

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"todoapp/todo-server/handler"
)

type Response struct {
	Message []byte `json:"message"`
	Error   error  `json:"error"`
}

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Server running.")
	})

	group := router.Group("/api")
	{
		group.GET("/task", func(c *gin.Context) {
			h := handler.NewTaskHandler(c)
			h.HandleTaskList()
		})

		group.POST("/task", func(c *gin.Context) {
			h := handler.NewTaskHandler(c)
			h.HandleTaskAdd()
		})
	}

	return router
}
