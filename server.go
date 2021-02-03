package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// var router = mux.NewRouter();
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/posts", getposts).Methods("GET")
	log.Println("Server Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
