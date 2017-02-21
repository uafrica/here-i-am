package main

import (
	"io"
	"net/http"
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
