package file_links_scraper

import (
	"acronis/file_links_scraper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetLinks(t *testing.T) {

	httpClient := http.Client{}
	mainFileLinksScraper := file_links_scraper.New(httpClient, os.Getenv("FILE_SERVER_URL"))

	assert.NotEmpty(t, mainFileLinksScraper.GetLinks())
}
