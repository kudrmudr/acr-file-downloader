package file_links_scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func FileServerEmulator(links []string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body><pre>")
		for _, link := range links {
			fmt.Fprintln(w, "<a href='"+link+"'>some name</a>")
		}
		fmt.Fprintln(w, "</pre></body></html>")
	}))

	return ts
}
