package logic

import (
	"context"
	"server/internal/dto"
	"server/internal/storage"
)

type MessagesLogic interface {
	GetHistoryMessage(ctx context.Context, req *dto.HistoryMessageReq) (data []*dto.HistoryMessageResp, err error)
	DeleteHistoryMessage(ctx context.Context, req *dto.DeleteHistoryMessageReq) (err error)
}

type messagesLogic struct {
	storage.HistoryMessageManager
	storage.CommonMessageManager
}

func NewMessagesLogic(history1 storage.HistoryMessageManager, history2 storage.CommonMessageManager) MessagesLogic {
	return &messagesLogic{
		HistoryMessageManager: history1,
		CommonMessageManager:  history2,
	}
}

func (m *messagesLogic) GetHistoryMessage(ctx context.Context, req *dto.HistoryMessageReq) (data []*dto.HistoryMessageResp, err error) {
	if req.AgentId == 99 {
		return m.HistoryMessageManager.GetSaveEarthAgentMessage(), nil
	}

	return m.CommonMessageManager.GetAgentMessageById(req.AgentId), nil
}

func (m *messagesLogic) DeleteHistoryMessage(ctx context.Context, req *dto.DeleteHistoryMessageReq) (err error) {
	if req.AgentId == 99 {
		m.HistoryMessageManager.DeleteSaveEarthAgentMessage()
		return
	}
	m.CommonMessageManager.DeleteAgentMessageById(req.AgentId)
	return
}
