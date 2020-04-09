// Code generated by go-bindata.
// sources:
// templates/html/inc_foot.html
// templates/html/inc_head.html
// templates/html/index.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHtmlInc_footHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\x4b\x8e\x4f\xcb\xcf\x2f\x51\x52\xa8\xad\xe5\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x56\x48\xcd\x4b\x01\xc9\x01\x02\x00\x00\xff\xff\xe4\xf0\x0a\x4b\x34\x00\x00\x00")

func templatesHtmlInc_footHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_footHtml,
		"templates/html/inc_foot.html",
	)
}

func templatesHtmlInc_footHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_footHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_foot.html", size: 52, mode: os.FileMode(420), modTime: time.Unix(1586200144, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x3f\x6f\xe3\x30\x0c\xc5\x77\x7f\x0a\x1e\xe7\x73\x8c\xdb\x6e\xb0\xbc\x05\x45\x81\x02\xed\xdc\xa5\x60\x24\x26\x56\x20\x8b\x86\xc5\x24\x4d\x0d\x7f\xf7\x42\x75\xdc\xa4\x7f\x26\x89\xef\xfd\x9e\x28\x0a\x1a\x47\x70\xbc\xf5\x91\x01\x7d\xb4\x2f\x2d\x93\x43\x98\xa6\xa2\xfe\xe3\xc4\xea\xb9\x67\x68\xb5\x0b\x4d\x51\xe7\x05\x02\xc5\x9d\x41\x8e\xd8\x14\x00\x75\x86\xf3\x06\xa0\xee\x58\x09\x6c\x4b\x43\x62\x35\x78\xd0\x6d\xf9\x1f\x2f\x96\x7a\x0d\xdc\xdc\xb1\x28\xed\xea\x6a\xae\x6e\x42\x91\x3a\x36\x78\xf4\x7c\xea\x65\x50\x04\x2b\x51\x39\xaa\xc1\x93\x77\xda\x1a\xc7\x47\x6f\xb9\xfc\x28\xfe\x82\x8f\x5e\x3d\x85\x32\x59\x0a\x6c\xfe\x2d\x2d\x92\x1d\x7c\xaf\x90\xaf\x6b\x50\xf9\x55\xab\x3d\x1d\x69\x56\x11\xd2\x60\x0d\xde\x28\xd5\xa9\x95\x24\x71\xeb\x87\xa4\xab\xc8\xba\xda\x27\x6c\xea\x6a\x36\x9b\x7c\x60\x1e\xae\x5a\xa6\xab\x37\xe2\xce\xe0\x48\xa9\x5c\xda\x07\x52\xaf\x07\xc7\x06\xc7\x11\x56\xf7\xb3\xfa\x70\x11\x61\x9a\xf0\x1b\x2e\x71\xf7\x0b\xbf\xa8\x3f\x03\x6f\x22\xdd\x17\xf6\x59\xa4\xbb\x62\x7d\x20\xcb\xad\x04\xc7\x43\xc9\xd1\xf5\xe2\xf3\x7b\x65\xfc\xe9\xea\xac\x2f\xc6\x35\x26\xdc\x6d\xd8\x7d\x26\xd2\x1c\x79\x5c\x67\x75\xa1\x53\xc6\x9b\x62\x1c\x81\xa3\xcb\xff\xe0\x3d\x00\x00\xff\xff\x2e\xf7\xe5\x06\x20\x02\x00\x00")

func templatesHtmlInc_headHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_headHtml,
		"templates/html/inc_head.html",
	)
}

func templatesHtmlInc_headHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_headHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 544, mode: os.FileMode(420), modTime: time.Unix(1586471433, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x4f\x6f\x9b\x4c\x10\xc6\xcf\xf0\x29\x46\x7b\xc7\x44\x79\xaf\xc0\x2d\xef\xa9\x52\xab\x56\x3d\x57\x03\x0c\xf1\xa6\xfb\x87\xee\x0e\x56\x5c\x94\xef\x5e\x2d\x8b\x6d\x48\x48\xa5\xba\xee\xc5\x5a\xcf\x30\xbf\x9d\x79\xe6\x41\x8c\x23\xb4\xd4\x49\x43\x20\xa4\x69\xe9\x59\xc0\xcb\x4b\x3a\x8e\xc0\xa4\x7b\x85\x3c\x85\x9b\x6f\x7b\xc2\x56\xc0\x2e\xe4\x0a\x25\xcd\x77\x70\xa4\x4a\xe1\xf9\xa8\xc8\xef\x89\x58\x00\x1f\x7b\x2a\x05\xd3\x33\xe7\x8d\xf7\x02\xf6\x8e\xba\x52\x84\x73\x3e\x71\x77\x53\x34\xaf\xd2\xc2\x37\x4e\xf6\xbc\x2c\x78\xc2\x03\xc6\xa8\x00\xef\x9a\x52\x2c\x22\xf9\x23\x59\xc6\xc7\x9d\xc6\xde\xef\x9e\xbc\xa8\x8a\x3c\x26\xae\x26\x35\xa8\xc9\xe1\x4d\x59\xd2\x48\xbe\x0d\x70\x1a\x73\x03\x37\x8e\x20\x3b\xd8\x3d\x18\xac\x15\x7d\x7c\xd0\x35\xb5\xd3\x32\xae\xbb\xc5\x52\x00\xdc\xa6\xe3\x99\xb5\x25\xc1\x1f\x3b\x65\x66\x9d\xac\x32\x8e\x40\xa6\x9d\x0d\x79\x19\xff\x93\xc2\x86\xf6\x56\xb5\xe4\xfe\x42\x83\xfe\x42\xb9\x8d\x10\x4b\xe0\x4d\xd4\x58\x02\xdf\x4a\x52\xb4\xf2\x00\x8d\x42\xef\x4b\xd1\x58\xc3\x28\x0d\x39\x51\xa5\x29\x00\xc0\xb6\x5d\x42\x26\xe6\x57\xc5\xe8\xda\x50\x97\xbc\x0e\x66\xe1\xad\x8f\xcc\x64\x2a\xea\xac\xd3\xa7\x07\xc2\x39\x93\x46\x49\x43\xa2\x4a\x93\xa4\x90\xa6\x1f\x78\x95\x0d\x5d\x39\xab\x40\xbb\xcc\xeb\xec\xfe\x34\xab\x27\x74\xcd\x5e\x80\x41\x4d\xa5\x88\x2b\xcf\x06\xa7\x04\xc8\x76\xfd\x7f\x21\x40\x29\x1e\x0c\x93\x03\x84\xaf\x9f\x3f\x00\x5b\x70\xc4\x4e\xd2\x81\x04\x1c\x50\x0d\x54\x0a\x01\xe8\x24\x66\x0a\xeb\x20\xf1\xff\xc4\xe1\x0e\x2f\x7f\x52\x29\xfe\xbb\x9b\xb4\x4b\x92\xa2\x1e\x98\xad\x39\x75\x59\xb3\x01\x7d\xcc\xee\xc3\x8f\xd7\xd9\xdd\xb9\xc3\xa1\xd6\x92\x57\xfd\x74\x13\xaf\x9a\xb0\x45\x1e\x31\xd5\x2c\x4b\x1e\xa6\xad\xd2\xa4\xc8\x5b\x79\xa8\xd2\x64\x56\x72\x51\x2d\x35\x3e\x92\x58\x69\x5b\xdb\xf6\x18\xec\x11\x4b\xde\x48\xdf\x59\xcb\xe4\x56\x2d\x68\x62\x3c\x57\xc4\x8b\xcf\xc7\xd3\xd2\x67\x6f\x9c\x43\x57\xed\xf9\x77\x2f\xdb\xbf\x33\xc2\x62\xd9\xd9\x8f\x81\xdc\x31\x0e\xbf\x11\xde\xb4\xc5\x14\x9c\x50\xc1\x1d\x11\x0d\x9d\x75\xef\xf8\xe3\x4b\xbc\xfb\x0a\x57\x54\xb1\xf4\x3d\x13\x9c\x55\x7f\xdd\xbc\x23\x3f\x28\xf6\x97\x9d\xaf\x57\xb6\xe5\x1e\x8d\xfd\x35\xa6\x79\xdd\x45\x47\xc8\x83\xa3\x45\xe1\xb6\x87\xe6\xf3\x9b\x8f\x7f\xc0\xce\x1f\xff\x4b\xbf\xbf\x02\x00\x00\xff\xff\x1b\x35\xe7\x7f\x38\x08\x00\x00")

func templatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIndexHtml,
		"templates/html/index.html",
	)
}

func templatesHtmlIndexHtml() (*asset, error) {
	bytes, err := templatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/index.html", size: 2104, mode: os.FileMode(420), modTime: time.Unix(1586471707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/html/inc_foot.html": templatesHtmlInc_footHtml,
	"templates/html/inc_head.html": templatesHtmlInc_headHtml,
	"templates/html/index.html": templatesHtmlIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"inc_foot.html": &bintree{templatesHtmlInc_footHtml, map[string]*bintree{}},
			"inc_head.html": &bintree{templatesHtmlInc_headHtml, map[string]*bintree{}},
			"index.html": &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

