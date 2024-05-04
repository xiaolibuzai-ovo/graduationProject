package logic

import (
	"context"
	"server/internal/dal"
	"server/internal/dto"
)

type AgentLogic interface {
	AgentList(ctx context.Context) (agents []*dto.AgentListData, err error)
	Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error)
}

type agentLogic struct {
	dal.AgentDal
}

func NewAgentLogic(agent dal.AgentDal) AgentLogic {
	return &agentLogic{AgentDal: agent}
}

func (a *agentLogic) AgentList(ctx context.Context) (agents []*dto.AgentListData, err error) {
	allAgents, err := a.AgentDal.GetAllAgents(ctx)
	if err != nil {
		return
	}
	for _, agent := range allAgents {
		item := &dto.AgentListData{
			Id:       agent.Id,
			Img:      agent.Img,
			Title:    agent.Title,
			Subtitle: agent.Subtitle,
			Content:  agent.Content,
		}
		agents = append(agents, item)
	}
	return
}

func (a *agentLogic) Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error) {
	// 拿到私有数据Embedding

	return
}
