package logic

import (
	"context"
	"server/internal/dto"
	"server/internal/storage"
)

type MessagesLogic interface {
	GetHistoryMessage(ctx context.Context, req *dto.HistoryMessageReq) (data []*dto.HistoryMessageResp, err error)
}

type messagesLogic struct {
	storage.HistoryMessageManager
}

func NewMessagesLogic(history storage.HistoryMessageManager) MessagesLogic {
	return &messagesLogic{
		HistoryMessageManager: history,
	}
}

func (m *messagesLogic) GetHistoryMessage(ctx context.Context, req *dto.HistoryMessageReq) (data []*dto.HistoryMessageResp, err error) {
	return m.HistoryMessageManager.GetSaveEarthAgentMessage(), nil
}
