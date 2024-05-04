package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"server/internal/dto"
)

func InitSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./mock.db")
	if err != nil {
		panic(err)
	}
	// 第一次创建数据库表
	//err = createAgents(db)
	//if err != nil {
	//	panic(err)
	//}
	return db
}

func createAgents(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS agents (
        id INTEGER PRIMARY KEY,
        img VARCHAR,
        title VARCHAR,
        subtitle VARCHAR,
        content VARCHAR
    )`)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO agents(img, title, subtitle, content) values(?,?,?,?)")
	if err != nil {
		return err
	}
	for _, agent := range []*dto.AgentListData{
		{
			Img:      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:    "MSD Mercks Al-Doctor",
			Subtitle: "@ MattChan",
			Content:  "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:       1,
		},
		{
			Img:      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:    "MSD Mercks Al-Doctor",
			Subtitle: "@ MattChan",
			Content:  "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:       2,
		},
		{
			Img:      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:    "MSD Mercks Al-Doctor",
			Subtitle: "@ MattChan",
			Content:  "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:       3,
		},
		{
			Img:      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:    "MSD Mercks Al-Doctor",
			Subtitle: "@ MattChan",
			Content:  "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:       4,
		},
		{
			Img:      "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
			Title:    "MSD Mercks Al-Doctor",
			Subtitle: "@ MattChan",
			Content:  "MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..",
			Id:       5,
		},
	} {
		stmt.Exec(agent.Img, agent.Title, agent.Subtitle, agent.Content)
	}
	return nil
}
