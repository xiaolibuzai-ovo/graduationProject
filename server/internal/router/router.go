package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handlers Handlers) (http.Handler, error) {
	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(
		gin.Logger(),
	)
	router.GET("", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	api := router.Group("api")
	{
		api.GET("test", handlers.Test)
		wsApi := api.Group("ws")
		{
			wsApi.GET("send", handlers.SendQuestion)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	return router, nil
}
