package wsHandler

import (
	"server/internal/logic/ws"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsHandler interface {
	SendQuestion(c *gin.Context)
}

type wsHandler struct {
	wsLogic.WsLogic
}

func NewWsHandler(logic wsLogic.WsLogic) WsHandler {
	return &wsHandler{
		WsLogic: logic,
	}
}

func (w *wsHandler) SendQuestion(c *gin.Context) {
	// 获取WebSocket连接
	wsConn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	for {
		mt, message, err := wsConn.ReadMessage()
		if err != nil {
			utils.ErrorInternalServerResponse(c, err)
			return
		}
		resp, err := w.WsLogic.SendQuestion(c.Request.Context(), string(message))
		if err != nil {
			utils.ErrorInternalServerResponse(c, err)
			return
		}
		err = wsConn.WriteMessage(mt, []byte(resp))
		if err != nil {
			utils.ErrorInternalServerResponse(c, err)
			return
		}
	}
}
