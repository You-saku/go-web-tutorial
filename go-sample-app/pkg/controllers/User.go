package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"go-sample-app/pkg/domain/repositories"
)

// ひとまずCRUD作る
func ShowUserHandler(w http.ResponseWriter, r *http.Request) {

	// var user models.User
	userId := strings.TrimPrefix(r.URL.Path, "/users/")
	user := repositories.FindUserById(userId)

	fmt.Fprintf(w, "id = %d name = %s\n", user.Id, user.Name)
}
