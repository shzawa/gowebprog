package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "Cookie1",
		Value:    "Value1",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "Cookie2",
		Value:    "Value2",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("Cookie1")
	if err != nil {
		fmt.Fprintln(w, "Cookie1 not found")
	}
	c2 := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, c2)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/get", getCookie)
	server.ListenAndServe()
}
