package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
		<head>
			<title>Go Web Programming</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
		</body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "そのようなサービスはありません。ほかを当たってください。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(http.StatusFound) // WriteHeaderは呼び出された後にヘッダの変更を防ぐ為、最後に呼び出す必要がある
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}
