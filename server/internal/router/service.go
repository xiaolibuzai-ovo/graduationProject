package router

import (
	"server/internal/handler"
	wsHandler "server/internal/handler/ws"
)

type Handlers interface {
	handler.AgentHandler
	wsHandler.WsHandler
}

type HandlersImpl struct {
	handler.AgentHandler
	wsHandler.WsHandler
}
