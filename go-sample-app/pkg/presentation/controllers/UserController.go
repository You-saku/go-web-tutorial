package controllers

import (
	"fmt"
	"net/http"
	"strings"

	Repository "go-sample-app/pkg/domain/repositories/user"
	"go-sample-app/pkg/domain/services"
)

// ひとまずCRUD作る
func ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	// var user models.User
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	repository := Repository.UserRepository{}
	service := services.UserService{
		Rep: repository,
	}

	user := service.ShowUser(userId)

	fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
}
