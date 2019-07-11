//初期化とルーティング
package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	_ "github.com/go-sql-driver/mysql"
	"./db/gormConnect"
)

var db  gormConnect.DB

//main()にはルーティングのみ書く
func main() {
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

//ここには初期化処理専用

func init() {
	// 初期時にdbと接続
	db := gormConnect.ConnectDB()
}