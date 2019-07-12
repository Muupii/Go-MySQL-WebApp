package main

import (
	"./models"                         //自分で作ったパッケージのインポート
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"github.com/jinzhu/gorm"
)

// ConnectDB はgormを使ってデータベースに接続します
func ConnectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "masayuki"
	PASS := "aaaa"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "gopractice"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := ConnectDB()
	db.CreateTable(&models.User{}) //&は変数のアドレスを取得している
	defer db.Close()
}
