package main

import (
	"context"
	"fmt"
	"server/internal/config"
	"server/internal/dal"
	"server/internal/handler"
	wsHandler "server/internal/handler/ws"
	"server/internal/logic"
	wsLogic "server/internal/logic/ws"
	"server/internal/ml"
	"server/internal/router"
)

type Service interface {
	NewHandlers() router.Handlers
}

type service struct {
	ctx context.Context
	*ml.LangChainGoClient
}

func NewService() Service {
	ctx := context.Background()
	langChainGoClient := ml.NewLangChainGoClient(ctx)
	return &service{
		ctx:               ctx,
		LangChainGoClient: langChainGoClient,
	}
}

func (s *service) NewHandlers() router.Handlers {
	db := config.InitSQLiteDB()
	fmt.Println("SQLite db init success")

	return &router.HandlersImpl{
		AgentHandler: handler.NewAgentHandler(logic.NewAgentLogic(dal.NewAgentDal(db))),
		WsHandler:    wsHandler.NewWsHandler(wsLogic.NewWsLogic(s.LangChainGoClient.LLM)),
	}
}
