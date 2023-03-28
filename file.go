package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle file upload
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Store file data

	fmt.Fprint(w, "File uploaded successfully")
}

func main() {
	http.HandleFunc("/upload", uploadFile)

	fmt.Println("File upload microservice started")
	http.ListenAndServe(":8002", nil)
}
