package downloader

import (
	"acronis/infrastructure/filesystem"
	"acronis/infrastructure/http_client"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync"
)

func Node(wg *sync.WaitGroup,
	waitSearching *sync.WaitGroup,
	link string,
	httpClient http_client.Client,
	fs filesystem.Filesystem,
	searchSymbol byte,
	minPosition *MinPosition,
) {
	defer wg.Done()

	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		//@TODO Ignore or Fatal?
		log.Fatal(err)
	}

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		//@TODO Ignore or Fatal?
		log.Fatal(err)
	}

	filename := getBasename(link)
	file, _ := fs.Create(filename)

	buf := make([]byte, 1)

	var position uint64

	shouldBeDownloaded := false

	isSymbolFound := false

	for {
		if shouldBeDownloaded == false && position > minPosition.Get() {
			log.Println(file.Name() + " stop downloading")
			break
		}

		// read a chunk
		n, err := response.Body.Read(buf)
		if err != nil && err != io.EOF {
			//@TODO Ignore or Fatal?
			log.Fatalf(err.Error())
		}
		if n == 0 {
			break
		}

		if isSymbolFound == false && buf[:n][0] == searchSymbol {
			log.Println(file.Name() + " found " + string(searchSymbol) + " on " + strconv.FormatUint(position, 10))
			minPosition.Set(position)
			isSymbolFound = true
			waitSearching.Done()
			waitSearching.Wait()

			shouldBeDownloaded = position == minPosition.Get()
		}

		// write a chunk
		if _, err := file.Write(buf[:n]); err != nil {
			//@TODO Ignore or Fatal?
			log.Fatalf(err.Error())
		}

		position++
	}

	file.Close()

	if isSymbolFound == false {
		waitSearching.Done()
	}

	if shouldBeDownloaded == false {
		fs.Remove(filename)
	}

	log.Println("Done " + file.Name())
}

func getBasename(link string) string {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	return path.Base(u.Path)
}
