package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func PostHandler(c *gin.Context) {
	// 获取前端传进来的JSON参数
	var json struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断JSON传递数据是否为空
	if json.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is empty"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"name": json})
	}
}
