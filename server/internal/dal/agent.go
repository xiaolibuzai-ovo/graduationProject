package dal

import (
	"context"
	"database/sql"
	"fmt"
)

type Agents struct {
	Id          int32  `json:"id"`
	Img         string `json:"img"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Content     string `json:"content"`
	TextDetail  string `json:"textDetail"`
	Greetings   string `json:"greetings"`
	Prompt      string `json:"prompt"`
	SupportFile int16  `json:"supportFile"`
}

type AgentDal interface {
	GetAllAgents(ctx context.Context) (agents []*Agents, err error)
	GetAgentsById(ctx context.Context, id int32) (agents []*Agents, err error)
}

type agentDal struct {
	db *sql.DB
}

func NewAgentDal(db *sql.DB) AgentDal {
	return &agentDal{db: db}
}

func (a *agentDal) GetAllAgents(ctx context.Context) (agents []*Agents, err error) {
	rows, err := a.db.Query("SELECT * FROM agents")
	if err != nil {
		return
	}
	for rows.Next() {
		var id int32
		var img string
		var title string
		var subtitle string
		var content string
		var textDetail string
		var greetings string
		var prompt string
		var supportFile int16

		err = rows.Scan(&id, &img, &title, &subtitle, &content, &textDetail, &greetings, &prompt, &supportFile)
		if err != nil {
			return
		}
		agents = append(agents, &Agents{
			Id:         id,
			Img:        img,
			Title:      title,
			Subtitle:   subtitle,
			Content:    content,
			TextDetail: textDetail,
			Greetings:  greetings,
		})
	}
	return
}

func (a *agentDal) GetAgentsById(ctx context.Context, agentId int32) (agents []*Agents, err error) {
	rows, err := a.db.Query(fmt.Sprintf("SELECT * FROM agents where id = %d", agentId))
	if err != nil {
		return
	}
	for rows.Next() {
		var (
			id          int32
			img         string
			title       string
			subtitle    string
			content     string
			textDetail  string
			greetings   string
			prompt      string
			supportFile int16
		)
		err = rows.Scan(&id, &img, &title, &subtitle, &content, &textDetail, &greetings, &prompt, &supportFile)
		if err != nil {
			return
		}
		agents = append(agents, &Agents{
			Id:          id,
			Img:         img,
			Title:       title,
			Subtitle:    subtitle,
			Content:     content,
			TextDetail:  textDetail,
			Greetings:   greetings,
			Prompt:      prompt,
			SupportFile: supportFile,
		})
	}
	return
}
