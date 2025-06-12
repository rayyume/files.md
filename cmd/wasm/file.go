package main

import (
	"io/fs"
	"os"
)

type File struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime int64
	isDir   bool
}

func NewFile(name string, modTime int64, isDir bool) *File {
	return &File{
		name:    name,
		modTime: modTime,
		isDir:   isDir,
	}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int64 {
	return 0
}

func (f *File) Mode() os.FileMode {
	return f.mode
}

func (f *File) ModTime() int64 {
	return f.modTime
}

func (f *File) IsDir() bool {
	return f.isDir
}

func (f *File) Sys() any {
	return nil
}
