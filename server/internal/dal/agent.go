package dal

import (
	"context"
	"database/sql"
)

type Agents struct {
	Id       int32  `json:"id"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}

type AgentDal interface {
	GetAllAgents(ctx context.Context) (agents []*Agents, err error)
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
		err = rows.Scan(&id, &img, &title, &subtitle, &content)
		if err != nil {
			return
		}
		agents = append(agents, &Agents{
			Id:       id,
			Img:      img,
			Title:    title,
			Subtitle: subtitle,
			Content:  content,
		})
	}
	return
}
