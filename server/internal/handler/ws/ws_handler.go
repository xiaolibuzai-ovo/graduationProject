package wsHandler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"server/internal/dto"
	"server/internal/logic/ws"
	"server/internal/utils"
)

type WsHandler interface {
	SendQuestion(c *gin.Context)
	SaveEarthAgent(c *gin.Context)
}

type wsHandler struct {
	wsLogic.WsLogic
}

func NewWsHandler(
	logic wsLogic.WsLogic,
) WsHandler {
	return &wsHandler{
		WsLogic: logic,
	}
}

func (w *wsHandler) SendQuestion(c *gin.Context) {
	req := new(dto.SendQuestionReq)
	if err := c.ShouldBindQuery(req); err != nil {
		fmt.Printf("[SendQuestion] ShouldBindJSON err,err=%v", err)
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	supportFile, err := w.WsLogic.GetSupportFile(c.Request.Context(), req.AgentId)
	if err != nil {
		fmt.Printf("[SendQuestion] GetSupportFile err,err=%v", err)
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	// 获取WebSocket连接
	wsConn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Printf("[SendQuestion] websocket.Upgrade err,err=%v", err)
		utils.ErrorBadRequestResponse(c, err)
		return
	}

	for {
		mt, message, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Printf("[SendQuestion] ReadMessage err,err=%v", err)
			return
		}
		var (
			msgChan        = make(chan string, 10)
			nextSuggestion = make(chan string, 10)
		)
		go func() {
			err = w.WsLogic.SendQuestion(c.Request.Context(), req, string(message), supportFile, msgChan, nextSuggestion)
			if err != nil {
				fmt.Printf("[SendQuestion] SendQuestion err,err=%v", err)
			}
		}()

		var msgData *chatReplyMsg
		sendOK := false // 标志变量，表示是否已经发送过 "ok" 消息
		sendSuggestOk := false

		for {
			select {
			case msg, ok := <-msgChan:
				if ok {
					msgData = &chatReplyMsg{
						MsgType: 1,
						Content: msg,
					}

					bytes, _ := json.Marshal(msgData)
					err = wsConn.WriteMessage(mt, bytes)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					if !sendOK {
						msgData = &chatReplyMsg{
							MsgType: 1,
							Content: "ok",
						}
						bytes, _ := json.Marshal(msgData)
						wsConn.WriteMessage(mt, bytes) // 表示一次消息发送完成
						sendOK = true
					}
				}
			case nextSuggest, ok := <-nextSuggestion:
				if ok {
					msgData = &chatReplyMsg{
						MsgType: 2,
						Content: nextSuggest,
					}
					bytes, _ := json.Marshal(msgData)
					err = wsConn.WriteMessage(mt, bytes)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					if !sendSuggestOk {
						msgData = &chatReplyMsg{
							MsgType: 2,
							Content: "suggestOk",
						}
						bytes, _ := json.Marshal(msgData)
						wsConn.WriteMessage(mt, bytes) // 表示一次消息发送完成
						sendSuggestOk = true
					}
				}
			}

			if sendOK && sendSuggestOk {
				fmt.Println("chan listening break")
				break
			}
		}

	}
}

type chatReplyMsg struct {
	MsgType int16  `json:"msgType"` // 1-回复 2-建议
	Content string `json:"content"`
}

func (w *wsHandler) SaveEarthAgent(c *gin.Context) {
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
		var (
			saveEarthAgentChan = make(chan string, 10)
			nextSuggestion     = make(chan string, 10)
		)
		go func() {
			err := w.WsLogic.SaveEarthAgent(c.Request.Context(), string(message), saveEarthAgentChan, nextSuggestion)
			if err != nil {
				fmt.Println(err)
			}
		}()
		var msgData *chatReplyMsg
		sendOK := false // 标志变量，表示是否已经发送过 "ok" 消息
		sendSuggestOk := false
		for {
			select {
			case msg, ok := <-saveEarthAgentChan:
				if ok {
					msgData = &chatReplyMsg{
						MsgType: 1,
						Content: msg,
					}

					bytes, _ := json.Marshal(msgData)
					err = wsConn.WriteMessage(mt, bytes)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					if !sendOK {
						msgData = &chatReplyMsg{
							MsgType: 1,
							Content: "ok",
						}
						bytes, _ := json.Marshal(msgData)
						wsConn.WriteMessage(mt, bytes) // 表示一次消息发送完成
						sendOK = true
					}
				}
			case nextSuggest, ok := <-nextSuggestion:
				if ok {
					msgData = &chatReplyMsg{
						MsgType: 2,
						Content: nextSuggest,
					}
					bytes, _ := json.Marshal(msgData)
					err = wsConn.WriteMessage(mt, bytes)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					if !sendSuggestOk {
						msgData = &chatReplyMsg{
							MsgType: 2,
							Content: "suggestOk",
						}
						bytes, _ := json.Marshal(msgData)
						wsConn.WriteMessage(mt, bytes) // 表示一次消息发送完成
						sendSuggestOk = true
					}
				}
			}
			if sendOK && sendSuggestOk {
				fmt.Println("chan listening break")
				break
			}
		}
		//for msg := range saveEarthAgentChan {
		//
		//	//if msg == "ok" {
		//	//	break
		//	//}
		//}
	}
}
