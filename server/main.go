package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.Run(":8000")
}
