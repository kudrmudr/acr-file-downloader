package downloader

import (
	"acronis/downloader"
	"acronis/infrastructure/filesystem"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestDownloadOne(t *testing.T) {

	httpClient := &http.Client{}
	fs := filesystem.New("/tmp/TestDownloadOne")

	searchSymbol := byte('A')
	expectedFile := myFile{path: "/file1", content: []byte("TralAlAlum")}

	files := []myFile{expectedFile}

	ts := HttpFileServerEmulator(files)
	defer ts.Close()

	links := buildDownloaderLinks(ts.URL, files)

	mainDownloader := downloader.New(httpClient, fs, searchSymbol)
	mainDownloader.Run(links)

	filenames, _ := fs.ReadAll()
	assert.NotEmpty(t, filenames)
	assert.Len(t, filenames, 1)
	content, _ := os.ReadFile(filenames[0])

	assert.Equal(t, expectedFile.content, content)
}

func TestNoDownloads(t *testing.T) {

	httpClient := &http.Client{}
	fs := filesystem.New("/tmp/TestNoDownloads")

	searchSymbol := byte('A')
	files := []myFile{
		{path: "/file2", content: []byte("othersymbols")},
	}

	ts := HttpFileServerEmulator(files)
	defer ts.Close()

	links := buildDownloaderLinks(ts.URL, files)

	mainDownloader := downloader.New(httpClient, fs, searchSymbol)
	mainDownloader.Run(links)

	filenames, _ := fs.ReadAll()
	assert.Empty(t, filenames)
}

func TestDownloadSeveral(t *testing.T) {

	httpClient := &http.Client{}
	fs := filesystem.New("/tmp/TestDownloadSeveral")

	searchSymbol := byte('B')
	expectedFileOne := myFile{path: "/file3", content: []byte("aBcde")}
	expectedFileTwo := myFile{path: "/file5", content: []byte("uBtumpurum")}
	expecedNumberOfFiles := 2
	notExpectedFileOne := myFile{path: "/file4", content: []byte("abBcd")}
	notExpectedFileTwo := myFile{path: "/file6", content: []byte("othertext")}

	files := []myFile{
		expectedFileOne,
		notExpectedFileOne,
		expectedFileTwo,
		notExpectedFileTwo,
	}

	ts := HttpFileServerEmulator(files)
	defer ts.Close()

	links := buildDownloaderLinks(ts.URL, files)
	assert.Len(t, links, 4)

	mainDownloader := downloader.New(httpClient, fs, searchSymbol)
	mainDownloader.Run(links)

	filenames, _ := fs.ReadAll()
	assert.NotEmpty(t, filenames)
	assert.Len(t, filenames, expecedNumberOfFiles)

	//@TODO assumption: order of files is matter in the test
	contentOne, _ := os.ReadFile(filenames[0])
	contentTwo, _ := os.ReadFile(filenames[1])

	assert.Equal(t, expectedFileOne.content, contentOne)
	assert.Equal(t, expectedFileTwo.content, contentTwo)
}
