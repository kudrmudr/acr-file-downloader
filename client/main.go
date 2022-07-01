package main

import (
	"acronis/downloader"
	"acronis/file_links_scraper"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Client Run")

	//DI
	httpClient := http.Client{}
	mainFileLinksScraper := file_links_scraper.New(httpClient, os.Getenv("FILE_SERVER_URL"))
	mainDownloader := downloader.New(httpClient, "/usr/share/downloads", byte('A'))
	//DI Done

	links := mainFileLinksScraper.GetLinks()
	mainDownloader.Run(links)

	log.Println("Client Done")
}
