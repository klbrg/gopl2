// Package stringfs implements a read-only file system whose contents
// are held in a map from path name to file contents.
package stringfs

import (
	"errors"
	"io"
	"io/fs"
	"path"
	"slices"
	"strings"
	"time"
)

// An FS maps slash-separated path names to file contents.
type FS map[string]string

func (fsys FS) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrInvalid}
	}
	if text, ok := fsys[name]; ok {
		info := fileInfo{path.Base(name), int64(len(text)), false}
		return &openFile{info, strings.NewReader(text)}, nil
	}
	entries, err := fsys.ReadDir(name) // not a regular file; try a directory
	if err != nil {
		return nil, err
	}
	return &openDir{fileInfo{path.Base(name), 0, true}, entries, 0}, nil
}

// ReadDir returns the entries of directory dir, sorted by name.
func (fsys FS) ReadDir(dir string) ([]fs.DirEntry, error) {
	if !fs.ValidPath(dir) {
		return nil, &fs.PathError{Op: "readdir", Path: dir, Err: fs.ErrInvalid}
	}
	prefix := ""
	if dir != "." {
		prefix = dir + "/"
	}
	seen := make(map[string]fs.DirEntry)
	for name, text := range fsys {
		rest, ok := strings.CutPrefix(name, prefix)
		if !ok || rest == "" {
			continue
		}
		if base, _, isDir := strings.Cut(rest, "/"); isDir {
			seen[base] = fs.FileInfoToDirEntry(fileInfo{base, 0, true})
		} else {
			seen[base] = fs.FileInfoToDirEntry(fileInfo{base, int64(len(text)), false})
		}
	}
	if len(seen) == 0 {
		return nil, &fs.PathError{Op: "readdir", Path: dir, Err: fs.ErrNotExist}
	}
	entries := make([]fs.DirEntry, 0, len(seen))
	for _, e := range seen {
		entries = append(entries, e)
	}
	slices.SortFunc(entries, func(a, b fs.DirEntry) int {
		return strings.Compare(a.Name(), b.Name())
	})
	return entries, nil
}

// An openFile is a regular file being read. The embedded *strings.Reader
// supplies Read, and also ReadAt and Seek.
type openFile struct {
	info fileInfo
	*strings.Reader
}

func (f *openFile) Stat() (fs.FileInfo, error) { return f.info, nil }
func (f *openFile) Close() error               { return nil }

// An openDir is a directory being read.
type openDir struct {
	info    fileInfo
	entries []fs.DirEntry
	offset  int
}

func (d *openDir) Stat() (fs.FileInfo, error) { return d.info, nil }
func (d *openDir) Close() error               { return nil }

func (d *openDir) Read([]byte) (int, error) {
	return 0, &fs.PathError{Op: "read", Path: d.info.name, Err: errors.New("is a directory")}
}

func (d *openDir) ReadDir(n int) ([]fs.DirEntry, error) {
	rest := len(d.entries) - d.offset
	if n <= 0 {
		list := d.entries[d.offset:]
		d.offset = len(d.entries)
		return list, nil
	}
	if rest == 0 {
		return nil, io.EOF
	}
	list := d.entries[d.offset : d.offset+min(n, rest)]
	d.offset += len(list)
	return list, nil
}

type fileInfo struct {
	name  string
	size  int64
	isDir bool
}

func (i fileInfo) Name() string { return i.name }
func (i fileInfo) Size() int64  { return i.size }
func (i fileInfo) Mode() fs.FileMode {
	if i.isDir {
		return fs.ModeDir | 0555
	}
	return 0444
}
func (i fileInfo) ModTime() time.Time { return time.Time{} }
func (i fileInfo) IsDir() bool        { return i.isDir }
func (i fileInfo) Sys() any           { return nil }
