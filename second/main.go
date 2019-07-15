// GoでのWebApp二度目の挑戦。参考URL→https://qiita.com/Anharu/items/ce644c521a4d52fafb7e & https://qiita.com/hyo_07/items/59c093dda143325b1859
package main

import (
	"strconv"

	"github.com/gin-gonic/gin" // gin はGo製のフレームワーク。軽いらしい。

	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"github.com/jinzhu/gorm"
)

// Champion Championを定義
type Champion struct {
	gorm.Model   // gorm.Model はgormの標準モデルでid, created_at, updated_at, deleted_atで構成されている
	Health       int
	Armor        int
	AttackDamage int
	AttackSpeed  int
}

// ConnectDB はDBへの接続
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

// dbInit DBのマイグレート
func dbInit() {
	db := ConnectDB()
	db.AutoMigrate(&Champion{}) // db.AutoMigrate() はファイルが無ければ生成を行い、すでにファイルがありマイグレートも行われていれば何も行いません
	defer db.Close()
}

// dbInsert DBにデータを追加
func dbInsert(health int, armor int, attackDamage int, attackSpeed int) {
	db := ConnectDB()
	champion := Champion{Health: health, Armor: armor, AttackDamage: attackDamage, AttackSpeed: attackSpeed}
	db.Create(&champion)
	defer db.Close()
}

//dbGetAll 全取得
func dbGetAll() []Champion {
	db := ConnectDB()
	champions := []Champion{}
	db.Order("created_at desc").Find(&champions)
	db.Close()
	return champions
}

// dbGetOne id指定一つ取得
func dbGetOne(id int) Champion {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	db.Close()
	return champion
}

// dbUpdate はid指定でデータ編集
func dbUpdate(id int, health int, armor int, attackDamage int, attackSpeed int) {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	champion.Health = health
	champion.Armor = armor
	champion.AttackDamage = attackDamage
	champion.AttackSpeed = attackSpeed
	db.Save(&champion)
	db.Close()
}

// dbDelete はデータ削除
func dbDelete(id int) {
	db := ConnectDB()
	champion := Champion{}
	db.First(&champion, id)
	db.Delete(&champion)
	db.Close()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html") // HTML Templatesを読み込む

	dbInit()

	//Index
	router.GET("/", func(ctx *gin.Context) {
		champions := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"champions": champions,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		healthString := ctx.PostForm("health")
		armorString := ctx.PostForm("armor")
		attackDamageString := ctx.PostForm("attackDamage")
		attackSpeedString := ctx.PostForm("attackSpeed")
		abilities := []string{healthString, armorString, attackDamageString, attackSpeedString}
		abilitiesName := []string{"health", "armor", "attackDamage", "attackSpeed"}
		abilitiesMap := map[string]int{}
		for i := 0; i < len(abilities); i++ {
			abilityInt, err := strconv.Atoi(abilities[i])
			if err != nil {
				panic(err)
			}
			abilitiesMap[abilitiesName[i]] = abilityInt
		}
		health := abilitiesMap["health"]
		armor := abilitiesMap["armor"]
		attackDamage := abilitiesMap["attackDamage"]
		attackSpeed := abilitiesMap["attackSpeed"]

		dbInsert(health, armor, attackDamage, attackSpeed)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n) // int型に変換
		if err != nil {
			panic(err)
		}
		champion := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"champion": champion})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}

		healthString := ctx.PostForm("health")
		armorString := ctx.PostForm("armor")
		attackDamageString := ctx.PostForm("attackDamage")
		attackSpeedString := ctx.PostForm("attackSpeed")
		abilities := []string{healthString, armorString, attackDamageString, attackSpeedString}
		abilitiesName := []string{"health", "armor", "attackDamage", "attackSpeed"}
		abilitiesMap := map[string]int{}
		for i := 0; i < len(abilities)-1; i++ {
			abilityInt, err := strconv.Atoi(abilities[i])
			if err != nil {
				panic(err)
			}
			abilitiesMap[abilitiesName[i]] = abilityInt
		}
		health := abilitiesMap["health"]
		armor := abilitiesMap["armor"]
		attackDamage := abilitiesMap["attackDamage"]
		attackSpeed := abilitiesMap["attackSpeed"]

		dbUpdate(id, health, armor, attackDamage, attackSpeed)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		champion := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"champion": champion})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")

	})

	router.Run()
}
