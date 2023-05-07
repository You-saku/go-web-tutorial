package main

import (
	"fmt"
	"log"
	"net/http"

	userController "go-sample-app/pkg/presentation/controllers/user"
)

// routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	// net/http
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/users", userController.UserGetHandler)
	http.HandleFunc("/users/", userController.UserShowHandler)

	/**
	TODO: postやput, deleteをどうするのか
	requestMethod := r.Method // これでhttpリクエストのメソッドを取得
	if requestMethod != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	**/

	log.Fatal(http.ListenAndServe(":80", nil))
}
