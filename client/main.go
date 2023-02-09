package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("uploadfile", "example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.WriteField("filename", "example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Post("http://localhost:8080/upload", writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	fmt.Println("File uploaded successfully.")
}
