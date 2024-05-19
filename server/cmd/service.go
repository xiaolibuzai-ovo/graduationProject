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
	"server/internal/storage"
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
	gcsClient := config.NewGcs()
	fmt.Println("gcsClient init success")
	saveStorageManager := storage.NewHistoryMessageManager()
	agentStorageManager := storage.NewCommonMessageManager()
	return &router.HandlersImpl{
		AgentHandler:    handler.NewAgentHandler(logic.NewAgentLogic(dal.NewAgentDal(db), ml.NewLangChainGoClient(context.Background()))),
		WsHandler:       wsHandler.NewWsHandler(wsLogic.NewWsLogic(s.LangChainGoClient, saveStorageManager, agentStorageManager, dal.NewAgentDal(db))),
		MessagesHandler: handler.NewMessagesHandler(logic.NewMessagesLogic(saveStorageManager, agentStorageManager)),
		CommonHandler:   handler.NewCommonHandler(gcsClient),
	}
}
