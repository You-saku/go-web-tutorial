package controllers

import (
	"fmt"
	"go-architecture-proto/entities/models"
	"net/http"

	usersUsecase "go-architecture-proto/usecases/user"
)

// GET:users
// POST:users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		users = usersUsecase.ShowUsers()
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
		usersUsecase.CreateUser()
		w.WriteHeader(http.StatusCreated)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
