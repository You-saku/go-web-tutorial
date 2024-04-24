package controllers

import (
	"api/entities/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	userRepository "api/entities/repositories/user"
	userUsecase "api/usecases/user"
)

// GET:users/{id}
// PUT:users/{id}
// DELETE:users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// 本番用リポジトリ層
	ur := userRepository.UserRepository{}
	// サービス層作成
	userUsecase := userUsecase.NewUserUsecase(&ur)

	userId := strings.TrimPrefix(r.URL.Path, "/users/")
	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

	// GET
	if requestMethod == "GET" {
		user := userUsecase.GetUser(userId)

		json.NewEncoder(w).Encode(user)
		//fmt.Fprintf(w, "id = %d name = %s email = %s\n", user.Id, user.Name, user.Email) // for debug
		return
	}
	// PUT
	if requestMethod == "PUT" {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード
		fmt.Fprintf(w, "name = %s\n", user.Name)
		userUsecase.UpdateUser(userId, user)
		return
	}
	// DELETE
	if requestMethod == "DELETE" {
		userUsecase.DeleteUser(userId)
		return
	}

	// それ以外のメソッドはエラー
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}
