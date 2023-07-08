package controllers

import (
	"fmt"
	"net/http"
	"strings"

	repository "go-architecture-proto/entities/repositories/user"
	usecase "go-architecture-proto/usecases/user"
)

// GET:users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := usecase.NewUserService(&repository)

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		user := usecase.ShowUser(userId)
		fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
		return
	}
	// PUT
	if requestMethod == "PUT" {
		usecase.UpdateUser(userId)
		return
	}
	// DELETE
	if requestMethod == "DELETE" {
		usecase.DeleteUser(userId)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
