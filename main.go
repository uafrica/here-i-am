package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var port string

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Serving HTTP connections on port %s\n", port)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
