package main

import (
	"acronis/downloader"
	"acronis/file_links_scraper"
	"log"
	"os"
)

func main() {
	log.Println("Client Run")

	mainFileLinksScraper := file_links_scraper.New(os.Getenv("FILE_SERVER_URL"))
	mainDownloader := downloader.New("/usr/share/downloads", byte('A'))

	links := mainFileLinksScraper.GetLinks()
	mainDownloader.Run(links)

	log.Println("Client Done")
}
