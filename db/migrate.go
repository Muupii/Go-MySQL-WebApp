package main

import (
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"./models" //自分で作ったパッケージのインポート
	"./gormConnect"
)




func main(){
	db := gormConnect.ConnectDB()
	db.CreateTable(&models.User{}) //&は変数のアドレスを取得している
	defer db.Close()	
}