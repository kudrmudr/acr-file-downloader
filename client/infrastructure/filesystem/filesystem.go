package filesystem

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Filesystem interface {
	Create(name string) (*os.File, error)
	Remove(name string) error
	ReadAll() ([]string, error)
	Clean() error
}

func New(dir string) Filesystem {
	fs := filesystem{
		dir: dir,
	}
	fs.Init()
	return fs
}

type filesystem struct {
	dir string
}

func (fs filesystem) Create(name string) (*os.File, error) {
	log.Println("Create", name)
	return os.Create(path.Join(fs.dir, name))
}

func (fs filesystem) Remove(name string) error {
	log.Println("Remove", name)
	return os.Remove(path.Join(fs.dir, name))
}

func (fs filesystem) ReadAll() ([]string, error) {

	result := []string{}

	err := filepath.Walk(fs.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() == false {
			result = append(result, path)
		}
		return nil
	})

	return result, err
}

func (fs filesystem) Clean() error {
	err := os.RemoveAll(fs.dir)
	fs.Init()
	return err
}

func (fs filesystem) Init() {
	if err := os.MkdirAll(fs.dir, 0755); err != nil {
		log.Fatalf(err.Error())
	}
}
