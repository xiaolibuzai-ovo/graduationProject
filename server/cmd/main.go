package main

import (
	"fmt"
	"net/http"
	"server/internal/router"
)

const (
	defaultPort uint16 = 8888
)

func main() {
	service := NewService()
	handlers := service.NewHandlers()

	router, err := router.NewRouter(handlers)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%d", defaultPort)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
