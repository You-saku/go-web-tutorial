package controllers

import (
	"fmt"
	"net/http"
	"strings"

	repository "go-sample-app/pkg/domain/repositories/user"
	service "go-sample-app/pkg/domain/services/user"
)

// GET:users/{id}
func UserShowHandler(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := service.NewUserService(&repository)

	user := service.ShowUser(userId)
	fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
}
