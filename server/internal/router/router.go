package router

import (
	"net/http"
	"server/internal/mw"

	"github.com/gin-gonic/gin"
)

func NewRouter(handlers Handlers) (http.Handler, error) {
	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(
		gin.Logger(),
		mw.Cors(),
	)
	router.GET("", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	api := router.Group("api")
	{
		agentApi := api.Group("agent")
		{
			agentApi.GET("list", handlers.AgentList)
			agentApi.POST("detail", handlers.AgentDetail)
			agentApi.POST("embedding", handlers.Embedding) // 嵌入向量
			agentApi.POST("suggests", handlers.Suggests)
			agentApi.POST("loadPdfData", handlers.LoadPdfData)
		}
		wsApi := api.Group("ws")
		{
			wsApi.GET("send", handlers.SendQuestion)
			wsApi.GET("saveEarthAgent", handlers.SaveEarthAgent)
		}
		msgApi := api.Group("message")
		{
			msgApi.POST("messages", handlers.GetHistoryMessage)
			msgApi.POST("delete", handlers.DeleteHistoryMessage)
		}
		commonApi := api.Group("common")
		{
			commonApi.POST("upload", handlers.UploadFile)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})
	return router, nil
}
