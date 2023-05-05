package main

import (
	"fmt"
	"log"
	"net/http"

	"go-sample-app/pkg/controllers"
)

// routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	// net/http
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/users/", controllers.ShowUserHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
