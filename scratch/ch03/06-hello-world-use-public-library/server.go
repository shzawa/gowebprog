package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter" // 外部ライブラリである為 go get が必要
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := "ああああ"

	if p.ByName("name") != "" {
		name = p.ByName("name")
	}

	fmt.Fprintf(w, "hello, %s!\n", name)
}

func main() {
	mux := httprouter.New()

	// httprouterの仕様上、'/hello/:name/'にだけハンドラ関数を紐付けていると、
	// '/hello/'に遷移した場合にNotFoundになる
	mux.GET("/hello/:name/", hello)
	mux.GET("/hello/", hello)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
