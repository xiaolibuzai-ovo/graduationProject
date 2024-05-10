package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/milvus"
	"os"
	"strings"
)

var ctx = context.Background()

func main() {
	//answer()
	//addData()
}

func answer() {
	prompt := "第1到2周任务是什么"
	ret, err := useRetriaver(getStore(), prompt, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(ret)
	fmt.Println(string(marshal))
	fmt.Println()
	getAnswer, err := GetAnswer(ctx, getLlm(), ret, prompt)
	if err != nil {
		fmt.Println("[GetAnswer] ", err)
		return
	}
	fmt.Println("ai answer is: ", getAnswer)
}

func addData() {
	err := Embedding(ctx, `
				四、进度计划
		一、（1-2周）确定研究方向和目标
		确定地理信息大模型的研究方向，包括语言理解、地理推理等具体应用方面。
		确立研究目标，明确地理信息大模型的预期功能和应用场景。
		二、（3-4周）文献调研和需求分析
		对大模型技术在地理信息领域的现有研究进行深入调研和分析。
		与相关领域的专家和用户进行需求调研，明确地理信息领域的痛点和应用需求。
		三、（5-6周）模型设计与训练
		基于选定的研究方向和需求分析结果，设计地理信息大模型的结构和算法。
		利用收集的地理信息数据对模型进行训练，并进行性能优化和调整。
		四、（7-8周）地理信息应用平台开发
		开发地理信息应用平台的前端和后端功能，包括用户界面设计、数据接口开发等。
		将训练好的地理信息大模型集成到应用平台中，并进行系统联调和测试。
		五、（9-10周）应用测试与优化
		对开发的地理信息应用平台进行功能测试和性能评估，发现并解决存在的问题和bug。
		根据用户反馈和测试结果，优化地理信息大模型和应用平台的性能和用户体验。
		六、（11-12周）论文撰写和成果总结
		撰写研究论文，总结地理信息大模型的设计与实现过程，以及应用平台的开发与应用。准备项目成果展示和演示材料	
	`)
	if err != nil {
		fmt.Println(err)
	}
}

func Embedding(ctx context.Context, text string) (err error) {
	// 拿到私有数据Embedding
	documents, err := textToChunks(text, 5, 2)
	if err != nil {
		return
	}
	documents = []schema.Document{
		{
			PageContent: text,
		},
	}
	marshal, _ := json.Marshal(documents)
	fmt.Println(string(marshal))

	err = storeDocs(documents, getStore(milvus.WithDropOld()))
	if err != nil {
		return
	}
	return
}

// useRetriaver 函数使用检索器
func useRetriaver(store *milvus.Store, prompt string, topk int) ([]schema.Document, error) {
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

// GetAnswer 获取答案
func GetAnswer(ctx context.Context, llm llms.Model, docRetrieved []schema.Document, prompt string) (string, error) {
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
	}
	// 运行链
	res, err := chains.Run(ctx, executor, prompt, options...)
	if err != nil {
		fmt.Println("[chains.Run]", err)
		return "", err
	}

	return res, nil
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

func getLlm() *openai.LLM {
	llm, err := openai.New(
		openai.WithModel("gpt-3.5-turbo"),
		openai.WithEmbeddingModel("text-embedding-ada-002"),
	)
	if err != nil {
		panic(err)
	}
	return llm
}

func getStore(opts ...milvus.Option) *milvus.Store {
	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
		panic("OPENAI_API_KEY NOT SET")
	}
	embedder, err := embeddings.NewEmbedder(getLlm())
	if err != nil {
		panic(err)
	}

	url := os.Getenv("MILVUS_URL")
	if len(url) == 0 {
		panic("MILVUS_URL NOT SET")
	}
	config := client.Config{
		Address: url,
	}
	idx, err := entity.NewIndexAUTOINDEX(entity.L2)
	if err != nil {
		panic(err)
	}
	//var opts []milvus.Option
	opts = append(
		opts,
		milvus.WithEmbedder(embedder),
		milvus.WithIndex(idx),
		milvus.WithCollectionName("collection1"),
	)
	store, err := milvus.New(
		ctx,
		config,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return &store
}
