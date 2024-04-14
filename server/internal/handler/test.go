package handler

import (
	"net/http"
	"server/internal/logic"

	"github.com/gin-gonic/gin"
)

type TestHandler interface {
	Test(c *gin.Context)
}

type testHandler struct {
	testLogic logic.TestLogic
}

func NewTestHandler() TestHandler {
	return &testHandler{
		testLogic: logic.NewTestLogic(),
	}
}

func (h *testHandler) Test(c *gin.Context) {
	c.Status(http.StatusOK)
}
