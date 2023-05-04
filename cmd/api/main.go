package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!!",
		})
	})
	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
