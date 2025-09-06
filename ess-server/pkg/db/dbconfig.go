package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("DB 環境変数が設定されていません")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable application_name=%s",
		host, port, user, password, dbname, "ess-server")

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("データベース接続に失敗しました:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("データベースが起動していません:", err)
	}

	fmt.Println("データベース接続成功")
}
