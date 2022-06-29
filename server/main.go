package main

import (
	"log"
	"net/http"
)

func main() {
	println("Run server")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/files"))))
}
