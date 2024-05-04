package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/dto"
	"server/internal/logic"
	"server/internal/utils"
)

type AgentHandler interface {
	AgentList(c *gin.Context)
	Embedding(c *gin.Context)
}

type agentHandler struct {
	logic.AgentLogic
}

func NewAgentHandler(agent logic.AgentLogic) AgentHandler {
	return &agentHandler{
		AgentLogic: agent,
	}
}

func (a *agentHandler) AgentList(c *gin.Context) {
	agents, err := a.AgentLogic.AgentList(c.Request.Context())
	if err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, &dto.AgentListResp{Agents: agents})
}

func (a *agentHandler) Embedding(c *gin.Context) {
	var (
		req = new(dto.EmbeddingReq)
		err error
	)
	if err = c.ShouldBindJSON(req); req != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	err = a.AgentLogic.Embedding(c.Request.Context(), req)
	if err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, nil)
}
