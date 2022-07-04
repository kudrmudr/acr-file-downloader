package downloader

import (
	"log"
	"net/http"
	"net/http/httptest"
)

type myFile struct {
	path    string
	content []byte
}

func HttpFileServerEmulator(files []myFile) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, file := range files {
			if r.URL.Path == file.path {
				_, err := w.Write(file.content)
				if err != nil {
					log.Fatal(err.Error())
				}
				break
			}
		}
	}))

	return ts
}

func buildDownloaderLinks(url string, myFiles []myFile) []string {
	links := make([]string, len(myFiles))
	for i, myFile := range myFiles {
		links[i] = url + myFile.path
	}

	return links
}
