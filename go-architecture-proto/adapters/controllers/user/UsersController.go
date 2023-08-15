package controllers

import (
	"encoding/json"
	"go-architecture-proto/entities/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	userRepository "go-architecture-proto/entities/repositories/user"
	userUsecase "go-architecture-proto/usecases/user"
)

var validate *validator.Validate

// バリデーションエラーを格納する構造体
type ErrorMessages struct {
	Errors []Error `json:"errors"`
}

// Userの各フィールドのバリデーションエラーを格納する構造体
type Error struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

// GET:users
// POST:users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// 本番用リポジトリ層
	ur := userRepository.UserRepository{}
	// サービス層作成
	userUsecase := userUsecase.NewUserUsecase(&ur)

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		users = userUsecase.GetUsers()

		// 文字列で返す場合 for debug
		// var output = ""
		// for _, user := range users {
		// 	output += fmt.Sprintf("id = %d name = %s email = %s\n", user.Id, user.Name, user.Email)
		// }

		// ステータスコードを設定
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users) // memo structで定義されてるものが全て返る

		// fmt.Fprintf(w, output) // for debug
		return
	}

	// POST
	if requestMethod == "POST" {
		var user models.User
		validate = validator.New()

		json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード

		err := validate.Struct(user)
		if err != nil {

			var errors []Error
			for _, err := range err.(validator.ValidationErrors) {
				error := Error{
					Field: err.Field(),
					Tag:   err.Tag(),
				}
				errors = append(errors, error)
			}

			// 422エラーを返す
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ErrorMessages{Errors: errors})

			return
		}

		userUsecase.CreateUser(user)
		w.WriteHeader(http.StatusCreated)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
