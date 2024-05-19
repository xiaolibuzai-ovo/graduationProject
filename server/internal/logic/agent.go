package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores/milvus"
	"io"
	"net/http"
	"server/internal/constant"
	"server/internal/dal"
	"server/internal/dto"
	"server/internal/ml"
	"strings"
	"sync"
)

type AgentLogic interface {
	AgentList(ctx context.Context) (agents []*dto.AgentListData, err error)
	AgentDetail(ctx context.Context, req *dto.AgentDetailReq) (info *dto.AgentDetailResp, err error)
	Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error)
	Suggests(ctx context.Context, req *dto.SuggestsReq) (data []string, err error)
	LoadPdfData(ctx context.Context, req *dto.LoadPdfReq) (err error)
}

type agentLogic struct {
	dal.AgentDal
	*ml.LangChainGoClient
}

func NewAgentLogic(agent dal.AgentDal, client *ml.LangChainGoClient) AgentLogic {
	return &agentLogic{
		AgentDal:          agent,
		LangChainGoClient: client,
	}
}

func (a *agentLogic) AgentList(ctx context.Context) (agents []*dto.AgentListData, err error) {
	allAgents, err := a.AgentDal.GetAllAgents(ctx)
	if err != nil {
		fmt.Println("[AgentList] GetAllAgents err=", err)
		return
	}
	for _, agent := range allAgents {
		item := &dto.AgentListData{
			Id:       agent.Id,
			Img:      agent.Img,
			Title:    agent.Title,
			Subtitle: agent.Subtitle,
			Content:  agent.Content,
		}
		agents = append(agents, item)
	}
	return
}

func (a *agentLogic) AgentDetail(ctx context.Context, req *dto.AgentDetailReq) (info *dto.AgentDetailResp, err error) {
	agents, err := a.AgentDal.GetAgentsById(ctx, req.AgentId)
	if err != nil {
		return
	}
	if len(agents) == 0 {
		return
	}
	return &dto.AgentDetailResp{
		AgentInfo:   agents[0].TextDetail,
		Greetings:   agents[0].Greetings,
		SupportFile: agents[0].SupportFile,
	}, nil
}

func (a *agentLogic) Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error) {
	// 拿到私有数据Embedding
	documents, err := textToChunks(req.Text)
	if err != nil {
		return
	}
	err = storeDocs(documents, a.MilvusStore)
	if err != nil {
		return
	}
	return
}

// textToChunks 函数将文本文件转换为文档块
func textToChunks(content string) ([]schema.Document, error) {
	reader := strings.NewReader(content)
	// 创建一个新的文本文档加载器
	docLoaded := documentloaders.NewText(reader)
	// 创建一个新的递归字符文本分割器
	split := textsplitter.NewRecursiveCharacter()
	// 加载并分割文档
	docs, err := docLoaded.LoadAndSplit(context.Background(), split)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

// storeDocs 将文档存储到向量数据库
func storeDocs(docs []schema.Document, store *milvus.Store) error {
	// 如果文档数组长度大于0
	if len(docs) > 0 {
		// 添加文档到存储
		_, err := store.AddDocuments(context.Background(), docs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *agentLogic) Suggests(ctx context.Context, req *dto.SuggestsReq) (data []string, err error) {
	// 加载prompt信息
	var prompt string
	if req.AgentId == 99 {
		prompt = constant.SaveEarthPrompt
	} else {
		agents, err := a.AgentDal.GetAgentsById(ctx, req.AgentId)
		if err != nil {
			return nil, err
		}
		if len(agents) == 0 {
			return nil, nil
		}
		prompt = agents[0].Prompt
	}

	// 获取建议
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt),
		llms.TextParts(llms.ChatMessageTypeHuman, `
		结合背景,生成3到4条推荐的谈话问题, 返回一个string数组,不需要输出任何其他信息

		例如
		["你认为造成地球生命值下降的主要原因是什么？","你觉得我们可以采取什么措施来改变这种情况?"]
	`),
	}

	response, err := a.LLM.GenerateContent(ctx, messages)
	if err != nil {
		return
	}
	if response == nil || len(response.Choices) == 0 {
		return
	}
	err = json.Unmarshal([]byte(response.Choices[0].Content), &data)
	if err != nil {
		fmt.Println("[Suggests] json.Unmarshal", err)
		return
	}
	return
}

func (a *agentLogic) LoadPdfData(ctx context.Context, req *dto.LoadPdfReq) (err error) {
	// 加载pdf
	var (
		wg      sync.WaitGroup
		pdfChan = make(chan string, 5)
	)
	for _, file := range req.Files {
		itemFile := file
		wg.Add(1)
		go a.loadFileContent(itemFile, pdfChan, &wg)
	}
	// 消费chan
	go a.consumePdfChan(pdfChan)

	wg.Wait()
	close(pdfChan)

	return
}

func (a *agentLogic) loadFileContent(file string, pdfChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// 加载file数据
	resp, err := http.Get(file)
	if err != nil {
		fmt.Printf("[sendFileQuestion] http load file err,err=%v", err)
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[sendFileQuestion] io.ReadAll err,err=%v", err)
		return
	}
	doc, err := fitz.NewFromMemory(bytes)
	if err != nil {
		fmt.Printf("[sendFileQuestion] NewFromMemory err,err=%v", err)
		return
	}

	defer doc.Close()
	var pdfText strings.Builder
	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			continue
		}
		pdfText.WriteString(text)
	}
	pdfChan <- pdfText.String()
}

func (a *agentLogic) consumePdfChan(pdfChan <-chan string) {
	var str strings.Builder
	for pdf := range pdfChan {
		str.WriteString(pdf)
		str.WriteString("\n")
	}

	documents, err := textToChunks(str.String())
	if err != nil {
		fmt.Printf("[consumePdfChan] textToChunks err,err=%v", err)
		return
	}
	err = storeDocs(documents, a.MilvusStore)
	if err != nil {
		fmt.Printf("[consumePdfChan] storeDocs err,err=%v", err)
		return
	}
}
