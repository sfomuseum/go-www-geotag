// Code generated by go-bindata.
// sources:
// templates/html/editor.html
// templates/html/inc_foot.html
// templates/html/inc_head.html
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

var _templatesHtmlEditorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4b\x6f\xd4\x30\x10\x3e\x27\xbf\x62\xe4\xbb\xbb\x55\xb9\x26\xb9\x15\x24\x84\x28\xa2\x42\x1c\xd1\x64\x33\xa1\x96\xfc\x08\x8e\x93\x6a\x59\xed\x7f\x47\x7e\xe4\xb5\x9b\xf6\xc0\x25\xb2\x67\x3c\xdf\x7c\xf3\xcc\xf9\x0c\x0d\xb5\x42\x13\x30\x6a\x84\x33\x96\xc1\xe5\x92\x9f\xcf\xe0\x48\x75\x12\x1d\x01\x13\xfa\xf8\xeb\x85\xb0\x61\x70\xe7\x75\x79\xa1\x71\x84\xa3\xc4\xbe\x2f\x99\xc6\xb1\x46\xcb\x40\x34\x65\xb2\xe7\x49\x54\xe5\x79\x76\x3e\x83\x68\xc1\x58\xb8\x7b\xd4\x58\x4b\xfa\x69\x85\xa3\xf9\xf6\x4d\xe2\x91\x5e\x8c\x6c\xc8\xc2\xe5\x92\xe5\x19\x00\x40\xfc\x46\xc3\xdd\x77\xf1\x41\xd1\x1a\xab\x26\x12\xfe\xcc\x85\x96\x42\x13\xab\xf2\x2c\x2b\x84\xee\x06\xb7\xd1\x1e\x8d\x76\xd6\x48\x50\x96\xf7\x8a\x3f\x30\x70\xa7\x8e\x4a\xd6\x13\xda\xe3\x0b\x03\x8d\x8a\x4a\xd6\x2d\x8e\xf8\x9f\x81\xec\x29\xc6\xb5\x23\x5e\x89\x4a\xf6\xa8\x7d\x50\x18\x85\x01\x0a\x9c\x81\x08\x0d\xad\xcf\xe8\x88\x72\xa0\x92\x31\x40\x2b\x90\x4b\xac\x49\x96\xec\x39\xfa\x3e\x04\xc6\xf5\xe0\x9c\xd1\x13\xe5\xda\x69\x50\x27\xfe\xe0\x3f\xbd\xe2\xf7\x33\xdd\xa1\x56\xc2\xb1\x2a\x9a\x16\x87\x68\x55\xa5\x94\x1c\x7c\xa4\xd3\xa5\x11\xe3\x0d\x79\x4b\xfd\x20\x5d\xcf\xaa\xe2\xd0\x88\xb1\x9a\x53\x4d\xba\xf1\x05\x08\xe9\xdf\xc9\x7f\xaa\xda\x9c\xfa\xc4\xd5\xa3\xbf\x06\x15\xef\x71\x24\xb6\x26\x5f\x3b\xcd\x3b\x2b\x14\xda\x13\xab\x9e\x71\xa4\x2b\xae\xb3\xd3\xa9\xe8\x8b\x20\xdc\x8b\x83\xc6\xb1\xca\xf3\x10\x46\xc2\xf5\x25\x44\xa1\x69\xd3\x6d\xac\x8a\x06\x73\xcc\xd3\x63\xb4\x4d\x68\xc1\x6b\x21\xaf\x4d\x73\x0a\x9a\x1b\x93\x19\xbf\xba\x55\x5a\xf3\x1a\x5a\x6b\x27\x3d\x4f\x8f\xaa\x26\x4f\x3d\xe5\x6f\xeb\xd2\xc8\x48\xf7\x68\x24\x37\xe4\x5f\x06\x9c\x6b\xfc\xa8\xe2\x33\xed\x6c\x2e\x60\xd2\x08\x85\xbf\x97\x1c\xcf\x91\xc0\xca\x30\x85\x36\x15\x37\x7b\x03\x9f\xb7\xc6\xb8\x29\x8b\x49\xae\xc8\xe1\xb6\x2d\xd2\xd9\xc3\x4c\x78\xd7\xa5\x7b\x2f\x54\x85\xdd\x36\x4e\xaf\x08\xc2\x05\xf8\xfd\x4c\x66\xa9\x33\xde\x4a\x95\x6f\xf6\xc5\xc5\x7b\xeb\xe0\xbf\xf7\x41\xf2\x34\x58\xb9\x49\x56\xb8\xef\x6e\x80\x1f\xdf\xbf\xf8\xd1\xb7\xe4\xac\x20\x3f\x12\xbb\x83\xff\x91\x9c\xf7\xd1\x8b\xbf\x54\xb2\x0f\xf7\x0c\x0e\x91\xe5\xed\x0e\x58\x8f\xd1\x76\x05\xac\xf9\xb4\x01\xaf\x0a\xb0\x10\xfa\x64\x35\x6d\xdb\xd5\xb0\x2e\x6d\x76\x3d\x89\xd7\x95\xde\x69\x87\x54\x93\xe9\xb6\x33\x5e\x9b\xe6\x1a\x5c\x37\xb8\x65\x9c\xc8\xa1\x90\x7d\x68\xcd\x7e\x50\x3e\xaa\xea\x13\x99\xcf\xcf\x4f\x5f\x8b\xc3\x24\x58\xb7\x7e\x4b\xe8\x06\x4b\x37\x8d\x39\xe3\x24\xf9\x8a\x60\xda\x1e\xe1\x7c\xf3\x13\xf3\xdc\xd2\x4f\x6c\x09\xfb\x5f\x00\x00\x00\xff\xff\x73\x7b\x9e\x6d\x01\x07\x00\x00")

func templatesHtmlEditorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlEditorHtml,
		"templates/html/editor.html",
	)
}

func templatesHtmlEditorHtml() (*asset, error) {
	bytes, err := templatesHtmlEditorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/editor.html", size: 1793, mode: os.FileMode(420), modTime: time.Unix(1590000322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xcf\x8e\xd3\x30\x10\xc6\xef\x7d\x8a\xc1\x67\x9a\x88\x1b\x87\xb8\x17\x54\x21\x24\x24\xf6\x86\xc4\x05\x4d\xed\x49\xe2\xc5\xb1\x83\x3d\xdb\x6e\x89\xfa\xee\xc8\xf9\xb3\xcd\x2e\xad\x84\xda\x9c\x1a\x8f\xbf\xf9\x3c\xf3\x93\x3b\xee\x3a\xd0\x54\x1a\x47\x20\x8c\x53\x3f\x6b\x42\x2d\xe0\x74\x5a\x15\xef\xb4\x57\x7c\x6c\x09\x6a\x6e\xec\x66\x55\xa4\x1f\xb0\xe8\x2a\x29\xc8\x89\xcd\x0a\xa0\x48\xe2\xf4\x01\x50\x34\xc4\x08\xaa\xc6\x10\x89\xa5\x78\xe2\x72\xfd\x51\x8c\x5b\x6c\xd8\xd2\xe6\x33\x79\xc6\xaa\xc8\x87\xd5\x2c\xc9\x61\x43\x52\xec\x0d\x1d\x5a\x1f\x58\x80\xf2\x8e\xc9\xb1\x14\x07\xa3\xb9\x96\x9a\xf6\x46\xd1\xba\x5f\xbc\x07\xe3\x0c\x1b\xb4\xeb\xa8\xd0\x92\xfc\x30\x1d\x11\x55\x30\x2d\x43\x2a\x57\x0a\xa6\x67\xce\x1f\x71\x8f\x43\x54\x40\x0c\x4a\x8a\x59\x24\x3f\xd4\x3e\x7a\x57\x9a\x10\x39\x73\xc4\xd9\x63\x14\x9b\x22\x1f\x36\x07\xc7\xae\x03\x53\x02\xfd\x86\xec\x01\x2b\xfa\x94\x4a\x7a\x66\x10\x5b\x6d\xd8\x87\x9e\xcf\x2d\xe7\x56\x3d\x83\xac\xc1\x36\x5e\x38\xf3\x46\x37\x85\x0d\x05\xbc\xe4\x67\x8d\xfb\x05\x81\xac\x14\x91\x8f\x96\x62\x4d\xc4\x62\x6e\xaf\x62\x14\x50\x07\x2a\xa5\x48\xdf\x39\xf5\xed\x65\x7d\x38\x9f\x83\xc8\xb6\x0e\x77\x96\xbe\x6d\x9b\x1d\xe9\x3b\xbb\xf7\x94\x4c\x96\xeb\x7f\xf4\x4b\x37\x63\x11\x08\xa3\xdf\x1b\x08\xe4\x5e\xfa\x7e\x85\xe4\xc1\xa2\xa2\xda\x5b\x4d\xe1\x4e\x2e\xed\xd9\x69\x39\x38\x73\xd3\xc5\x08\xcd\x4d\xff\x13\xd3\xf7\x60\xf8\x6e\x42\xd8\x9a\xe5\xc8\x1c\xfa\x8a\x16\xf7\x5b\x0c\xf2\xe8\x77\x9d\xef\x7d\xe3\xe2\x5a\x9d\xb7\x99\x8e\x83\xe3\x8a\xe9\xbc\xee\x22\x9f\x1e\x8d\x62\xe7\xf5\x11\x34\x32\xae\xa7\xa9\x6e\x91\x0d\x3f\x69\x92\xa2\xeb\x20\xfb\x32\x44\xbf\x8e\x41\x38\x9d\xc4\x1b\xb9\x77\xd5\x05\xfd\x14\xfd\x37\xe1\x8f\xf7\xcd\x2b\xed\x0f\xef\x9b\xb3\x6c\x76\xaf\xd7\xe4\x74\xeb\x4d\x7a\x86\x92\x7c\xf6\x2f\xdf\x8e\x1b\xe7\xb4\x61\x60\xbc\x64\xc4\x21\x65\x98\x95\x93\x3a\x26\xf9\x66\x75\x26\xf1\x37\x00\x00\xff\xff\x4e\xd7\xf6\x5d\x77\x07\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 1911, mode: os.FileMode(420), modTime: time.Unix(1588960763, 0)}
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
	"templates/html/editor.html": templatesHtmlEditorHtml,
	"templates/html/inc_foot.html": templatesHtmlInc_footHtml,
	"templates/html/inc_head.html": templatesHtmlInc_headHtml,
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
			"editor.html": &bintree{templatesHtmlEditorHtml, map[string]*bintree{}},
			"inc_foot.html": &bintree{templatesHtmlInc_footHtml, map[string]*bintree{}},
			"inc_head.html": &bintree{templatesHtmlInc_headHtml, map[string]*bintree{}},
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

