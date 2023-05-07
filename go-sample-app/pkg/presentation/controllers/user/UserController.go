package controllers

import (
	"fmt"
	"net/http"
	"strings"

	repository "go-sample-app/pkg/domain/repositories/user"
	service "go-sample-app/pkg/domain/services/user"
)

// GET:users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := service.NewUserService(&repository)

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		user := service.ShowUser(userId)
		fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
		return
	}
	// PUT
	if requestMethod == "PUT" {
		service.UpdateUser(userId)
		return
	}
	// DELETE
	if requestMethod == "DELETE" {
		service.DeleteUser(userId)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
