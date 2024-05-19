package storage

import (
	"server/internal/dto"
)

type CommonMessageManager interface {
	GetAgentMessageById(agentId int32) []*dto.HistoryMessageResp
	AddAgentMessageById(agentId int32, msg *dto.HistoryMessageResp)
	DeleteAgentMessageById(agentId int32)
}

type commonMessageManager struct {
	AgentMessages map[int32][]*dto.HistoryMessageResp
}

type HistoryMessageManager interface {
	GetSaveEarthAgentMessage() []*dto.HistoryMessageResp
	AddSaveEarthAgentMessage(msg *dto.HistoryMessageResp)
	DeleteSaveEarthAgentMessage()
}

type historyMessageManager struct {
	SaveEarthAgentMessages []*dto.HistoryMessageResp
}

func NewCommonMessageManager() CommonMessageManager {
	return &commonMessageManager{AgentMessages: make(map[int32][]*dto.HistoryMessageResp)}
}

func NewHistoryMessageManager() HistoryMessageManager {
	var saveEarthAgentMessages = make([]*dto.HistoryMessageResp, 0)
	return &historyMessageManager{
		SaveEarthAgentMessages: saveEarthAgentMessages,
	}
}

func (h *commonMessageManager) GetAgentMessageById(agentId int32) []*dto.HistoryMessageResp {
	if val, ok := h.AgentMessages[agentId]; ok {
		return val
	}
	return nil
}

func (h *commonMessageManager) AddAgentMessageById(agentId int32, msg *dto.HistoryMessageResp) {
	h.AgentMessages[agentId] = append(h.AgentMessages[agentId], msg)
}

func (h *commonMessageManager) DeleteAgentMessageById(agentId int32) {
	if _, ok := h.AgentMessages[agentId]; ok {
		h.AgentMessages[agentId] = make([]*dto.HistoryMessageResp, 0)
	}
}

func (h *historyMessageManager) GetSaveEarthAgentMessage() []*dto.HistoryMessageResp {
	return h.SaveEarthAgentMessages
}

func (h *historyMessageManager) AddSaveEarthAgentMessage(msg *dto.HistoryMessageResp) {
	h.SaveEarthAgentMessages = append(h.SaveEarthAgentMessages, msg)
}

func (h *historyMessageManager) DeleteSaveEarthAgentMessage() {
	h.SaveEarthAgentMessages = make([]*dto.HistoryMessageResp, 0)
}
