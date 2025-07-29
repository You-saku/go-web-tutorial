package main

import (
	"fmt"
	"log"
	"net/http"
	"web/adapters/controllers"
)

// sample routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang !!")
}

func main() {
	// UserControllerのインスタンスを作成
	userController := controllers.NewUserController()

	router := http.NewServeMux()
	router.HandleFunc("GET /", sampleHandler)
	router.HandleFunc("GET /users", userController.GetAll)
	router.HandleFunc("GET /users/", userController.GetById)
	router.HandleFunc("POST /users", userController.Create)
	router.HandleFunc("PUT /users/", userController.Update)
	router.HandleFunc("DELETE /users/", userController.Delete)

	log.Fatal(http.ListenAndServe(":80", router))
}
