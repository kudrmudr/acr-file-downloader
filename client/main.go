package main

import (
	"acronis/downloader"
	"acronis/file_links_scraper"
	"log"
	"os"
)

func main() {

	log.Println("Run client")

	fileLinksScraper := file_links_scraper.New(os.Getenv("FILE_SERVER_URL"))
	downloaderService := downloader.New("/usr/share/downloads")

	for _, link := range fileLinksScraper.GetLinks() {
		downloaderService.Download(link)
	}
}
