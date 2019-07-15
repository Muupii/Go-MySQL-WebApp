package main

import (
	"github.com/wcl48/valval"
	"github.com/zenazn/goji/web"

	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"./db/models"
)

// Password とIDの設定。ベーシック認証用。
const Password = "user:user"

var tpl *template.Template

// FormData はバリデーションエラーを画面に表示するために使うuserモデルとエラーメッセージを持つ構造体
type FormData struct {
	User models.User
	Mess string
}

// Test はテストで使う
type Test struct {
	Users []models.User
	User  models.User
}

// UserIndex はusersテーブルのデータ一覧を出す
func UserIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	Users := []models.User{} // [] はスライスを作るときに使う。[]int{1, 3, 5} だったら[1 3 5]のスライスの作成。
	User := models.User{}
	db.Find(&Users) // SELECT * FROM users;
	User.ID = 1
	db.First(&User)
	tpl = template.Must(template.ParseFiles("view/user/index.tpl"))
	tpl.Execute(w, Test{Users, User})
}

// UserNew はテーブルに新しいデータを登録できるページの表示
func UserNew(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("view/user/new.tpl"))
	tpl.Execute(w, FormData{models.User{}, ""})
}

// UserCreate は新しいusersテーブルに新しいデータをinsertする
func UserCreate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{Name: r.FormValue("Name")}
	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("view/user/new.tpl"))
		tpl.Execute(w, FormData{User, Mess})
	} else {
		db.Create(&User)
		http.Redirect(w, r, "/user/index", 301)
	}
}

// UserEdit はuserテーブルの編集
func UserEdit(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	tpl = template.Must(template.ParseFiles("view/user/edit.tpl"))
	tpl.Execute(w, FormData{User, ""})
}

// UserUpdate は全フィールドの更新
func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	User.Name = r.FormValue("Name")
	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("view/user/edit.tpl"))
		tpl.Execute(w, FormData{User, Mess})
	} else {
		db.Save(&User) // UPDATE users SET name=Name
		http.Redirect(w, r, "/user/index", 301)
	}
}

// UserDelete はテーブルからデータを削除する
func UserDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Delete(&User)
	http.Redirect(w, r, "/user/index", 301)
}

// SuperSecure はベーシック認証する処理
func SuperSecure(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			pleaseAuth(w)
			return
		}

		password, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil || string(password) != Password {
			pleaseAuth(w)
			return
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Authヘッダを受け付けるための処理
func pleaseAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Go away!\n"))
}
