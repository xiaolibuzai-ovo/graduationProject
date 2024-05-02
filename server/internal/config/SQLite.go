package config

import (
	"database/sql"
)

func InitSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./mock.db")
	if err != nil {
		panic(err)
	}
	// 创建数据库表
	err = createTestDb(db)
	if err != nil {
		panic(err)
	}
	return db
}
func createTestDb(db *sql.DB) error {
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
