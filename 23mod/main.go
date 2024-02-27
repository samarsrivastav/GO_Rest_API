package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") //calling the  endpoints

	log.Fatal(http.ListenAndServe(":4000", r))
}
func greet() {
	fmt.Println("hello")
}

// making endpoints
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome  to home server </h1>"))
}
