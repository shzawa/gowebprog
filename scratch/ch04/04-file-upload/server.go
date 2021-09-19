package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	strImageData, err := getImageParsedStringData(r.MultipartForm.File["upload_image"][0])
	if err != nil {
		fmt.Println(w, err)
	}
	fmt.Fprintln(w, strImageData)
}

func getImageParsedStringData(fh *multipart.FileHeader) (image string, err error) {
	file, err := fh.Open()
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	image = string(data)
	return
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
