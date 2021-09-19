package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	// 文字列の配列を取得
	h2 := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h2)

	// 文字列の一覧をカンマ区切りの文字列で取得 (ヘッダ内での実際の文字列)
	h1 := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h1)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", headers)
	server.ListenAndServe()
}
