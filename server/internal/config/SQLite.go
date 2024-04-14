package config

import (
	"database/sql"
)

var (
	db *sql.DB
)

func GetDb() *sql.DB {
	return db
}

func InitSQLiteDB() {
	var err error
	db, err = sql.Open("sqlite3", "./mock.db")
	if err != nil {
		panic(err)
	}
	// 创建数据库表
	err = createTestDb(db)
	if err != nil {
		panic(err)
	}
}

func createTestDb(db *sql.DB) error {
	// 创建订单表
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY,
        uid INTEGER,
        weight DOUBLE,
        created_at DATETIME
    )`)
	if err != nil {
		return err
	}
	return nil
}
