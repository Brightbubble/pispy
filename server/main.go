package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./files/"+r.FormValue("filename"), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "File uploaded successfully.")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed.")
	}
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
