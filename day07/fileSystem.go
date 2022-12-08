package day07

import (
	"strings"
)

type fileSystem struct {
	root      *directory
	paths     map[string]*directory
	sizeCache map[*directory]int
}

func newFileSystem() *fileSystem {
	root := newDirectory("", nil)
	fs := &fileSystem{
		root:      root,
		paths:     make(map[string]*directory),
		sizeCache: make(map[*directory]int),
	}
	fs.paths[""] = root
	fs.paths["/"] = root
	return fs
}

func (f *fileSystem) getDirectory(path string) *directory {
	if dir, ok := f.paths[path]; ok {
		return dir
	}
	paths := strings.Split(path, "/")
	pathCutoff := len(paths)
	parent := f.getDirectory(strings.Join(paths[:pathCutoff-1], "/"))
	newDir := parent.addFolder(paths[pathCutoff-1])
	f.paths[newDir.String()] = newDir
	return newDir
}

func (f *fileSystem) iterateWith(handler func(*directory)) []*directory {
	result := make([]*directory, 0)
	for _, d := range f.paths {
		handler(d)
	}
	return result
}
