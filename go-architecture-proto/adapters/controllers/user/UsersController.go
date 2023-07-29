package controllers

import (
	"encoding/json"
	"fmt"
	"go-architecture-proto/entities/models"
	"net/http"

	userRepository "go-architecture-proto/entities/repositories/user"
	userUsecase "go-architecture-proto/usecases/user"
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
		var output = ""
		for _, user := range users {
			output += fmt.Sprintf("id = %d name = %s email = %s\n", user.Id, user.Name, user.Email.String)
		}

		// ステータスコードを設定
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
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
