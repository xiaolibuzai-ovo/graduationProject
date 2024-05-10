package storage

import (
	"server/internal/dto"
)

type HistoryMessageManager interface {
	GetSaveEarthAgentMessage() []*dto.HistoryMessageResp
	AddSaveEarthAgentMessage(msg *dto.HistoryMessageResp)
}
type historyMessageManager struct {
	SaveEarthAgentMessages []*dto.HistoryMessageResp
}

func NewHistoryMessageManager() HistoryMessageManager {
	var saveEarthAgentMessages []*dto.HistoryMessageResp
	return &historyMessageManager{
		SaveEarthAgentMessages: saveEarthAgentMessages,
	}
}

func (h historyMessageManager) GetSaveEarthAgentMessage() []*dto.HistoryMessageResp {
	return h.SaveEarthAgentMessages
}

func (h historyMessageManager) AddSaveEarthAgentMessage(msg *dto.HistoryMessageResp) {
	h.SaveEarthAgentMessages = append(h.SaveEarthAgentMessages, msg)
}
