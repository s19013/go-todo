package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// これで同じディレクトリ内のschema.sqlが読み込まれるらしい
//
//go:embed schema.sql
var schema string

// fatalで止めるのでerrorを返す必要はない
func NewDB() *sql.DB {
	// envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("環境変数を読み込めませんでした: %v", err)
	}

	// データベースへのパス?
	var dbpath string
	if os.Getenv("DB_CONNECTION") == "sqlite3" {
		dbpath = os.Getenv("DB_PATH")
	} else {
		dbpath = fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	}

	db, err := sql.Open(os.Getenv("DB_CONNECTION"), dbpath)
	if err != nil {
		log.Fatalf("データベースに接続できませんでした: %v", err)
	}

	// スキーマ実行
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("スキーマを実行できませんでした: %v", err)
	}

	return db
}
