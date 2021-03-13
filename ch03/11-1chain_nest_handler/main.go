package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareOne を実行")
		next.ServeHTTP(w, r)
		log.Println("middlewareOne をもう一度実行")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareTwo を実行")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("middlewareTwo をもう一度実行")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("finalHandler を実行")
	w.Write([]byte("OK"))
}

func main() {
	finalHandler := http.HandlerFunc(final)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	http.ListenAndServe(":3000", nil)
}
