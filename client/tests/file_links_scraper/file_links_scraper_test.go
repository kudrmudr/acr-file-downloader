package file_links_scraper

import (
	"acronis/file_links_scraper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetLinks(t *testing.T) {
	expectedLinks := []string{
		"file1",
		"file2",
	}
	ts := FileServerEmulator(expectedLinks)
	defer ts.Close()

	mainFileLinksScraper := file_links_scraper.New(&http.Client{}, ts.URL)

	links := mainFileLinksScraper.GetLinks()

	assert.NotEmpty(t, links)

	for i, link := range links {
		assert.Equal(t, ts.URL+"/"+expectedLinks[i], link)
	}
}

func TestGetLinksEmpty(t *testing.T) {
	ts := FileServerEmulator([]string{})
	defer ts.Close()

	mainFileLinksScraper := file_links_scraper.New(&http.Client{}, ts.URL)

	assert.Empty(t, mainFileLinksScraper.GetLinks())
}
