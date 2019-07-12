//初期化とルーティング
package main

import (
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var db *gorm.DB

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

//main()にはルーティングのみ書く
func main() {
	db = ConnectDB()
	defer db.Close()
	user := web.New()
	goji.Handle("/user/*", user)

	user.Use(middleware.SubRouter)
	user.Use(SuperSecure) // ベーシック認証処理追加
	user.Get("/index", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)

	goji.Serve()
}

