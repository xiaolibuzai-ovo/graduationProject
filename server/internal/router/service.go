package router

import (
	"server/internal/handler"
	wsHandler "server/internal/handler/ws"
)

type Handlers interface {
	handler.TestHandler
	wsHandler.WsHandler
}

type HandlersImpl struct {
	handler.TestHandler
	wsHandler.WsHandler
}
