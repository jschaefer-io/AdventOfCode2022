package day07

import (
	"fmt"
)

type directory struct {
	name    string
	parent  *directory
	folders []*directory
	files   map[string]int
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:    name,
		parent:  parent,
		folders: make([]*directory, 0),
		files:   make(map[string]int, 0),
	}
}

func (d *directory) String() string {
	if d.parent == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", d.parent.String(), d.name)
}

func (d *directory) addFolder(name string) *directory {
	folder := newDirectory(name, d)
	folder.parent = d
	d.folders = append(d.folders, folder)
	return folder
}

func (d *directory) addFile(size int, name string) {
	d.files[name] = size
}

func (d *directory) getSize(fs *fileSystem) int {
	if size, ok := fs.sizeCache[d]; ok {
		return size
	}
	count := 0
	for _, f := range d.files {
		count += f
	}
	for _, child := range d.folders {
		count += child.getSize(fs)
	}
	fs.sizeCache[d] = count
	return count
}
