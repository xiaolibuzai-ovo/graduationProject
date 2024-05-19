package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/dto"
	"server/internal/logic"
	"server/internal/utils"
)

type MessagesHandler interface {
	GetHistoryMessage(c *gin.Context)
	DeleteHistoryMessage(c *gin.Context)
}

type messagesHandler struct {
	logic.MessagesLogic
}

func NewMessagesHandler(logic logic.MessagesLogic) MessagesHandler {
	return &messagesHandler{
		MessagesLogic: logic,
	}
}

func (m *messagesHandler) GetHistoryMessage(c *gin.Context) {
	req := new(dto.HistoryMessageReq)
	if err := c.ShouldBindJSON(req); err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	data, err := m.MessagesLogic.GetHistoryMessage(c.Request.Context(), req)
	if err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, data)
}

func (m *messagesHandler) DeleteHistoryMessage(c *gin.Context) {
	req := new(dto.DeleteHistoryMessageReq)
	if err := c.ShouldBindJSON(req); err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	err := m.MessagesLogic.DeleteHistoryMessage(c.Request.Context(), req)
	if err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, nil)
}
