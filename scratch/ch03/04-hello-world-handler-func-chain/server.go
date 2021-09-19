package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// ハンドル関数 = HandlerFunc型
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(next).Pointer()).Name()
		fmt.Println("log Handler function called - " + name)
		next(w, r)
	}
}

func protect(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("protect Handler called - %T\n", next)
		next(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/hello", protect(log(hello)))
	// -> http.HandleFunc(pattern = "/hello", handler = log(hello))
	// -> http.HandleFunc(pattern = "/hello", handler = hello)

	server.ListenAndServe()
}
