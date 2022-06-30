package downloader

import (
	"acronis/url_helper"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func New(dir string) Downloader {
	return Downloader{dir: dir}
}

type Downloader struct {
	dir string
}

func (downloader *Downloader) Download(link string) {

	fileBasename := url_helper.GetBasename(link)
	fileTo := path.Join(downloader.dir, fileBasename)

	out, _ := os.Create(fileTo)
	defer out.Close()

	resp, _ := http.Get(link)
	defer resp.Body.Close()

	n, _ := io.CopyN(out, resp.Body, 1)

	fmt.Println(n)
}
