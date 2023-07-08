package controllers

import (
	"fmt"
	"net/http"
	"strings"

	userUsecases "go-architecture-proto/usecases/user"
)

// GET:users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		user := userUsecases.ShowUser(userId)
		fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
		return
	}
	// PUT
	if requestMethod == "PUT" {
		userUsecases.UpdateUser(userId)
		return
	}
	// DELETE
	if requestMethod == "DELETE" {
		userUsecases.DeleteUser(userId)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
