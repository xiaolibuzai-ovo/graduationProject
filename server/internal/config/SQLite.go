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
			Img:         "https://lf16-alice-tos-sign.oceanapi-i18n.com/obj/ocean-cloud-tos-sg/FileBizType.BIZ_BOT_ICON/7348600504488576002_1710982299024767106.jpeg?lk3s=50ccb0c5&x-expires=1716206956&x-signature=B%2BrwlGe5Vf3WJKOHYalE8A5qimA%3D",
			Title:       "自然灾害管理专家",
			Subtitle:    "@ MattChan",
			Content:     "自然灾害管理专家小助手帮助研究和分析地震、洪水、飓风等自然灾害的预防和应对策略，为减少灾害影响提供专业见解。",
			Id:          1,
			TextDetail:  "自然灾害管理专家小助手将带你深入了解自然灾害的预防和应对策略。这个小助手会分析各种自然灾害的特征，探讨有效的防灾减灾方法，并提供应急管理建议。加入这个小助手，提升你的灾害管理能力。",
			Greetings:   "欢迎使用自然灾害管理专家小助手！",
			Prompt:      "你是一位自然灾害管理教授。",
			SupportFile: 0,
		},
		{
			Img:         "https://p16-flow-product-sign-sg.ibyteimg.com/tos-alisg-i-bfte7mpw5s-sg/bf46c838bc9e4b29b098b2d1e0b75883~tplv-bfte7mpw5s-resize:128:128.image?rk3s=2e2596fd&x-expires=1718715913&x-signature=MtayMP%2FuvxttohdSuGF26kcB9wA%3D",
			Title:       "旅行小助手",
			Subtitle:    "@ MattChan",
			Content:     "地图专家提供精准的地理位置信息和路线规划，让您轻松找到目的地的所以情况，可以为您提前量身定做游玩攻略,当地美食",
			Id:          2,
			TextDetail:  "地图专家提供精准的地理位置信息和路线规划，让您轻松找到目的地的所以情况，可以为您提前量身定做游玩攻略,当地美食",
			Greetings:   "你好，我是旅游专家，我可以帮你轻松制作任何一个城市的旅游攻略和前往方案",
			Prompt:      "你是旅游专家，可以帮我轻松制作任何一个城市的旅游攻略和前往方案",
			SupportFile: 0,
		},
		{
			Img:         "https://lf16-alice-tos-sign.oceanapi-i18n.com/obj/ocean-cloud-tos-sg/FileBizType.BIZ_BOT_ICON/7331559207001752578_1708493843729596758.jpeg?lk3s=50ccb0c5&x-expires=1716210514&x-signature=mqaGyTj%2BU%2FsLIHqXCx56hhWxzBA%3D",
			Title:       "GIS小助手",
			Subtitle:    "@ MattChan",
			Content:     "GIS是一种强大的工具，用于捕获、存储、分析和管理空间数据。无论你是城市规划者、环境科学家，还是地理数据的学习者，GIS都能为你提供宝贵的见解和解决方案。",
			Id:          3,
			TextDetail:  "GIS (地理信息系统) 是用于捕获、存储、分析和管理空间和地理数据的强大工具。无论你是城市规划者、环境科学家，还是希望更好地了解地理空间数据的学生，GIS都能为你提供宝贵的见解和解决方案。",
			Greetings:   "欢迎来到GIS小助手！ 有什么问题我可以帮你解答呢",
			Prompt:      "你是一名精通地理信息(GIS)的大师，回答我提出的地理信息相关问题",
			SupportFile: 0,
		},
		{
			Img:         "https://p19-flow-product-sign-sg.ibyteimg.com/tos-alisg-i-bfte7mpw5s-sg/26d478507e9847c396eed7e1bcecbea7~tplv-bfte7mpw5s-resize:128:128.image?rk3s=2e2596fd&x-expires=1718716855&x-signature=Qhmhmannw8K7CjfhqG9kiWSamUU%3D",
			Title:       "气候分析小助手",
			Subtitle:    "@ MattChan",
			Content:     "气候分析小助手可以帮助你分析各种气候类型及其对环境和人类的影响，无论你是气候科学家还是学生，它都能提供有价值的见解。",
			Id:          4,
			TextDetail:  "气候分析小助手可以帮助你分析全球各种气候类型，并了解它们对环境和人类活动的影响。无论你是研究气候变化的科学家，还是对气候影响感兴趣的学生，这个小助手都能为你提供有价值的分析和见解。",
			Greetings:   "欢迎使用气候分析小助手！",
			Prompt:      "你是一位气候学教授。",
			SupportFile: 0,
		},
		{
			Img:         "https://lf16-alice-tos-sign.oceanapi-i18n.com/obj/ocean-cloud-tos-sg/FileBizType.BIZ_BOT_ICON/7323631676981642241_1708106515773069442.jpeg?lk3s=50ccb0c5&x-expires=1716211228&x-signature=4WvkA6hksyKGPb%2BdmqbF%2FXfGAb8%3D",
			Title:       "遥感技术专家",
			Subtitle:    "@ MattChan",
			Content:     "遥感技术专家小助手帮助理解和应用遥感技术进行地理数据采集和分析，为环境科学家、城市规划者和学生提供专业见解和分析。",
			Id:          5,
			TextDetail:  "遥感技术专家小助手将带你探索遥感技术的世界。我会介绍各种遥感技术，展示如何使用这些技术进行环境监测、土地利用分析和资源管理, 和我开启一段科技与地理交汇的探索之旅吧！",
			Greetings:   "我是你的遥感技术专家小助手，带你探索遥感技术的世界。我们将介绍各种遥感技术，展示如何使用这些技术进行环境监测、土地利用分析和资源管理。加入我，开启这段科技与地理交汇的探索之旅吧！",
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
