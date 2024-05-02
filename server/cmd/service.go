package main

import (
	"github.com/tmc/langchaingo/llms/openai"
	"server/internal/handler"
	wsHandler "server/internal/handler/ws"
	wsLogic "server/internal/logic/ws"
	"server/internal/router"
)

type Service interface {
	NewHandlers() router.Handlers
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) NewHandlers() router.Handlers {
	//db := config.InitSQLiteDB()
	//fmt.Println("SQLite db init success")
	//_ = db
	langChainGo, err := openai.New()
	if err != nil {
		panic(err)
	}
	return &router.HandlersImpl{
		TestHandler: handler.NewTestHandler(),
		WsHandler:   wsHandler.NewWsHandler(wsLogic.NewWsLogic(langChainGo)),
	}
}
