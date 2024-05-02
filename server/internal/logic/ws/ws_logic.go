package wsLogic

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type WsLogic interface {
	SendQuestion(ctx context.Context, message string) (string, error)
}

type wsLogic struct {
	langChainGo *openai.LLM
}

func NewWsLogic(langChainGo *openai.LLM) WsLogic {
	return &wsLogic{
		langChainGo: langChainGo,
	}
}

func (w *wsLogic) SendQuestion(ctx context.Context, message string) (string, error) {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a professor of geography. "),
		llms.TextParts(llms.ChatMessageTypeHuman, message),
	}
	response, err := w.langChainGo.GenerateContent(ctx, messages)
	if err != nil {
		return "", err
	} else if len(response.Choices) == 0 {
		return "", fmt.Errorf("response is empty")
	}
	return response.Choices[0].Content, nil
}
