package file_links_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"path"
)

func New(fileServerUrl string) FileLinksScraper {

	return FileLinksScraper{
		fileServerUrl: fileServerUrl,
	}
}

type FileLinksScraper struct {
	fileServerUrl string
}

func (scraper *FileLinksScraper) GetLinks() []string {
	var links []string

	res, err := http.Get(scraper.fileServerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("body pre a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		links = append(links, buildUrlBy(scraper.fileServerUrl, href))
	})

	return links
}

func buildUrlBy(urlPath string, filename string) string {

	u, err := url.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, filename)
	return u.String()
}
