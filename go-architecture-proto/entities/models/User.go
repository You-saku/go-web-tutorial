package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// validation設定
type UserCustomValidator struct{}

func (cv *UserCustomValidator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("名前は必須入力です"),
			validation.RuneLength(5, 20).Error("名前は 5～20 文字です"),
		),
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスは必須入力です"),
			is.Email.Error("メールアドレスを入力してください"),
		),
		validation.Field(
			&u.Age,
			validation.Required.Error("年齢は必須項目です。"),
			// TODO:数値型をバリデーションしたい
		),
	)
}
