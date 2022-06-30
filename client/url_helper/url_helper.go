package url_helper

import (
	"log"
	"net/url"
	"path"
)

func BuildUrlBy(urlPath string, filename string) string {

	u, err := url.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, filename)
	return u.String()
}

func GetBasename(link string) string {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	return path.Base(u.Path)
}
