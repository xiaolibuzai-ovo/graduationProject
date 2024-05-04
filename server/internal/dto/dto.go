package dto

type AgentListResp struct {
	Agents []*AgentListData `json:"agents,omitempty"`
}

type AgentListData struct {
	Id       int32  `json:"id"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}

type EmbeddingReq struct {
	Text string `json:"text"`
}
