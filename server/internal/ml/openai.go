package ml

import (
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores/milvus"
	"golang.org/x/net/context"
	"os"
)

type LangChainGoClient struct {
	Ctx            context.Context
	LLM            *openai.LLM
	OpenaiEmbedder embeddings.Embedder
	MilvusStore    milvus.Store
}

func NewLangChainGoClient(ctx context.Context) *LangChainGoClient {
	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
		panic("OPENAI_API_KEY NOT SET")
	}

	opts := []openai.Option{
		openai.WithModel("gpt-3.5-turbo-0125"),
		openai.WithEmbeddingModel("text-embedding-ada-002"),
	}

	llm, err := openai.New(opts...)
	if err != nil {
		panic(err)
	}
	c := &LangChainGoClient{
		Ctx: ctx,
		LLM: llm,
	}

	c.InitEmbedder()
	c.initMilvusStore()

	return c
}

func (c *LangChainGoClient) InitEmbedder() {
	embedder, err := embeddings.NewEmbedder(c.LLM)
	if err != nil {
		panic(err)
	}
	c.OpenaiEmbedder = embedder
}

func (c *LangChainGoClient) initMilvusStore() {
	e := c.OpenaiEmbedder

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
	var opts []milvus.Option
	opts = append(
		opts,
		milvus.WithEmbedder(e),
		milvus.WithIndex(idx),
		milvus.WithCollectionName("collection"),
	)
	store, err := milvus.New(
		c.Ctx,
		config,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	c.MilvusStore = store
}
