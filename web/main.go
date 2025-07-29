package main

import (
	"fmt"
	"log"
	"net/http"

	userController "web/adapters/controllers/user"
)

// sample routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang !!")
}

func main() {
	// net/http
	http.HandleFunc("/", sampleHandler)

	http.HandleFunc("/users", userController.UsersHandler)
	http.HandleFunc("/users/", userController.UserHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
