// deprecated: This file is now replaced by web/adapters/controllers/UserController.go
package controllers

import (
	"web/entities/models"
	"encoding/json"
	"net/http"

	userRepository "web/entities/repositories/user"
	userUsecase "web/usecases/user"
)

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
		json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード

		userUsecase.CreateUser(user)
		w.WriteHeader(http.StatusCreated)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
