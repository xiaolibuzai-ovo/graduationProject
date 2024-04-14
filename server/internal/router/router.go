package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handlers Handlers) (http.Handler, error) {
	router := gin.New()
	router.Use()
	router.GET("", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	api := router.Group("api")
	{
		api.GET("test", handlers.Test)
	}
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	return router, nil
}
