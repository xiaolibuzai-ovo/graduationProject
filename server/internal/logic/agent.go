package logic

import (
	"context"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores/milvus"
	"server/internal/dal"
	"server/internal/dto"
	"server/internal/ml"
	"strings"
)

type AgentLogic interface {
	AgentList(ctx context.Context) (agents []*dto.AgentListData, err error)
	AgentDetail(ctx context.Context, req *dto.AgentDetailReq) (info *dto.AgentDetailResp, err error)
	Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error)
	Suggests(ctx context.Context, req *dto.SuggestsReq) (data []string, err error)
}

type agentLogic struct {
	dal.AgentDal
	*ml.LangChainGoClient
}

func NewAgentLogic(agent dal.AgentDal) AgentLogic {
	return &agentLogic{AgentDal: agent}
}

func (a *agentLogic) AgentList(ctx context.Context) (agents []*dto.AgentListData, err error) {
	allAgents, err := a.AgentDal.GetAllAgents(ctx)
	if err != nil {
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
		AgentInfo: agents[0].TextDetail,
		Greetings: agents[0].Greetings,
	}, nil
}

func (a *agentLogic) Embedding(ctx context.Context, req *dto.EmbeddingReq) (err error) {
	// 拿到私有数据Embedding
	documents, err := textToChunks(req.Text, req.ChunkSize, req.ChunkOverlap)
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
func textToChunks(content string, chunkSize, chunkOverlap int) ([]schema.Document, error) {
	reader := strings.NewReader(content)
	// 创建一个新的文本文档加载器
	docLoaded := documentloaders.NewText(reader)
	// 创建一个新的递归字符文本分割器
	split := textsplitter.NewRecursiveCharacter()
	// 设置块大小
	split.ChunkSize = chunkSize
	// 设置块重叠大小
	split.ChunkOverlap = chunkOverlap
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

	// 获取建议

	return
}
