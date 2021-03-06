// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// TemplateAssets contains project assets.
var TemplateAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 4, 6, 0, 22, 43, 758505986, time.UTC),
		},
		"/login.html": &vfsgen۰CompressedFileInfo{
			name:             "login.html",
			modTime:          time.Date(2020, 4, 6, 0, 22, 43, 758393366, time.UTC),
			uncompressedSize: 1295,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcd\x6e\xdb\x3c\x10\x3c\x47\x4f\xb1\x1f\xef\xfe\x88\xdc\x29\x15\x45\xfa\x67\xc0\x6d\x82\xd6\x39\xe4\x64\x50\xe2\x46\x22\x4a\x91\x04\xb9\x74\xe3\x06\x79\xf7\x82\xfa\x41\x6c\x43\x28\x8c\xa2\xf5\xc5\xda\xe1\xec\xcc\x2c\x28\xad\xf8\xef\xdd\xed\xcd\xf6\xe1\xee\x3d\x7c\xda\x7e\xde\x54\xa2\xa3\xde\x80\x91\xb6\x2d\x19\x5a\x56\x15\xa2\x43\xa9\xaa\x02\x00\x40\x90\x26\x83\xd5\xc6\xb5\xda\x0a\x3e\x16\x85\xe0\x23\x41\xd4\x4e\x1d\xaa\x42\x3c\xba\xd0\x43\x8f\xd4\x39\x55\x32\xef\x22\x31\x68\x8c\x8c\xb1\x64\x26\xf7\xad\xbc\x6c\x91\x81\x6c\x48\x3b\x5b\x32\x2e\x13\x75\x2e\xe8\x9f\xc8\x26\x8f\xee\x3a\x1b\xc0\xda\xc2\xd6\xc1\x83\x4b\x01\xde\x36\x8d\x4b\x96\x04\xef\xae\x27\x8e\xb6\x3e\x11\xec\xa5\x49\x58\xb2\xe7\x67\xf8\xff\xc6\x68\xb4\xb4\x56\xf0\xf2\xc2\xc0\xca\x1e\x4b\xd6\x0c\xd0\x4e\x2b\x06\x74\xf0\x58\xb2\x4e\x2b\x35\x4c\xb4\x2c\xf1\x15\x95\x0e\xd8\xd0\x7d\xd0\x47\x2a\x61\x42\x77\x29\xe8\x8b\x85\xa2\x77\x36\xe2\xf6\xe0\xf1\x44\x69\x84\x77\x59\xe4\x42\xa9\x6f\x8d\x3b\xd1\x88\xb9\xbe\xb4\x97\x24\x9d\xf4\xe6\x7a\xb9\x57\xe9\xfd\x7c\x47\xf9\xf6\x56\xb5\x7b\x9a\x8e\x16\x8f\xf7\x18\x48\x37\xd2\x1c\x71\x16\x79\x6d\x70\xc9\x9f\x91\x06\xa2\x91\x35\x9a\x13\xea\x80\xb0\xea\x3e\x62\xc8\x69\x05\x1f\x80\x85\xd6\x71\xcc\x71\x22\xec\xa5\x36\xf3\x44\x84\x4f\xaf\x6f\xda\xa0\x39\x50\xcf\x33\x72\xa5\xf7\x7f\x3d\xf6\x9d\x8c\xf1\x87\x0b\xea\xb2\xd8\x7e\x62\xcf\xc9\x5f\xeb\x7f\x9a\xbe\x4e\x44\xce\x4e\x9e\x31\xd5\xbd\x3e\xff\x32\x47\x06\xab\x36\xb7\x1f\xd7\x5f\x04\x1f\xcb\x3f\x4f\x70\x95\x7f\x22\x7a\x69\x8f\x8e\x5b\x47\xab\x79\xe0\x95\xd1\xf6\xfb\x4c\xbc\x12\x12\xba\x80\x8f\x25\xe3\x67\x34\x56\x7d\x18\x00\x38\xe4\x65\x30\xa3\x6f\x04\x97\xb3\x07\xcf\x26\xbf\x0d\x7a\x54\x4e\x8f\x22\xdb\xf4\xf9\x7f\xda\x5b\x3c\x2f\xbe\xaa\xf8\x15\x00\x00\xff\xff\x6a\x66\xa3\x63\x0f\x05\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/login.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
