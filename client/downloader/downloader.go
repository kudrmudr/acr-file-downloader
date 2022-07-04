package downloader

import (
	"acronis/infrastructure/filesystem"
	"acronis/infrastructure/http_client"
	"log"
	"math"
	"sync"
)

func New(httpClient http_client.Client, fs filesystem.Filesystem, searchSymbol byte) Downloader {
	return Downloader{
		httpClient:   httpClient,
		fs:           fs,
		searchSymbol: searchSymbol,
	}
}

type Downloader struct {
	httpClient   http_client.Client
	fs           filesystem.Filesystem
	searchSymbol byte
}

func (d *Downloader) Run(links []string) {
	log.Println("Downloader run")

	d.clean()

	minPosition := &MinPosition{value: math.MaxUint64}
	wg := &sync.WaitGroup{}
	waitSearching := &sync.WaitGroup{}

	for _, link := range links {
		wg.Add(1)
		waitSearching.Add(1)
		go Node(wg, waitSearching, link, d.httpClient, d.fs, d.searchSymbol, minPosition)
	}

	wg.Wait()
	log.Println("Downloader Done")
}

func (d *Downloader) clean() {
	log.Println("Clean download folder")
	d.fs.RemoveAll()
}
