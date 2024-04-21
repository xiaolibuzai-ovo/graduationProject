package wsHandler

import (
	"fmt"
	"server/internal/logic/ws"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type WsHandler interface {
	AddWsConnection(c *gin.Context)
}

type wsHandler struct {
	ws.WsLogic
}

func NewWsHandler(logic ws.WsLogic) WsHandler {
	return &wsHandler{
		WsLogic: logic,
	}
}

func (w *wsHandler) AddWsConnection(c *gin.Context) {
	// 升级并建立ws连接
	err := w.WsLogic.CreateWsConnection(c.Writer, c.Request)
	if err != nil {
		fmt.Printf("[AddWsConnection] createWsConnection err,err=%v", err)
		utils.ErrorInternalServerResponse(c, err)
		return
	}

	// 检测心跳
	go w.WsLogic.HeartBeatCheck()
	// 循环推送消息
	go w.WsLogic.PushLoop()

}
