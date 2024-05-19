package config

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"server/internal/dto"
)

func InitSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./mock.db")
	if err != nil {
		panic(err)
	}
	//第一次创建数据库表
	err = createAgents(db)
	if err != nil {
		panic(err)
	}
	return db
}

func createAgents(db *sql.DB) error {
	db.Exec("drop table agents")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS agents (
        id INTEGER PRIMARY KEY,
        img VARCHAR,
        title VARCHAR,
        subtitle VARCHAR,
        content VARCHAR,
        textDetail VARCHAR,
	  	greetings VARCHAR,
	  	prompt VARCHAR,
	  	supportFile int 
    )`)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO agents(img, title, subtitle, content, textDetail,greetings,prompt,supportFile) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	for _, agent := range []*dto.AgentListInsertData{
		{
			Img:         "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:       "MSD Mercks Al-Doctor",
			Subtitle:    "@ MattChan",
			Content:     "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:          1,
			TextDetail:  "I'm Gordon Ramsay, taking you on a wild culinary ride. We're uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It's gonna be a mouthwatering experience!",
			Greetings:   "test greetings",
			Prompt:      "You are a professor of geography. ",
			SupportFile: 1,
		},
		{
			Img:         "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:       "MSD Mercks Al-Doctor",
			Subtitle:    "@ MattChan",
			Content:     "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:          2,
			TextDetail:  "I'm Gordon Ramsay, taking you on a wild culinary ride. We're uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It's gonna be a mouthwatering experience!",
			Greetings:   "test greetings",
			Prompt:      "You are a professor of geography. ",
			SupportFile: 0,
		},
		{
			Img:         "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:       "MSD Mercks Al-Doctor",
			Subtitle:    "@ MattChan",
			Content:     "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:          3,
			TextDetail:  "I'm Gordon Ramsay, taking you on a wild culinary ride. We're uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It's gonna be a mouthwatering experience!",
			Greetings:   "test greetings",
			Prompt:      "You are a professor of geography. ",
			SupportFile: 0,
		},
		{
			Img:         "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:       "MSD Mercks Al-Doctor",
			Subtitle:    "@ MattChan",
			Content:     "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:          4,
			TextDetail:  "I'm Gordon Ramsay, taking you on a wild culinary ride. We're uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It's gonna be a mouthwatering experience!",
			Greetings:   "test greetings",
			Prompt:      "You are a professor of geography. ",
			SupportFile: 0,
		},
		{
			Img:         "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:       "MSD Mercks Al-Doctor",
			Subtitle:    "@ MattChan",
			Content:     "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:          5,
			TextDetail:  "I'm Gordon Ramsay, taking you on a wild culinary ride. We're uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It's gonna be a mouthwatering experience!",
			Greetings:   "test greetings",
			Prompt:      "You are a professor of geography. ",
			SupportFile: 0,
		},
	} {
		_, err := stmt.Exec(agent.Img, agent.Title, agent.Subtitle, agent.Content, agent.TextDetail, agent.Greetings, agent.Prompt, agent.SupportFile)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
