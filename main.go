package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println(fmt.Sprintf(`"Starting up on %s"`, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%s`, port), nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
