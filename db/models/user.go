package models

import(
	"github.com/wcl48/valval" //validationを扱うフレームワーク的な？
	"time" //時間を扱うパッケージ
	"regexp" //正規表現を扱える
)

type User struct { //typeとstructについての記事→https://qiita.com/tenntenn/items/45c568d43e950292bc31
	Id int64
	Name string `sql:"size:255"` //この「`」のことをグレイブ・アクセントと呼ぶらしい
	CreatedAt time.Time //time.Time型
	UpdatedAt time.Time
	DeletedAt time.Time
}

func UserValidate(user User)(error) {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)),
		),
	})

	return Validator.Validate(user)
}