// Package stringfs implements a read-only file system whose contents
// are held in a map from path name to file contents.
//
// This is the first draft of Section 19.7; it does not yet satisfy
// fstest.TestFS.
package stringfs

import (
	"io/fs"
	"path"
	"strings"
	"time"
)

// An FS maps slash-separated path names to file contents.
type FS map[string]string

func (fsys FS) Open(name string) (fs.File, error) {
	text, ok := fsys[name]
	if !ok {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}
	return &openFile{fileInfo{path.Base(name), int64(len(text))}, strings.NewReader(text)}, nil
}

// An openFile is a file being read. The embedded *strings.Reader
// supplies Read, and also ReadAt and Seek.
type openFile struct {
	info fileInfo
	*strings.Reader
}

func (f *openFile) Stat() (fs.FileInfo, error) { return f.info, nil }
func (f *openFile) Close() error               { return nil }

type fileInfo struct {
	name string
	size int64
}

func (i fileInfo) Name() string       { return i.name }
func (i fileInfo) Size() int64        { return i.size }
func (i fileInfo) Mode() fs.FileMode  { return 0444 }
func (i fileInfo) ModTime() time.Time { return time.Time{} }
func (i fileInfo) IsDir() bool        { return false }
func (i fileInfo) Sys() any           { return nil }
