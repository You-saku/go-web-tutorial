package main

import (
	"fmt"
	"log"
	"net/http"

	todoController "go-sample-app/pkg/presentation/controllers/todo"
	userController "go-sample-app/pkg/presentation/controllers/user"
)

// routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	// net/http
	http.HandleFunc("/", sampleHandler)

	// user関連
	http.HandleFunc("/users", userController.UsersHandler)
	http.HandleFunc("/users/", userController.UserHandler)

	// todo関連
	http.HandleFunc("/todos", todoController.TodoCreateHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
