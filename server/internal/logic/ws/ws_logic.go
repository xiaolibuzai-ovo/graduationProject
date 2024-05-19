package wsLogic

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/milvus"
	"server/internal/constant"
	"server/internal/dal"
	"server/internal/dto"
	"server/internal/ml"
	"server/internal/storage"
)

type WsLogic interface {
	GetSupportFile(ctx context.Context, agentId int64) (supportFile int16, err error)
	SendQuestion(ctx context.Context, req *dto.SendQuestionReq, message string, supportFile int16, msgChan, nextSuggestion chan<- string) error
	SaveEarthAgent(ctx context.Context, message string, SaveEarthAgentChan chan<- string, nextSuggestion chan<- string) error
}

type wsLogic struct {
	langChainGo *ml.LangChainGoClient
	storage.HistoryMessageManager
	storage.CommonMessageManager
	dal.AgentDal
}

func NewWsLogic(langChainGo *ml.LangChainGoClient, storage1 storage.HistoryMessageManager, storage2 storage.CommonMessageManager, agent dal.AgentDal) WsLogic {
	return &wsLogic{
		langChainGo:           langChainGo,
		HistoryMessageManager: storage1,
		CommonMessageManager:  storage2,
		AgentDal:              agent,
	}
}

func (w *wsLogic) GetSupportFile(ctx context.Context, agentId int64) (supportFile int16, err error) {
	agentList, err := w.AgentDal.GetAgentsById(ctx, int32(agentId))
	if err != nil {
		return
	}
	if len(agentList) == 0 {
		return
	}
	return agentList[0].SupportFile, nil
}

func (w *wsLogic) SendQuestion(ctx context.Context, req *dto.SendQuestionReq, message string, supportFile int16, msgChan, nextSuggestion chan<- string) error {
	switch supportFile {
	case 0: // false
		return w.sendTextQuestion(ctx, req, message, msgChan, nextSuggestion)
	case 1:
		return w.sendFileQuestion(ctx, req, message, msgChan, nextSuggestion)
	}
	return nil
}

func (w *wsLogic) sendTextQuestion(ctx context.Context, req *dto.SendQuestionReq, message string, msgChan, nextSuggestion chan<- string) error {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a professor of geography. "),
		llms.TextParts(llms.ChatMessageTypeHuman, message),
	}
	response, err := w.langChainGo.LLM.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		msgChan <- string(chunk)
		return nil
	}))
	if err != nil {
		fmt.Println(err)
		return err
	} else if len(response.Choices) == 0 {
		fmt.Println("response is empty")
		return fmt.Errorf("response is empty")
	}
	close(msgChan)
	fmt.Println("close msgChan")

	// 存储消息记录
	go func() {
		w.AddAgentMessageById(int32(req.AgentId), &dto.HistoryMessageResp{
			Text:   message,
			Sender: dto.SenderUser,
		})
		w.AddAgentMessageById(int32(req.AgentId), &dto.HistoryMessageResp{
			Text:   response.Choices[0].Content,
			Sender: dto.SenderAI,
		})
	}()

	// 下一步建议
	err = w.nextSuggest(ctx, messages, response.Choices[0].Content, nextSuggestion)
	if err != nil {
		return err
	}
	return nil
}

func (w *wsLogic) sendFileQuestion(ctx context.Context, req *dto.SendQuestionReq, message string, msgChan, nextSuggestion chan<- string) error {
	// 搜索向量数据
	ret, err := w.useRetriever(w.langChainGo.MilvusStore, message, 10)
	if err != nil {
		fmt.Printf("[sendFileQuestion] useRetriever err,err=%v", err)
		return err
	}
	err = w.getAnswer(ctx, w.langChainGo.LLM, req.AgentId, ret, message, msgChan)
	if err != nil {
		fmt.Printf("[sendFileQuestion] getAnswer err,err=%v", err)
		return err
	}

	return nil
}

// useRetriever 函数使用检索器
func (w *wsLogic) useRetriever(store *milvus.Store, prompt string, topk int) ([]schema.Document, error) {
	// 设置选项向量
	optionsVector := []vectorstores.Option{
		vectorstores.WithScoreThreshold(0.80), // 设置分数阈值
	}

	// 创建检索器
	retriever := vectorstores.ToRetriever(store, topk, optionsVector...)
	// 搜索
	docRetrieved, err := retriever.GetRelevantDocuments(context.Background(), prompt)
	if err != nil {
		return nil, fmt.Errorf("检索文档失败: %v", err)
	}

	// 返回检索到的文档
	return docRetrieved, nil
}

// getAnswer 获取答案
func (w *wsLogic) getAnswer(ctx context.Context, llm llms.Model, agentId int64, docRetrieved []schema.Document, prompt string, msgChan chan<- string) error {
	// 创建一个新的聊天消息历史记录
	history := memory.NewChatMessageHistory()
	// 将检索到的文档添加到历史记录中
	for _, doc := range docRetrieved {
		history.AddAIMessage(ctx, doc.PageContent)
	}
	// 使用历史记录创建一个新的对话缓冲区
	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))
	executor := agents.NewExecutor(
		agents.NewConversationalAgent(llm, nil),
		nil,
		agents.WithMemory(conversation),
	)
	// 设置链调用选项
	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
		chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			msgChan <- string(chunk)
			return nil
		}),
	}
	// 运行链
	response, err := chains.Run(ctx, executor, prompt, options...)
	if err != nil {
		fmt.Println("[chains.Run]", err)
		return err
	}
	close(msgChan)

	w.AddAgentMessageById(int32(agentId), &dto.HistoryMessageResp{
		Text:   prompt,
		Sender: dto.SenderUser,
	})
	w.AddAgentMessageById(int32(agentId), &dto.HistoryMessageResp{
		Text:   response,
		Sender: dto.SenderAI,
	})

	return nil
}

func (w *wsLogic) SaveEarthAgent(ctx context.Context, message string, SaveEarthAgentChan chan<- string, nextSuggestion chan<- string) error {
	// 创建system信息
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, constant.SaveEarthPrompt),
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

	response, err := w.langChainGo.LLM.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		SaveEarthAgentChan <- string(chunk)
		return nil
	}))
	//SaveEarthAgentChan <- "ok"
	if err != nil {
		fmt.Println(err)
		return err
	}
	close(SaveEarthAgentChan)
	fmt.Println("close SaveEarthAgentChan")
	go func() {
		w.AddSaveEarthAgentMessage(&dto.HistoryMessageResp{
			Text:   message,
			Sender: dto.SenderUser,
		})
		w.AddSaveEarthAgentMessage(&dto.HistoryMessageResp{
			Text:   response.Choices[0].Content,
			Sender: dto.SenderAI,
		})
	}()

	// 下一步问题建议
	err = w.nextSuggest(ctx, messages, response.Choices[0].Content, nextSuggestion)
	if err != nil {
		return err
	}
	return nil
}

func (w *wsLogic) nextSuggest(ctx context.Context, messages []llms.MessageContent, response string, nextSuggestion chan<- string) error {
	// 加载上一次的message回复
	messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, response))
	messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, `
		请根据上面的聊天生成3条接下来的聊天内容建议,返回一个string数组,不需要输出任何其他信息

		例如
		["xxxx","xxxx","xxxx"]
	`))
	suggestResp, err := w.langChainGo.LLM.GenerateContent(ctx, messages)
	if err != nil {
		return err
	} else if len(suggestResp.Choices) == 0 {
		return fmt.Errorf("no response")
	}
	nextSuggestion <- suggestResp.Choices[0].Content
	close(nextSuggestion)
	fmt.Println("close nextSuggestion")

	return nil
}
