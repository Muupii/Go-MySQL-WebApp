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
	Name         string
	Health       int
	Armor        int
	AttackDamage int
	AttackSpeed  int
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
		name := ctx.PostForm("name")
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

		dbInsert(name, health, armor, attackDamage, attackSpeed)
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
		name := ctx.PostForm("name")
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

		dbUpdate(id, name, health, armor, attackDamage, attackSpeed)
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
