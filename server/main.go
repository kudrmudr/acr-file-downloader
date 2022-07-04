package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Run server")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Getenv("UPLOAD_PATH")))))
}
