package main

import (
	"encoding/json"
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

func Hello(out http.ResponseWriter, in *http.Request) {
	// set cors
	out.Header().Set("Content-Type", "application/json")
	out.Header().Set("Access-Control-Allow-Headers", "*")
	out.Header().Set("Access-Control-Allow-Origin", "*")
	out.WriteHeader(http.StatusOK)

	var response = make(map[string]interface{})
	response["status"] = http.StatusOK
	response["data"] = "Hello World"
	_ = json.NewEncoder(out).Encode(response)
}
