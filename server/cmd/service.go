package main

import (
	"fmt"
	"server/internal/config"
	"server/internal/handler"
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
	config.InitSQLiteDB()
	fmt.Println("SQLite db init success")

	return &router.HandlersImpl{
		TestHandler: handler.NewTestHandler(),
	}
}
