package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, r.FormValue("hello"))
	// fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
