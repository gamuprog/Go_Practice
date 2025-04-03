package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv" // .env ファイルから環境変数を読み込むためのパッケージ
	_ "github.com/lib/pq"      // PostgreSQL ドライバの読み込み
)

func main() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Println("※ .env ファイルが見つかりませんでした。環境変数をシステムの設定から取得します。")
	}

	// 環境変数から接続情報を取得
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")  // docker-compose で定義したサービス名 "db" を指定することが一般的です
	if dbHost == "" {
		dbHost = "db"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	// PostgreSQL の接続文字列を構築
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// データベースに接続
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("データベース接続エラー: ", err)
	}
	defer db.Close()

	// 接続確認のため Ping を実施
	err = db.Ping()
	if err != nil {
		log.Fatal("データベースに接続できませんでした: ", err)
	}
	fmt.Println("データベースに正常に接続しました！")

	// サンプルとして、現在の時刻をデータベースから取得
	var currentTime time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		log.Fatal("クエリ実行エラー: ", err)
	}
	fmt.Println("DB から取得した現在時刻:", currentTime)
}
