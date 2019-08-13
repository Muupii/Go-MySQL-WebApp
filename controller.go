package main

import "github.com/jinzhu/gorm"

// ConnectDB はDBへの接続
func ConnectDB() *gorm.DB {
	DBMS := "****"
	USER := "****"
	PASS := "****"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "gopractice"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local" // utf8coding対応していなかったため問題が起きてた
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// dbInit DBのマイグレート
func dbInit() {
	db := ConnectDB()
	db.AutoMigrate(&Champion{}) // db.AutoMigrate() はファイルが無ければ生成を行い、すでにファイルがありマイグレートも行われていれば何も行いません
	defer db.Close()
}

// dbInsert DBにデータを追加
func dbInsert(name string, health int, armor int, attackDamage int, attackSpeed int) {
	db := ConnectDB()
	champion := Champion{Name: name, Health: health, Armor: armor, AttackDamage: attackDamage, AttackSpeed: attackSpeed}
	db.Create(&champion)
	defer db.Close()
}

//dbGetAll 全取得
func dbGetAll() []Champion {
	db := ConnectDB()
	champions := []Champion{}
	db.Order("created_at desc").Find(&champions)
	defer db.Close()
	return champions
}

// dbGetOne id指定一つ取得
func dbGetOne(id int) Champion {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	defer db.Close()
	return champion
}

// dbUpdate はid指定でデータ編集
func dbUpdate(id int, name string, health int, armor int, attackDamage int, attackSpeed int) {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	db.Model(&champion).Updates(map[string]interface{}{"Name": name, "Health": health, "Armor": armor, "AttackDamage": attackDamage, "AttackSpeed": attackSpeed})
	defer db.Close()
}

// dbDelete はデータ削除
func dbDelete(id int) {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	db.Delete(&champion)
	defer db.Close()
}
