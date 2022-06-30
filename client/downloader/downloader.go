package downloader

import (
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
)

func New(dir string, stopSymbol byte) Downloader {
	return Downloader{
		dir:        dir,
		stopSymbol: stopSymbol,
	}
}

type Downloader struct {
	dir         string
	stopSymbol  byte
	minPosition *MinPosition
	wg          *sync.WaitGroup
}

func (d *Downloader) Run(links []string) {
	log.Println("Downloader run")

	d.minPosition = &MinPosition{value: math.MaxUint64}
	d.wg = &sync.WaitGroup{}

	for _, link := range links {
		d.wg.Add(1)
		go d.download(link)
	}

	d.wg.Wait()
	log.Println("Downloader Done")
}

func (d *Downloader) download(link string) {
	log.Println("Start " + link)
	defer d.wg.Done()

	fileTo := buildFilePathBy(link, d.dir)

	out, _ := os.Create(fileTo)
	defer out.Close()

	response, _ := http.Get(link)
	defer response.Body.Close()

	buf := make([]byte, 1)

	var position uint64

	var myMinPosition uint64 = math.MaxUint64
	isSymbolFound := false

	for {
		if myMinPosition != d.minPosition.Get() && position > d.minPosition.Get() {
			log.Println(link + " stop downloading")
			break
		}

		// read a chunk
		n, err := response.Body.Read(buf)
		if err != nil && err != io.EOF {
			//@TODO  error
			log.Fatalf(err.Error())
		}
		if n == 0 {
			break
		}

		if isSymbolFound == false && buf[:n][0] == d.stopSymbol {
			log.Println(link + " found " + string(d.stopSymbol) + " on " + strconv.FormatUint(position, 10))
			d.minPosition.Set(position)
			myMinPosition = position
			isSymbolFound = true
		}

		// write a chunk
		if _, err := out.Write(buf[:n]); err != nil {
			log.Fatalf(err.Error())
		}

		position++
	}

	out.Close()

	if isSymbolFound == false || myMinPosition > d.minPosition.Get() {
		os.Remove(fileTo)
	}

	log.Println("Done " + link)
}

func buildFilePathBy(link string, dir string) string {
	return path.Join(dir, getBasename(link))
}

func getBasename(link string) string {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	return path.Base(u.Path)
}
