package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(hello http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("log Handler called - %T\n", hello)
		hello.ServeHTTP(w, r)
	})
}

func protect(log http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("protect Handler called - %T\n", log)
		log.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	// http.Handle("/hello/", protect(log(hello))) // 末尾が'/'なので '/hello/hoge'にアクセスするとここに遷移される
	server.ListenAndServe()
}
