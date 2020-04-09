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

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xbd\x6e\xe3\x30\x0c\x80\x77\x3f\x05\x8f\xf3\x39\xc6\x6d\x37\x48\xde\x82\xc3\x01\x05\x3a\x74\xeb\x52\x30\x12\x13\x2b\xb0\x45\xc3\x62\x92\xa6\x86\xdf\xbd\x50\x1c\x37\xe9\xcf\x24\xf1\xe3\x47\x51\x94\xc6\x11\x3c\x6f\x43\x64\xc0\x10\xdd\x4b\xc3\xe4\x11\xa6\xa9\x30\xbf\xbc\x38\x3d\xf7\x0c\x8d\x76\x6d\x5d\x98\xbc\x40\x4b\x71\x67\x91\x23\xd6\x05\x80\xc9\x72\xde\x00\x98\x8e\x95\xc0\x35\x34\x24\x56\x8b\x07\xdd\x96\x7f\xf1\x9a\xd2\xa0\x2d\xd7\xff\x58\x94\x76\xa6\x9a\xa3\xbb\xa2\x48\x1d\x5b\x3c\x06\x3e\xf5\x32\x28\x82\x93\xa8\x1c\xd5\xe2\x29\x78\x6d\xac\xe7\x63\x70\x5c\x5e\x82\xdf\x10\x62\xd0\x40\x6d\x99\x1c\xb5\x6c\xff\x2c\x2d\x92\x1b\x42\xaf\x90\xaf\x6b\x51\xf9\x55\xab\x3d\x1d\x69\xa6\x08\x69\x70\x16\xef\x48\x75\x6a\x24\x49\xdc\x86\x21\xe9\x2a\xb2\xae\xf6\x09\x6b\x53\xcd\xc9\x3a\x1f\x98\x87\xab\x96\xe9\xcc\x46\xfc\x19\x3c\x29\x95\x4b\xfb\x96\x34\xe8\xc1\xb3\xc5\x71\x84\xd5\xff\x99\x3e\x5c\x21\x4c\x13\x7e\xd1\x25\xee\x7e\xf0\x17\xfa\xbd\xe0\x4d\xa4\xfb\xe4\x3e\x8b\x74\x37\x2d\x31\x0d\xae\x29\x39\xfa\x5e\x42\x7e\xaa\x6c\x3e\x5d\xe0\xfa\xca\x6e\xb2\x70\xb7\x61\xff\x21\xa7\xd9\x7e\x5c\x67\xba\xd8\x29\xeb\x75\x31\x8e\xc0\xd1\xe7\xdf\x7f\x0f\x00\x00\xff\xff\x47\xe2\x31\xb8\x16\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 534, mode: os.FileMode(420), modTime: time.Unix(1586298350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x41\x8f\x9b\x3c\x10\x3d\x93\x5f\x31\x9a\x3b\x61\xb5\xdf\x15\xb8\xed\x77\xaa\x54\xa9\x55\xcf\xd5\x00\xc3\xc6\x5b\xdb\x50\x7b\x88\x36\x8d\xf6\xbf\x57\xc6\x24\x81\x5d\xd2\x43\x9a\x5e\x22\x67\x26\xef\xf1\x66\xde\x73\x38\x1e\xa1\xe1\x56\x59\x06\x54\xb6\xe1\x57\x84\xb7\xb7\xcd\xf1\x08\xc2\xa6\xd7\x24\x63\xb9\xfe\xbe\x63\x6a\x10\xb6\xa1\x97\x6b\x65\x7f\x80\x63\x5d\xa0\x97\x83\x66\xbf\x63\x16\x04\x39\xf4\x5c\xa0\xf0\xab\x64\xb5\xf7\x08\x3b\xc7\x6d\x81\xe1\x9c\x8d\xbc\xdb\xb1\x9a\x95\x37\xe0\x7b\x4d\x35\xef\x3a\xdd\xb0\xbb\xb0\xf8\xda\xa9\x5e\xe6\xb0\x17\xda\x53\xac\x22\x78\x57\x17\x38\xab\x64\xcf\xdc\x09\x3d\x6f\x0d\xf5\x7e\xfb\xe2\xb1\xcc\xb3\xd8\xb8\x99\xa9\x26\xc3\x8e\xee\xca\xa5\xac\x92\xfb\x10\x8e\x63\xae\xd0\x1d\x8f\xa0\x5a\xd8\x3e\x59\xaa\x34\x7f\x7e\x32\x15\x37\xa3\xa5\xb7\x3d\xa5\xe3\x40\x70\x1f\xc5\x13\xd7\xba\x66\xb6\xcd\x94\xca\x8b\xfa\xaf\x4c\xae\xde\xfd\x85\xfa\x79\xaa\xee\x32\xc2\x9c\xf0\x8f\x73\xe4\x8d\xda\x43\xad\xc9\xfb\x02\xeb\xce\x0a\x29\xcb\x0e\xcb\xcd\x06\x00\x60\xdd\xa2\xd0\x89\xfd\x05\x98\x5c\x13\x70\xc9\xfb\x62\x1a\xee\x6b\xe4\x4c\x46\x50\xdb\x39\x73\xfa\x41\x38\xa7\xca\x6a\x65\x19\xcb\x4d\x92\xe4\xca\xf6\x83\x2c\xba\x41\x95\xeb\x34\x18\x97\x7a\x93\x3e\x9e\x6e\xa7\x1f\x77\x8e\x60\xc9\x70\x81\xd1\xb2\x74\x70\x1a\x41\x35\xcb\xef\xb3\x5d\x14\xf8\x64\x85\x1d\x10\x7c\xfb\xf2\x09\xa4\x03\xc7\xe2\x14\xef\x19\x61\x4f\x7a\xe0\x02\x11\xc8\x29\x4a\x35\x55\xe1\x4f\xe1\x7f\x96\xf0\x0c\xaf\x7e\x71\x81\xff\x3d\x8c\xb7\x3d\x49\xf2\x6a\x10\xe9\xec\x49\x65\x25\x16\xcc\x21\x7d\x0c\x1f\xde\xa4\x0f\x67\x85\x43\x65\x94\x2c\xf4\xb4\x23\x5f\x39\xd2\xe6\x59\xa4\x29\xa7\xb5\x64\x61\xda\x72\x93\xe4\x59\xa3\xf6\xe5\x26\x99\x36\x39\x43\x2b\x43\xcf\x8c\x8b\xdd\x56\x5d\x73\x08\xc6\x46\xc8\x87\xd5\xb7\x5d\x27\xec\x16\x12\x0c\x0b\x9d\x11\xf1\xc1\xe7\xe3\xc9\xf4\x29\x1b\xe7\xd2\x4d\x3e\x5f\xb9\x21\xff\x2e\x03\x33\x9f\xd3\x9f\x03\xbb\x43\x9c\x7b\xa5\xbc\x9a\x88\xb1\x38\x52\x85\x60\x44\x6a\x68\x3b\x77\x25\x1a\x71\x22\xbc\x21\x10\x65\x84\x5e\xf3\xff\xbc\xf0\xf7\xe2\x1d\xfb\x41\x8b\xbf\xd8\xbd\x74\x6b\x2d\x38\x86\xfa\x5b\xf2\xf2\x5e\x45\xcb\x24\x83\xe3\x19\x70\x3d\x3e\xd3\xf9\xc3\x1b\x3b\xd0\x4e\x6f\xec\x8b\xde\xdf\x01\x00\x00\xff\xff\x39\x98\x64\x7d\xed\x07\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/index.html", size: 2029, mode: os.FileMode(420), modTime: time.Unix(1586298216, 0)}
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
	"templates/html/index.html":    templatesHtmlIndexHtml,
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
			"index.html":    &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
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
