package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}
