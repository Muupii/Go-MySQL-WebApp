package models

import (
	"regexp" //正規表現を扱える
	"time"   //時間を扱うパッケージ

	"github.com/wcl48/valval" //validationを扱うフレームワーク的な？
)

// User Userテーブルのモデル
type User struct { //typeとstructについての記事→https://qiita.com/tenntenn/items/45c568d43e950292bc31
	ID        int64
	Name      string    `sql:"size:255"` //この「`」のことをグレイブ・アクセントと呼ぶらしい
	CreatedAt time.Time //time.Time型
	UpdatedAt time.Time
	DeletedAt time.Time
}

// UserValidate バリデーションをする関数
func UserValidate(user User) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(user)
}
