package filesystem

import (
	"os"
	"path"
)

type Filesystem interface {
	Create(name string) (*os.File, error)
	Remove(name string) error
	RemoveAll() error
}

func New(dir string) Filesystem {
	return filesystem{
		dir: dir,
	}
}

type filesystem struct {
	dir string
}

func (fs filesystem) Create(name string) (*os.File, error) {
	return os.Create(path.Join(fs.dir, name))
}

func (fs filesystem) Remove(name string) error {
	return os.Remove(path.Join(fs.dir, name))
}

func (fs filesystem) RemoveAll() error {
	return os.RemoveAll(fs.dir)
}
