package main

import (
	"fmt"
	"net/http"
)

// MyHandlerはハンドラと定義されている。
// ハンドラ   ServeHTTP(w http.ResponseWriter, r *http.Request) というシグネチャを持つインターフェスのこと
// シグネチャ   例: Javaなどの言語は、シグネチャを用いて同じ名前のメソッドだが引数の型などが異なるものを複数同時に宣言できる仕組みを提供している
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
