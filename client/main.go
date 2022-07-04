package main

import (
	"acronis/downloader"
	"acronis/file_links_scraper"
	"acronis/infrastructure/filesystem"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Client Run")

	//DI
	httpClient := &http.Client{}
	fs := filesystem.New("/usr/share/downloads")
	mainFileLinksScraper := file_links_scraper.New(httpClient, os.Getenv("FILE_SERVER_URL"))
	mainDownloader := downloader.New(httpClient, fs, byte('A'))
	//DI Done

	links := mainFileLinksScraper.GetLinks()
	mainDownloader.Run(links)

	log.Println("Client Done")
}
