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

func New(httpClient http.Client, dir string, stopSymbol byte) Downloader {
	return Downloader{
		httpClient: httpClient,
		dir:        dir,
		stopSymbol: stopSymbol,
	}
}

type Downloader struct {
	httpClient    http.Client
	dir           string
	stopSymbol    byte
	minPosition   *MinPosition
	wg            *sync.WaitGroup
	waitSearching *sync.WaitGroup
}

func (d *Downloader) Run(links []string) {
	log.Println("Downloader run")

	d.clean()

	d.minPosition = &MinPosition{value: math.MaxUint64}
	d.wg = &sync.WaitGroup{}
	d.waitSearching = &sync.WaitGroup{}

	for _, link := range links {
		d.wg.Add(1)
		go d.download(link)
	}

	d.wg.Wait()
	log.Println("Downloader Done")
}

func (d *Downloader) clean() {
	log.Println("Clean download folder")
	os.RemoveAll(d.dir)
}

func (d *Downloader) download(link string) {
	log.Println("Start " + link)
	defer d.wg.Done()

	fileTo := buildFilePathBy(link, d.dir)

	out, _ := os.Create(fileTo)
	defer out.Close()

	response, _ := d.httpClient.Get(link)
	defer response.Body.Close()

	buf := make([]byte, 1)

	var position uint64

	isThisIt := false

	isSymbolFound := false

	d.waitSearching.Add(1)

	for {
		if isThisIt == false && position > d.minPosition.Get() {
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
			isSymbolFound = true
			d.waitSearching.Done()
			d.waitSearching.Wait()

			isThisIt = position == d.minPosition.Get()
		}

		// write a chunk
		if _, err := out.Write(buf[:n]); err != nil {
			log.Fatalf(err.Error())
		}

		position++
	}

	out.Close()

	if isSymbolFound == false {
		d.waitSearching.Done()
	}

	if isThisIt == false {
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
