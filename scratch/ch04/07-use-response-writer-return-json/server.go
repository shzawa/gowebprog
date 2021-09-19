package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Shotaro Ozawa",
		Threads: []string{"1番目", "2番目", "3番目"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/", jsonExample)
	server.ListenAndServe()
}
