package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type Response struct {
	Message []byte `json:"message"`
	Error   error  `json:"error"`
}

var router *gin.Engine

func setupRouter() *gin.Engine {
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
		c.String(200, "Server running")
	})

	group := router.Group("/api")
	{
		group.GET("/task", func(c *gin.Context) {
			result, err := listTasks(c)

			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": result})
		})

		group.POST("/task", func(c *gin.Context) {
			result, err := addTask(c)

			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": result})
		})
	}

	router.Run(":5000")
	return router
}
