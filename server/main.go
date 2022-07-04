package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Run server")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/files"))))
}
