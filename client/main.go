package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {

	println("Main run client")

	links := FileLinksScraper(os.Getenv("FILE_SERVER_URL"))

	for _, link := range links {

		Download(link)
	}

}

func FileLinksScraper(url string) []string {

	var links []string

	res, err := http.Get(url)
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
		links = append(links, BuildUrl(url, href))
	})

	return links
}

func BuildUrl(urlPath string, file string) string {

	u, err := url.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, file)
	return u.String()
}

func GetFilename(link string) string {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	return path.Base(u.Path)
}

func Download(link string) {

	filename := GetFilename(link)

	out, _ := os.Create("/tmp/" + filename)
	defer out.Close()

	resp, _ := http.Get(link)
	defer resp.Body.Close()

	n, _ := io.CopyN(out, resp.Body, 1)

	fmt.Println(n)

}
