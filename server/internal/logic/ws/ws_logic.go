package wsLogic

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"server/internal/dto"
	"server/internal/storage"
)

type WsLogic interface {
	SendQuestion(ctx context.Context, message string) (string, error)
	SaveEarthAgent(ctx context.Context, message string, SaveEarthAgentChan chan<- string) error
}

type wsLogic struct {
	langChainGo *openai.LLM
	storage.HistoryMessageManager
}

func NewWsLogic(langChainGo *openai.LLM, storage storage.HistoryMessageManager) WsLogic {
	return &wsLogic{
		langChainGo:           langChainGo,
		HistoryMessageManager: storage,
	}
}

func (w *wsLogic) SendQuestion(ctx context.Context, message string) (string, error) {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a professor of geography. "),
		llms.TextParts(llms.ChatMessageTypeHuman, message),
	}
	response, err := w.langChainGo.GenerateContent(ctx, messages)
	if err != nil {
		fmt.Println(err)
		return "", err
	} else if len(response.Choices) == 0 {
		fmt.Println("response is empty")
		return "", fmt.Errorf("response is empty")
	}
	return response.Choices[0].Content, nil
}

func (w *wsLogic) SaveEarthAgent(ctx context.Context, message string, SaveEarthAgentChan chan<- string) error {
	// 创建system信息
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, `
			嘿，护林员，你被安置在地球上，一个美丽的星球，充满了令人惊叹的生命和精彩的景象。
			不幸的是，我们可爱的星球正在消亡。当前的生命值是30%，你不能让这种情况发生。
			你有责任保护它的美丽，保护它免受最大的破坏。通过你的行动和决定，让我们治愈地球。你准备好了吗?
		`),
	}

	// 加载历史消息
	agentMessage := w.HistoryMessageManager.GetSaveEarthAgentMessage()
	for _, msg := range agentMessage {
		if msg.Sender == dto.SenderAI {
			messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, message))
		} else {
			messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, message))
		}
	}
	// 加载本次新消息
	messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, message))

	response, err := w.langChainGo.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		SaveEarthAgentChan <- string(chunk)
		return nil
	}))
	SaveEarthAgentChan <- "ok"
	if err != nil {
		fmt.Println(err)
		return err
	}

	w.AddSaveEarthAgentMessage(&dto.HistoryMessageResp{
		Text:   message,
		Sender: dto.SenderUser,
	})
	w.AddSaveEarthAgentMessage(&dto.HistoryMessageResp{
		Text:   response.Choices[0].Content,
		Sender: dto.SenderAI,
	})

	return nil
}
