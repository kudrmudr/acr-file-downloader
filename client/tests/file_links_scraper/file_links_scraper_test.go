package file_links_scraper

import (
	"acronis/file_links_scraper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLinks(t *testing.T) {
	expectedLinks := []string{
		"file1",
		"file2",
	}
	ts := EmulateFileServer(expectedLinks)
	defer ts.Close()

	mainFileLinksScraper := file_links_scraper.New(&http.Client{}, ts.URL)

	links := mainFileLinksScraper.GetLinks()

	assert.NotEmpty(t, links)

	for i, link := range links {
		assert.Equal(t, ts.URL+"/"+expectedLinks[i], link)
	}
}

func TestGetLinksEmpty(t *testing.T) {
	ts := EmulateFileServer([]string{})
	defer ts.Close()

	mainFileLinksScraper := file_links_scraper.New(&http.Client{}, ts.URL)

	assert.Empty(t, mainFileLinksScraper.GetLinks())
}

func EmulateFileServer(links []string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body><pre>")
		for _, link := range links {
			fmt.Fprintln(w, "<a href='"+link+"'>some name</a>")
		}
		fmt.Fprintln(w, "</pre></body></html>")
	}))

	return ts
}
