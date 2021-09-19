package main

import (
	"fmt"
	"net/http"
)

// ハンドラ関数   メソッドServeHTTPと同じシグネチャを持つ関数
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/hello", hello)

	// ↑ のHandleFuncが内部で行っていることを具体化したもの
	worldHandler := http.HandlerFunc(world) // HandlerFuncは、メソッドではなく型
	http.Handle("/world", &worldHandler)

	server.ListenAndServe()
}
