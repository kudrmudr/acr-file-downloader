package filesystem

import "os"

type Filesystem interface {
	Create(name string) (*os.File, error)
	Remove(name string) error
	RemoveAll(path string) error
}

func New() Filesystem {
	return filesystem{}
}

type filesystem struct {
}

func (fs filesystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (fs filesystem) Remove(name string) error {
	return os.Remove(name)
}

func (fs filesystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
