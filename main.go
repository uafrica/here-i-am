package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	ct := "text/plain"
	out := "Here I am!"

	if r.Header.Get("Content-Type") == "application/json" {
		ct = "application/json"
		out = `{"status": "Here I am!"}`
	}

	w.Header().Set("Content-Type", ct)
	io.WriteString(w, out)
}

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
