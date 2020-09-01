// Code generated by vfsgen; DO NOT EDIT.

package resource

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

// Defaults statically implements the virtual filesystem provided to vfsgen.
var Defaults = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 5, 8, 14, 47, 22, 896862493, time.UTC),
		},
		"/app": &vfsgen۰DirInfo{
			name:    "app",
			modTime: time.Date(2020, 4, 23, 0, 4, 32, 255592894, time.UTC),
		},
		"/app/defaults-app.properties": &vfsgen۰FileInfo{
			name:    "defaults-app.properties",
			modTime: time.Date(2020, 3, 16, 22, 37, 55, 58928482, time.UTC),
			content: []byte("\x70\x72\x6f\x66\x69\x6c\x65\x3d\x64\x65\x66\x61\x75\x6c\x74\x0a"),
		},
		"/discovery": &vfsgen۰DirInfo{
			name:    "discovery",
			modTime: time.Date(2020, 2, 27, 18, 56, 54, 491061721, time.UTC),
		},
		"/discovery/consulprovider": &vfsgen۰DirInfo{
			name:    "consulprovider",
			modTime: time.Date(2020, 5, 6, 12, 3, 8, 514064328, time.UTC),
		},
		"/discovery/consulprovider/defaults-discovery.properties": &vfsgen۰FileInfo{
			name:    "defaults-discovery.properties",
			modTime: time.Date(2020, 4, 3, 20, 22, 7, 409821879, time.UTC),
			content: []byte("\x73\x70\x72\x69\x6e\x67\x2e\x63\x6c\x6f\x75\x64\x2e\x63\x6f\x6e\x73\x75\x6c\x2e\x64\x69\x73\x63\x6f\x76\x65\x72\x79\x2e\x69\x6e\x73\x74\x61\x6e\x63\x65\x49\x64\x3d\x75\x75\x69\x64\x0a"),
		},
		"/fs": &vfsgen۰DirInfo{
			name:    "fs",
			modTime: time.Date(2020, 4, 7, 15, 24, 17, 861700325, time.UTC),
		},
		"/fs/defaults-fs.properties": &vfsgen۰CompressedFileInfo{
			name:             "defaults-fs.properties",
			modTime:          time.Date(2020, 3, 18, 2, 1, 55, 520673807, time.UTC),
			uncompressedSize: 131,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcc\xb1\x0d\x42\x31\x0c\x04\xd0\xfe\x4f\xf1\x0b\xea\x78\x02\x86\x71\x2c\xff\xe8\xa4\x60\x47\x76\x42\x83\xd8\x1d\x05\x06\xa0\x3b\x9d\xee\xdd\x95\x25\xdc\xe7\x79\x3f\xe9\xd8\x59\xd3\x57\x88\xe6\x2e\x9e\x1c\xd4\x51\xe9\xf6\xca\x11\xb0\x56\x78\x8c\x0e\xe1\x09\xb7\x62\xfc\xd0\xf7\x26\xe2\x76\xa1\x7d\x81\x4e\xf9\x33\xae\x30\x0e\xfc\xee\x57\x06\x55\xd8\xf1\x09\x00\x00\xff\xff\xb5\x3a\xe4\x1a\x83\x00\x00\x00"),
		},
		"/leader": &vfsgen۰DirInfo{
			name:    "leader",
			modTime: time.Date(2020, 3, 18, 13, 9, 35, 358725984, time.UTC),
		},
		"/leader/defaults-leader.yml": &vfsgen۰CompressedFileInfo{
			name:             "defaults-leader.yml",
			modTime:          time.Date(2020, 3, 18, 13, 9, 35, 358637102, time.UTC),
			uncompressedSize: 177,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x41\x0a\xc2\x30\x10\x85\xe1\x7d\x4e\x31\x8b\x6e\x4d\xf7\xb9\x82\x08\x5e\x61\x4c\x5e\x25\x38\x26\x61\x26\x15\xa4\xf4\xee\x12\xba\x93\xee\xde\xea\xfb\x5f\xac\xc5\x56\xf1\x02\x4e\x50\x0f\x41\xec\xb9\x96\xe0\x88\x50\xf8\x21\x48\x81\x16\x16\x83\x23\x4a\x58\x78\x95\x7e\x63\xeb\xd0\x2b\xbe\x81\x0c\xfa\xc9\x11\xf3\xb4\x59\xd3\x5c\x9e\x9e\x5b\x93\x1c\x79\x10\xbe\xf0\x1b\xfb\x7c\xc0\x8e\xe8\x18\x77\xad\x0d\xda\x33\x6c\x24\x88\x2e\xf4\x1a\xd0\xb4\x9d\xff\xf0\xff\xcd\xdd\xfd\x02\x00\x00\xff\xff\x1f\x71\x63\xb6\xb1\x00\x00\x00"),
		},
		"/security": &vfsgen۰DirInfo{
			name:    "security",
			modTime: time.Date(2020, 4, 3, 20, 22, 7, 410990359, time.UTC),
		},
		"/security/idmdetailsprovider": &vfsgen۰DirInfo{
			name:    "idmdetailsprovider",
			modTime: time.Date(2020, 5, 6, 12, 3, 8, 514396809, time.UTC),
		},
		"/security/idmdetailsprovider/defaults-security-idmdetailsprovider.properties": &vfsgen۰FileInfo{
			name:    "defaults-security-idmdetailsprovider.properties",
			modTime: time.Date(2020, 3, 16, 22, 38, 32, 803682492, time.UTC),
			content: []byte("\x73\x65\x63\x75\x72\x69\x74\x79\x2e\x74\x6f\x6b\x65\x6e\x2e\x64\x65\x74\x61\x69\x6c\x73\x2e\x61\x63\x74\x69\x76\x65\x2d\x63\x61\x63\x68\x65\x2e\x74\x74\x6c\x3a\x20\x35\x73\x0a"),
		},
		"/transit": &vfsgen۰DirInfo{
			name:    "transit",
			modTime: time.Date(2020, 5, 8, 12, 5, 41, 341920880, time.UTC),
		},
		"/transit/defaults-transit.properties": &vfsgen۰CompressedFileInfo{
			name:             "defaults-transit.properties",
			modTime:          time.Date(2020, 3, 19, 18, 59, 25, 223165396, time.UTC),
			uncompressedSize: 268,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x0e\xc2\x30\x0c\x04\xd0\x9d\xaf\xc8\x0f\xb8\x03\x12\x95\x18\xfa\x31\x6e\x38\x50\x14\xe3\x58\x8e\x51\x9b\xbf\x47\x2c\x4c\x48\xc0\x7e\xef\xee\x0c\x4e\x01\x65\x0d\x82\x66\x1f\x16\xa5\xe9\x04\xe5\x55\x70\x49\x4b\xba\xb2\x74\x1c\x3e\xa7\x58\x36\x1e\x9d\xb2\x83\x03\x54\x31\xfa\x17\x50\x31\xc8\xbc\x19\x3c\x0a\xfa\x14\xc3\x90\x96\xc4\xe8\xc7\xd3\x4c\xb7\x7c\x3f\xcf\xbf\x41\xec\xd6\x3c\x5e\x17\xff\x1b\x64\x91\xb6\x91\x09\x17\x0d\xec\x41\x2b\xe7\xfa\xb0\x77\xc7\x33\x00\x00\xff\xff\x9f\x8e\xd6\xc3\x0c\x01\x00\x00"),
		},
		"/webservice": &vfsgen۰DirInfo{
			name:    "webservice",
			modTime: time.Date(2020, 5, 8, 12, 5, 42, 960581030, time.UTC),
		},
		"/webservice/defaults-webservice.properties": &vfsgen۰CompressedFileInfo{
			name:             "defaults-webservice.properties",
			modTime:          time.Date(2020, 3, 20, 19, 37, 36, 274820398, time.UTC),
			uncompressedSize: 119,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x2d\x2a\x4b\x2d\xd2\x2b\x2e\x49\x2c\xc9\x4c\xd6\x2d\x48\x2c\xc9\x50\xb0\x55\xd0\x2f\x2f\x2f\xe7\x82\xca\x94\x14\x25\x26\xa7\xea\xa6\xe6\x25\x26\xe5\xa4\xa6\x28\xd8\x2a\xa4\x25\xe6\x14\xa7\xc2\x24\x0b\xf2\x8b\x4a\x14\x6c\x15\x2c\x0c\x2c\x0c\x60\x42\x19\xf9\xc5\x20\x21\x03\x3d\x30\x84\x89\xa2\xeb\x07\x04\x00\x00\xff\xff\x53\x54\x31\xd7\x77\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app"].(os.FileInfo),
		fs["/discovery"].(os.FileInfo),
		fs["/fs"].(os.FileInfo),
		fs["/leader"].(os.FileInfo),
		fs["/security"].(os.FileInfo),
		fs["/transit"].(os.FileInfo),
		fs["/webservice"].(os.FileInfo),
	}
	fs["/app"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/defaults-app.properties"].(os.FileInfo),
	}
	fs["/discovery"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/discovery/consulprovider"].(os.FileInfo),
	}
	fs["/discovery/consulprovider"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/discovery/consulprovider/defaults-discovery.properties"].(os.FileInfo),
	}
	fs["/fs"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/fs/defaults-fs.properties"].(os.FileInfo),
	}
	fs["/leader"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/leader/defaults-leader.yml"].(os.FileInfo),
	}
	fs["/security"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/security/idmdetailsprovider"].(os.FileInfo),
	}
	fs["/security/idmdetailsprovider"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/security/idmdetailsprovider/defaults-security-idmdetailsprovider.properties"].(os.FileInfo),
	}
	fs["/transit"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/transit/defaults-transit.properties"].(os.FileInfo),
	}
	fs["/webservice"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/webservice/defaults-webservice.properties"].(os.FileInfo),
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
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
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

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
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