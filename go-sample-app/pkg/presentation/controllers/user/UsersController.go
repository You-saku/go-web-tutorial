package controllers

import (
	"fmt"
	"net/http"

	"go-sample-app/pkg/domain/models"
	repository "go-sample-app/pkg/domain/repositories/user"
	service "go-sample-app/pkg/domain/services/user"
)

// GET:users
// POST:users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := service.NewUserService(&repository)

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		users = service.ShowUsers()
		var output = ""
		for _, user := range users {
			output += fmt.Sprintf("id = %d name = %s ", user.Id, user.Name)
		}

		// ステータスコードを設定
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
		return
	}

	// POST
	if requestMethod == "POST" {
		service.CreateUser()
		w.WriteHeader(http.StatusCreated)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
