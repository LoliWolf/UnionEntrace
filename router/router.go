package router

import (
	"UnionEntrace/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	testGroup := r.Group("/test")
	{
		testGroup.GET("/get", handler.TestHandler)
	}
	return r
}
