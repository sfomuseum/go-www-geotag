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

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3d\x4f\xc3\x30\x10\x86\xf7\xfc\x8a\xc3\x33\xa1\x62\x63\xb0\xb3\x21\x84\xc4\xc6\xc6\x82\x0e\xfb\x9a\x58\x72\xee\xa2\xe4\xda\xaa\x58\xf9\xef\xc8\x4d\x22\xbe\x3a\xd9\xf7\xf8\x79\xa5\xd7\x97\x33\x04\xda\x47\x26\x30\x91\xfd\x7b\x47\x18\x0c\xcc\x73\x65\x6f\x82\x78\x3d\x0f\x04\x9d\xf6\xa9\xa9\x6c\x39\x20\x21\xb7\xce\x10\x9b\xa6\x02\xb0\x45\x2e\x17\x00\xdb\x93\x22\xf8\x0e\xc7\x89\xd4\x99\x83\xee\xeb\x07\xb3\x3e\x69\xd4\x44\xcd\x13\x89\x62\x6b\x77\xcb\xf4\x23\xc4\xd8\x93\x33\xc7\x48\xa7\x41\x46\x35\xe0\x85\x95\x58\x9d\x39\xc5\xa0\x9d\x0b\x74\x8c\x9e\xea\xcb\x70\x0b\x91\xa3\x46\x4c\xf5\xe4\x31\x91\xbb\x5f\x6a\xec\xb6\x1e\xf6\x43\xc2\x19\x02\x2a\xd6\x9b\x98\x50\xa3\x1e\x02\x39\x93\x33\xdc\x3d\x2f\xf4\x65\x85\x30\xcf\xe6\x8f\x2e\xdc\x5e\xf1\x37\xfa\x3f\xf0\x29\xd2\xff\x72\xdf\x44\xfa\x6f\x6d\x22\x1c\x7d\x57\x13\x87\x41\x62\xf9\x54\x31\x5f\x2f\xf0\x71\x65\x45\x6e\xaa\x9c\x81\x38\x94\xc5\x7f\x05\x00\x00\xff\xff\xb8\xf4\xba\x98\x91\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 401, mode: os.FileMode(420), modTime: time.Unix(1586215562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4b\x6e\xeb\x3a\x0c\x1d\x27\xab\x20\x38\x77\xfc\x5e\x87\xad\xed\x59\x57\xd0\x05\x3c\xd0\x16\xdd\xa8\x4f\x1f\x5f\x8a\x0e\x1a\x04\xdd\xfb\x85\xe4\xb8\x75\x3f\xa3\x22\x93\x84\x3e\xc4\x39\x3a\xfc\x48\x97\x0b\x18\x1e\x6d\x60\x40\x1b\x0c\xbf\x22\xbc\xbd\xed\x2f\x17\x50\xf6\x93\x23\x2d\xf0\xf0\xdf\x91\xc9\x20\x1c\x72\xae\x71\x36\xfc\x0f\xc2\xae\xc5\xa4\x67\xc7\xe9\xc8\xac\x08\x7a\x9e\xb8\x45\xe5\x57\xad\x87\x94\x10\x8e\xc2\x63\x8b\x39\xae\x8b\xee\xa1\xa0\x75\xf7\x0b\xfe\xe4\x68\xe0\x63\x74\x86\xe5\x43\x25\x0d\x62\x27\xdd\xd2\x5e\xe8\x44\x0b\x8a\x90\x64\x68\x71\x83\xd4\xcf\x1c\x95\x9e\x0f\x9e\xa6\x74\x78\x49\xd8\x35\xf5\x92\xf8\xb5\xd2\x40\x9e\x85\x6e\xa3\xb5\x2d\xf0\x96\xe6\x6c\xb0\x7a\x1b\xc1\xd2\xb7\xdb\xc9\x6d\x0b\xfe\x49\xd5\xd8\x13\x0c\x8e\x52\x6a\x71\x88\x41\xc9\x06\x16\xec\xf6\x7b\x00\x80\x26\xd0\x09\xca\xe6\xb4\xe8\x49\x9e\x6d\xa8\x34\x4e\xf7\xff\x0a\xfb\x87\xeb\x77\x1f\x55\xa3\x5f\x20\xec\x16\x56\x1d\xe8\xb4\x2a\x6c\xe5\x49\x4c\x56\xde\x7d\x05\xab\xbc\xf2\xcb\xa9\xbb\x4c\xba\x5c\xc0\x8e\x70\x78\x0c\xd4\x3b\x7e\x62\x92\xe1\x98\x6f\x43\xc9\x35\xc9\x93\x73\xab\xa9\x31\x06\xad\x4a\x7c\x6f\x95\x9c\x1d\x1e\xb0\x7b\x52\x61\x56\x20\x63\x84\x53\xe2\x04\x14\x0c\x90\x95\x29\x8a\x26\x20\x61\x08\x51\x61\x98\x45\x38\xa8\x3b\x43\x9a\xa7\x9c\x62\xd3\xd4\x45\xbb\xdb\xed\xca\x49\xd7\xf3\xc6\x28\x7e\x35\x9b\xe3\xca\x06\x67\x03\x63\xb7\xdf\xed\x1a\x1b\xa6\x59\x3f\x65\x73\x0f\x25\x3a\xf0\x52\x25\x5f\xdd\xad\x97\x2d\x95\x2a\x10\x02\x79\x6e\x71\x33\x93\xea\xcf\xcc\x72\x46\xb0\xe6\x47\x78\x03\xb5\xf8\x18\x94\x05\x68\x01\x8b\x14\x68\x84\x45\x1a\xc6\x28\x08\x27\x72\x33\xb7\x88\x40\x62\xa9\x72\xd4\xe7\xab\xbf\x74\x10\xeb\xe2\xb8\x9f\x55\x63\x58\x2d\xf7\x1a\xc0\x9f\xab\xbb\xfc\x93\x7c\xf5\xcf\xbb\xdd\xb9\xf7\x56\xb1\x5b\xa8\x4d\xbd\xb0\xba\x6b\x4b\xea\x5c\xe9\xfa\x91\x67\xf9\xd5\xbc\x70\x9a\x9d\x96\x3d\x33\xf6\xd4\xbd\x4f\x95\x83\x29\x83\x5c\xe1\xeb\x2a\x64\xba\xa7\x09\x3f\xed\x44\x1f\xcd\xf9\x43\xe0\xdb\xca\x8c\x31\x6a\x59\x99\xcf\x2e\x46\x26\x9d\x85\x37\xc4\xe5\x7f\xf1\xfd\x1e\xee\xaf\xf1\xb7\xc7\x37\xcb\x5e\x1f\xdf\x0f\xbf\x7f\x03\x00\x00\xff\xff\x72\xbf\x9d\x93\xb8\x05\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/index.html", size: 1464, mode: os.FileMode(420), modTime: time.Unix(1586282378, 0)}
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

