package router

import "server/internal/handler"

type Handlers interface {
	handler.TestHandler
}

type HandlersImpl struct {
	handler.TestHandler
}
