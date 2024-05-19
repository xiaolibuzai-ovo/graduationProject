package dto

type AgentListResp struct {
	Agents []*AgentListData `json:"agents,omitempty"`
}

type AgentListInsertData struct {
	Id          int32  `json:"id"`
	Img         string `json:"img"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Content     string `json:"content"`
	TextDetail  string `json:"textDetail"`
	Greetings   string `json:"greetings"`
	Prompt      string `json:"prompt"`
	SupportFile int16  `json:"SupportFile"`
}

type AgentListData struct {
	Id         int32  `json:"id"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Content    string `json:"content"`
	TextDetail string `json:"textDetail"`
}

type AgentDetailReq struct {
	AgentId int32 `json:"agentId"`
}

type AgentDetailResp struct {
	AgentInfo   string `json:"agentInfo"`
	Greetings   string `json:"greetings"`
	SupportFile int16  `json:"supportFile"`
}

type EmbeddingReq struct {
	Text         string `json:"text"`
	ChunkSize    int    `json:"chunkSize"`
	ChunkOverlap int    `json:"chunkOverlap"`
}

type HistoryMessageReq struct {
	AgentId int32 `json:"agentId"`
}

type DeleteHistoryMessageReq struct {
	AgentId int32 `json:"agentId"`
}

var (
	SenderAI   = "bot"
	SenderUser = "user"
)

type HistoryMessageResp struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

type SuggestsReq struct {
	AgentId int32 `json:"agentId"`
}

type SuggestsResp struct {
	SuggestsData []string `json:"suggestsData"`
}

type SendQuestionReq struct {
	AgentId int64 `json:"agentId"`
}

type SendQuestionMessage struct {
	Content string `json:"content"`
}

type LoadPdfReq struct {
	Files []string `json:"files"`
}
