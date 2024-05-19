package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/internal/dto"
	"server/internal/logic"
	"server/internal/utils"
)

type AgentHandler interface {
	AgentList(c *gin.Context)
	AgentDetail(c *gin.Context)
	Embedding(c *gin.Context)
	Suggests(c *gin.Context)
	LoadPdfData(c *gin.Context)
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

func (a *agentHandler) AgentDetail(c *gin.Context) {
	req := new(dto.AgentDetailReq)
	if err := c.ShouldBindJSON(req); err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	info, err := a.AgentLogic.AgentDetail(c.Request.Context(), req)
	if err != nil {
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, info)
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

func (a *agentHandler) Suggests(c *gin.Context) {
	req := new(dto.SuggestsReq)
	if err := c.ShouldBindJSON(req); err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	data, err := a.AgentLogic.Suggests(c.Request.Context(), req)
	if err != nil {
		fmt.Println("[Suggests] Suggests err", err)
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, &dto.SuggestsResp{SuggestsData: data})
}

func (a *agentHandler) LoadPdfData(c *gin.Context) {
	req := new(dto.LoadPdfReq)
	if err := c.ShouldBindJSON(req); err != nil {
		utils.ErrorBadRequestResponse(c, err)
		return
	}
	fmt.Println(req.Files)
	err := a.AgentLogic.LoadPdfData(c.Request.Context(), req)
	if err != nil {
		fmt.Println("[Suggests] Suggests err", err)
		utils.ErrorInternalServerResponse(c, err)
		return
	}
	utils.SuccessResponse(c, nil)
}
