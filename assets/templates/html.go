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

var _templatesHtmlEditorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x41\x8f\x9b\x3c\x10\x3d\x93\x5f\x61\xcd\x9d\xb0\xda\xef\x0a\xdc\xf6\x3b\x55\x6a\xd5\x55\xd5\x63\x35\xc0\xb0\xf1\xd6\x36\xd4\x1e\x68\xd3\x68\xff\x7b\x65\x4c\x02\xec\xd2\x56\x4a\xe8\x25\x72\x66\x98\xe7\x79\x6f\x9e\x2d\x9f\x4e\xa2\xa2\x5a\x1a\x12\x40\x95\xe4\xc6\x82\x78\x79\xd9\x9d\x4e\x82\x49\xb7\x0a\x99\x04\x48\x53\x7e\x39\x10\x56\x20\xf6\x3e\x97\xba\xd2\xca\x96\x05\x1f\x5b\xca\x80\xe9\x07\x27\xcf\xd8\x63\x88\x82\x70\xb6\xcc\x60\x16\x49\x9e\xa8\x61\x7c\xda\x6b\x6c\xdd\xfe\xd9\x41\x9e\x26\x21\x91\x5f\x8b\x54\xa2\x26\x8b\x9b\x62\x49\x23\x79\x1b\xc0\x20\xe2\x2a\xa0\x92\xe6\xab\xb0\xa4\x32\x70\x7c\x54\xe4\x0e\x44\x0c\x73\xfc\xd2\x39\x10\x07\x4b\x75\x06\x7e\x9d\x8c\x58\x43\x38\xc9\xfd\x50\x64\x2d\xf6\x0f\x06\x0b\x45\xef\x1f\x74\x41\xd5\x0d\xf3\x68\xc8\x03\x6c\x43\x7a\xc4\xda\x84\xf4\x88\x35\x23\x4d\xa6\x1a\x3d\x39\xd1\xff\xa0\xb0\xa4\x43\xa3\x2a\xb2\x37\x68\xd0\x4e\x28\xdb\x08\x31\x07\xdc\x44\x8d\x39\xe0\x5f\x24\xf9\x6c\x25\xdf\xa4\x06\xb6\x72\x1b\x15\xbe\x0f\x9d\x6c\x8a\xb5\x89\x98\x23\xd6\x5b\x1d\xd3\x4a\xf6\xa2\x54\xe8\x5c\x06\x65\x63\x18\xa5\x21\x0b\xf9\x6e\x27\x84\x10\xeb\xc7\xce\x67\x42\x7e\x51\x8c\xb6\xf2\x75\xd1\xeb\x60\xec\x2f\xd0\x80\x19\x0d\x45\x75\x63\xf5\xf9\x03\xbf\x8e\xa5\x51\xd2\x10\xe4\xbb\x28\x4a\xa5\x69\x3b\x5e\x64\x7d\x57\xb6\x51\x42\xdb\xd8\xe9\xf8\xfe\x4c\xd3\x11\xda\xf2\x00\xc2\xa0\xa6\x0c\xc2\xd1\x89\x3b\xab\x40\xc8\x6a\xf9\x7f\x66\xa4\x0c\x1e\x8c\x77\x0a\x8a\x4f\x1f\xdf\x09\x6e\x84\x25\xb6\x92\x7a\x02\xd1\xa3\xea\x28\x03\x10\x68\x25\xc6\x0a\x0b\xaf\xee\xff\xc4\x7e\x0f\x27\x7f\x52\x06\xff\xdd\x0d\xda\x45\x51\x5a\x74\xcc\x8d\x39\x77\x59\xb0\x11\xfa\x18\xdf\xfb\x1f\xa7\xe3\xbb\x4b\x87\x5d\xa1\x25\x2f\xfa\xa9\x07\xbc\x7c\x80\x4d\x93\x00\x93\x8f\xb2\x24\x9e\x6d\xbe\x8b\xd2\xa4\x92\x7d\xbe\x8b\x46\x25\x67\xd5\x52\xe3\x13\xc1\x42\xdb\xa2\xa9\x8e\xde\x19\xa1\xe4\x8d\xf4\x75\xd3\x30\xd9\x45\x0b\x9a\x18\x2f\x15\x61\xe3\xcb\xf2\x3c\xf4\xd1\x1b\x97\xd0\x55\x73\xfe\xd3\xa5\xf5\xef\x8c\x30\x1b\x76\xfc\xad\x23\x7b\x0c\xe4\x57\xc2\xab\xb6\x18\x82\x03\x94\x77\x47\x80\x16\xb5\x7f\x1a\xac\xfa\xe3\x31\xec\x7d\x85\x2b\xf2\x50\xfa\x3b\x13\x5c\x54\x7f\xdd\xbc\x25\xd7\x29\x76\xd3\xcc\x97\x23\x5b\x73\x8f\xc6\xf6\x7a\xd3\x74\xdc\x76\x0c\xf9\xca\x48\xa7\x4b\x37\x74\x3b\xd2\xf7\x55\xe1\xb6\x89\x1d\xf6\x93\x5d\xbd\x1e\x05\x9b\xb8\xb5\x52\xa3\x3d\x42\xfe\x88\x3d\xbd\xa2\x3f\xe3\xb1\x10\xa0\x26\xe4\xce\xd2\xac\xe7\x75\xfb\x8e\xeb\x37\x4f\x38\xcf\x68\x7c\xc2\x4d\x5b\xfc\x0a\x00\x00\xff\xff\x25\xbe\x01\x9e\xff\x09\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/editor.html", size: 2559, mode: os.FileMode(420), modTime: time.Unix(1586547389, 0)}
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

