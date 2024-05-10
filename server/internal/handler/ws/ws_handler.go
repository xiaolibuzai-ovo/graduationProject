package wsHandler

import (
	"fmt"
	"server/internal/logic/ws"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsHandler interface {
	SendQuestion(c *gin.Context)
	SaveEarthAgent(c *gin.Context)
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
			fmt.Println(err)
			return
		}
		resp, err := w.WsLogic.SendQuestion(c.Request.Context(), string(message))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = wsConn.WriteMessage(mt, []byte(resp))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (w *wsHandler) SaveEarthAgent(c *gin.Context) {
	wsConn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	var saveEarthAgentChan = make(chan string, 10)

	for {
		mt, message, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			err := w.WsLogic.SaveEarthAgent(c.Request.Context(), string(message), saveEarthAgentChan)
			if err != nil {
				fmt.Println(err)
			}
		}()
		for msg := range saveEarthAgentChan {
			err = wsConn.WriteMessage(mt, []byte(msg))
			if err != nil {
				fmt.Println(err)
			}
			if msg == "ok" {
				break
			}
		}
		wsConn.WriteMessage(mt, []byte("ok")) // 表示一次消息发送完成
	}
}
