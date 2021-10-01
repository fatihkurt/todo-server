package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, response map[string]interface{}) {
	c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
}
