package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// create an handler using mux
	// second param is inline function
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(res, "Health check") // Fprintln write res on a client.
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Print("Server listing on Port :8080")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
