package controllers

import (
	"fmt"
	"net/http"

	"go-sample-app/pkg/domain/models"
	repository "go-sample-app/pkg/domain/repositories/user"
	service "go-sample-app/pkg/domain/services/user"
)

// GET:users
func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// 本番用リポジトリ層
	repository := repository.UserRepository{}
	// サービス層作成
	service := service.NewUserService(&repository)
	users = service.ShowUsers()

	var output = ""
	for _, user := range users {
		output += fmt.Sprintf("id = %d name = %s ", user.Id, user.Name)
	}

	fmt.Fprintf(w, output)
}
