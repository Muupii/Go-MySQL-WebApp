package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"./models" //自分で作ったパッケージのインポート
)

  //mysqlへの接続GolangのORMであるgormを使用。詳しくは→　https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81
  func gormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "masayuki"
	PASS     := "aaaa"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "gopractice"
  
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)
  
	if err != nil {
	  panic(err.Error())
	}
	return db
  }


func main(){
	db := gormConnect()
	db.CreateTable(&models.User{})
	defer db.Close()	
}