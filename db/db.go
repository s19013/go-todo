package db

import (
	"database/sql"
	_ "embed"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// NewDB returns go-sqlite3 driver based *sql.DB.
func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// SQLファイルを読み込む
	// ioutilは古くて非推奨らしいのでosを使う
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// スキーマ実行
	if _, err := db.Exec(string(schema)); err != nil {
		return nil, err
	}

	return db, nil
}
