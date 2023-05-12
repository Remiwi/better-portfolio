package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

func helloReq(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hi",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", helloReq)
	router.Run("0.0.0.0:8888")
}
