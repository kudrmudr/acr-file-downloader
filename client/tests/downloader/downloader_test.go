package downloader

import (
	"acronis/downloader"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type myFile struct {
	path    string
	content string
}

func TestDownloadOne(t *testing.T) {

	expectedFiles := []myFile{
		{path: "/file1", content: "tralalaA"},
	}

	ts := EmulateFileServer(expectedFiles)
	defer ts.Close()

	links := buildDownloaderLinks(ts.URL, expectedFiles)

	mainDownloader := downloader.New(http.Client{}, "/usr/share/downloads", byte('A'))
	mainDownloader.Run(links)
}

func buildDownloaderLinks(url string, myFiles []myFile) []string {
	links := make([]string, len(myFiles))
	for i, myFile := range myFiles {
		links[i] = url + myFile.path
	}

	return links
}

func EmulateFileServer(files []myFile) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, file := range files {
			if r.URL.Path == file.path {
				fmt.Fprintln(w, file.content)
				break
			}
		}
	}))

	return ts
}
