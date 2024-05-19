package router

import (
	"server/internal/handler"
	wsHandler "server/internal/handler/ws"
)

type Handlers interface {
	handler.AgentHandler
	wsHandler.WsHandler
	handler.MessagesHandler
	handler.CommonHandler
}

type HandlersImpl struct {
	handler.AgentHandler
	wsHandler.WsHandler
	handler.MessagesHandler
	handler.CommonHandler
}
