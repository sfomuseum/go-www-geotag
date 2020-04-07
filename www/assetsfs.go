// Code generated by go-bindata.
// sources:
// static/javascript/geotag.camera.init.js
// static/javascript/geotag.camera.js
// static/javascript/geotag.maps.init.js
// static/javascript/geotag.maps.js
// static/javascript/geotag.oembed.init.js
// static/javascript/geotag.oembed.js
// static/javascript/geotag.placeholder.init.js
// static/javascript/geotag.placeholder.js
// static/css/index.css
// static/css/placeholder.css
// DO NOT EDIT!

package www

import (
	"github.com/whosonfirst/go-bindata-assetfs"
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

var _staticJavascriptGeotagCameraInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8e\xb1\x4a\x04\x41\x0c\x86\xfb\x7d\x8a\x30\xd5\x2c\xec\xcd\x0b\x1c\x16\x16\x16\x07\x62\x65\x2f\xb9\x49\x1c\x02\x73\x09\xcc\x66\x76\x11\xf1\xdd\x65\x44\xb6\x10\x9b\x2b\x43\x3e\xbe\xef\xdf\x45\xc9\xf6\x84\x44\x4f\x1b\xab\x3f\xcb\xea\xac\xdc\x62\xa8\x86\x14\x16\x78\xef\x9a\x5d\x4c\x61\xdc\x91\x07\x33\x7f\x4e\x13\x00\xc0\x86\x0d\x44\xc5\xdf\x2a\x3a\x3c\x00\x59\xee\x37\x56\x4f\x57\xa3\x8f\x54\xd8\x1f\xdd\x9b\x5c\xbb\x73\x0c\x84\x8e\xa7\xc1\x0a\xd6\x53\x45\x17\xef\xc4\x61\x3e\xff\xf1\x98\xde\xe5\x31\x2d\x87\xe8\xc7\x94\x4d\x57\xab\x9c\xaa\x95\x18\x2e\x2f\x97\xd7\xb0\x1c\x0b\x97\xa3\xf1\x9b\x2d\x6c\x8e\x25\x65\xbc\x71\xc3\x34\x9e\xf1\x5f\xf6\x6b\x3e\x4f\xdf\x01\x00\x00\xff\xff\x18\xcc\xe1\xdd\x26\x01\x00\x00")

func staticJavascriptGeotagCameraInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagCameraInitJs,
		"static/javascript/geotag.camera.init.js",
	)
}

func staticJavascriptGeotagCameraInitJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagCameraInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.camera.init.js", size: 294, mode: os.FileMode(420), modTime: time.Unix(1586281466, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagCameraJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xd1\x8b\xda\x40\x10\xc6\x9f\x77\xff\x8a\xa9\x2f\x49\x20\x04\xdb\x47\x83\x0f\x22\xda\x1e\x78\xe5\xb8\x0a\x7d\x10\x29\x8b\x5d\xd3\xc0\xba\x13\x36\x6b\xe1\xf0\xfc\xdf\xcb\xcc\x6e\x4c\x3c\x3d\xe8\xbd\x48\xcc\xfc\xbe\x6f\x67\xbe\xcc\xfe\x55\x0e\x2a\x8d\x5e\x55\x30\xed\x1e\x5e\x5f\xe1\x74\x2e\xa5\x0c\x7f\x8b\x9d\x3a\x68\xa7\x60\x0a\xe9\xfe\x68\x77\xbe\x46\x9b\x66\x27\x29\x01\x00\x48\xfd\xeb\x52\xb7\x47\x63\xca\xbe\xd0\x6a\xb3\x87\x29\x9c\xa4\x14\x49\x6d\x6b\x9f\x4c\xe0\x62\x60\x94\xcf\xc1\xa0\x25\x23\x41\x02\xa3\x3c\x4c\xa1\x51\xae\xd5\x4b\x83\xca\x13\x91\x95\xb1\x86\xf6\x4d\x0d\x6d\x57\x0b\xbf\xf5\x1e\xd2\x4f\xe4\x91\x9d\xa4\x10\xc1\x6b\x5c\x8c\x23\x73\x96\x57\x14\x9f\x2a\x44\x70\xbd\xa5\x76\x68\x5b\x34\xba\x30\x58\xa5\xa3\x87\xef\x0f\x6b\x98\xcf\x1e\x17\xcf\xb3\x51\x0e\x97\xae\xa3\x84\x26\x2c\x68\xb4\x39\x47\x90\x0e\xeb\xe7\x5c\x0a\x29\x92\x4a\xc7\xe2\x70\x7a\x3a\x9f\xf4\x4e\xfb\xa3\xb3\x5d\x82\x41\x14\xc3\xba\x15\xdd\x44\x76\xdd\x28\xcc\x1e\xe1\xdb\xe2\x79\x71\xa7\x4d\xfa\x18\xe1\x88\x27\xac\x2d\x65\xb3\x31\x68\x99\xdb\xf6\x84\x57\xae\xd2\xfe\x3d\xa2\xe7\x1a\x22\x5a\xfe\xb0\x42\xf8\x97\x46\x4f\x20\x59\x6a\xe5\x8f\x4e\x27\xb9\x14\xa2\x71\xd8\x68\xe7\x6b\xdd\x4e\x18\x21\x9d\xb2\x95\xd1\x13\xf8\x32\x96\x82\x83\x11\x95\xc6\x83\xf6\xee\xa5\x47\xa2\xd3\xd7\x58\x98\xa3\x31\x9a\x07\x67\x53\x22\xa2\x86\x8d\x37\x52\x08\x41\xd2\xa1\x96\x7b\x67\x3c\xe6\x83\xee\x77\x6d\x95\x27\xc1\x60\x7e\xaa\x73\x13\x1f\x32\x18\xc4\xc3\x06\xb1\xa9\xad\xe4\x67\x5e\xa0\xab\x95\xec\xaf\xc5\xaa\x08\x17\xe9\xe9\x0f\x7a\x8c\xb7\x29\x0d\x29\xe6\x3c\xfe\xa1\xb6\xb3\x10\xd0\xe7\x71\xb4\x1a\x6c\x50\xab\xfd\x4a\xf9\x15\xda\xff\x5f\x86\x1f\x8b\xf5\x9d\x35\xb8\xea\xab\x68\xbb\xc5\x24\x73\x5b\xa5\x9b\x0e\xdf\x76\xfc\x80\x5c\xf3\xf0\xf7\xc8\x3b\xa7\xff\x5c\x2f\x47\xf9\x45\x5d\xbd\x39\x27\xcb\xde\xdf\x7d\x21\x63\x8e\x72\x40\xd0\x25\x2b\xf9\x85\x3c\x67\x69\x56\xca\x7f\x01\x00\x00\xff\xff\x53\x70\xca\x83\xb3\x04\x00\x00")

func staticJavascriptGeotagCameraJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagCameraJs,
		"static/javascript/geotag.camera.js",
	)
}

func staticJavascriptGeotagCameraJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagCameraJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.camera.js", size: 1203, mode: os.FileMode(420), modTime: time.Unix(1586281734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagMapsInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x8f\xda\x48\x10\x3d\xc3\xaf\xa8\xf5\x65\x8c\x04\x8e\xb4\xc7\x41\x1c\xb2\x51\x56\xc9\x8a\x49\x56\xbb\x64\x2f\x51\x64\x15\xee\xb2\x69\x4d\xbb\xda\xea\x2e\x87\x90\xd1\xfc\xf7\x55\xfb\x0b\xc3\x0c\x43\xe6\x82\xb0\xfb\xd5\xab\xf7\xea\xa3\xbd\xd7\xac\xec\x3e\x41\xa5\xde\x7f\x27\x96\xb5\xf6\x42\x4c\x2e\x8e\x8c\x45\x15\xcd\x21\xaf\x39\x13\x6d\x19\xc2\x73\x4c\x01\x33\x7b\x98\x4e\x01\x00\xbe\xa3\x03\xac\x74\x7a\x4f\x07\x58\x81\xb2\x59\x5d\x12\x4b\xb2\xb5\xea\x90\x14\x24\x6f\x45\x9c\xde\xd6\x42\x71\xa4\x50\x70\xc1\xf4\x43\x7e\x12\x2f\xb0\xd2\x8b\x7b\x3a\x44\xb3\xe5\xc0\xe2\xe5\x60\x28\xad\x9d\x79\x0d\x4f\x13\xb4\xa8\x9d\x19\x33\x89\x7e\x3d\x51\x88\xe9\x78\x02\x4d\x6b\x4e\xe7\x10\xff\xd6\xdb\x9b\x3d\x4c\x27\x99\x65\x6f\x0d\x25\xc6\x16\x71\x74\xa7\xbd\xd7\x5c\xc0\xdb\xbf\x3f\x42\xe7\x65\xe2\x48\x6a\xc7\xad\x94\xc7\xe6\x77\x44\x34\x38\xbc\x48\xd5\x20\xe0\xcb\x3f\xeb\xeb\x64\xbd\xc9\x8b\x5c\x01\x70\x8d\x2a\x54\x4b\xb3\x96\xd4\xa0\xfc\x4a\xb5\x02\x56\xa3\x59\x18\x14\x2d\xb5\xa2\xc0\x3d\xd2\xd4\x53\x5d\xd4\xd4\xc5\xc3\x38\xfe\xba\x36\xcb\xaf\xd2\x66\xb9\xb8\x2c\xce\xf2\x75\x71\x63\x82\xab\xea\x7e\x5a\x5b\xbe\x46\x5e\xc0\x3f\x33\x64\x03\xd7\x55\x79\x1d\xc1\x8b\xca\x4a\xac\x52\x3a\x99\xff\x82\xe4\xbd\xa1\xf0\xf7\x8f\xc3\x47\x15\x47\x25\x56\x67\xf5\x69\x63\x2e\xa6\x2f\xb1\x02\x6a\x19\xa2\xd9\x72\x72\x9e\xfe\x24\x35\xba\xc2\xc3\x0a\x1e\xa6\x93\xa8\x5b\x9e\xe8\xb6\x5f\xa3\xf9\x74\x12\x0d\x8b\x10\xdd\x1e\x97\x22\x1c\xf4\x43\x1d\xdd\x0e\xf3\x3d\x6f\x13\x74\x52\xdf\xbc\x81\x3d\x01\x13\x29\x10\x0b\xca\x82\xec\xb4\x87\x74\x4b\xb9\x75\x94\xc2\x06\xb9\x70\x58\x82\x17\x74\xe2\x41\xdc\xa1\xd9\x04\x0b\xca\xe1\x3e\x60\xb9\xf0\x3d\x4f\xeb\x37\x69\xf2\x27\x4a\xfb\xca\x60\xb8\xc2\xa2\xad\xb1\xd9\x7d\xb4\x7c\x52\x50\x58\x41\x41\x56\xb0\x48\x4a\xac\x7c\x28\xe8\x1d\x56\x71\xcb\x32\x1f\x6c\x3f\xa9\xe9\x93\x82\x7e\x61\xdc\x1a\x0a\xa2\x34\x7b\x41\x16\x8d\x42\xd0\xf5\xe3\x52\x55\x77\xe8\x77\xb0\x02\xa6\x3d\xac\x93\x0f\xe8\x77\x21\x71\x9f\xab\x07\xa4\x5e\x1c\xac\xc0\xd8\x0c\xc3\x6d\x9d\x84\x77\x23\x39\x3d\x24\x5c\xdd\x93\x10\x53\xa1\xf3\xa4\xce\x7c\x35\x2f\x9b\x0c\x03\x7e\x39\x9d\x4e\x02\x41\x8b\x0f\x86\x1a\xca\xe3\xbd\xd1\x1e\x7c\xbd\xe9\xd7\xfa\xe6\xdb\x72\x8c\x69\xf6\x77\xc0\xf4\xdb\x75\x0a\xea\xd6\xa8\x47\x85\xc7\x06\xf0\x78\x3e\xde\x25\x56\x89\x27\xf9\x4f\xd3\x3e\xfe\xda\x6b\x98\x0f\x99\xbe\xcd\x47\xab\x34\x0c\xcd\xb1\x4e\x19\x96\xe4\xf0\xe8\xb9\x7d\x0e\xdd\x7c\xd7\xfc\x8b\xbb\xef\x48\xf7\x1e\x95\xda\xd8\xae\xd6\x83\x84\xee\xcc\xf7\x31\x6b\x94\x35\x17\xcf\xaa\x39\x65\xf3\x24\x1b\x74\x05\xc9\x8b\x11\x43\x9e\x41\xb4\xe5\xb4\xae\x54\x18\x93\xd5\xf0\x2d\x8e\x29\x34\x22\x9c\xe6\xb0\x82\xa3\x8d\x3f\x35\x19\xf5\x39\x6f\xea\xd3\x74\x2e\x40\x5e\xbe\x0b\x72\x42\xa9\x5d\x7b\x5f\x36\x8d\x6e\xaf\x81\x90\x9d\x4c\xa2\x99\xc9\x7d\xd8\xdc\xad\xc3\x72\x44\xcb\xe1\x3d\x56\x15\xb1\x7a\xb7\xd3\x46\xc5\x8e\x58\x91\x4b\x3b\xa2\x38\x9f\xcd\x8e\xad\x1b\x0d\xe9\x29\x6c\x6c\x26\x1f\x86\x92\x38\x83\x15\xfc\xf5\xef\xe7\x4f\x89\x17\xa7\xb9\xd0\xf9\x21\xce\xe7\xc0\xb5\x31\x73\xf8\x3d\x10\x37\xb3\xdb\xc4\x0f\x96\x32\x47\x28\xd4\xb9\x8a\xa3\xaa\x75\x33\xa9\x1c\x9d\xe8\x3c\xc3\x6f\xe8\x87\x7c\xb2\x8a\x62\xe2\x6c\xd6\xb8\x6f\xd7\x2f\xb0\x2f\x7b\xf5\xe7\x6d\xb7\x1c\xdf\x64\x3b\xe4\x82\x6e\xe6\xc7\xce\x9c\x36\x3a\x60\x34\x57\xb5\x9c\x41\x1e\x67\xcb\xe9\xff\x01\x00\x00\xff\xff\xc0\x90\x1f\x83\x6d\x09\x00\x00")

func staticJavascriptGeotagMapsInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagMapsInitJs,
		"static/javascript/geotag.maps.init.js",
	)
}

func staticJavascriptGeotagMapsInitJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagMapsInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.maps.init.js", size: 2413, mode: os.FileMode(420), modTime: time.Unix(1586281746, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagMapsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x6d\x8b\xdc\x36\x10\xfe\xec\xfd\x15\x13\x17\x6a\x3b\x2c\x76\x2e\xa5\x50\xec\x38\x90\x94\x96\x06\xae\x39\xb8\xe6\x3e\x1d\xcb\xa2\xb5\x67\x6d\x37\xb2\xe4\x4a\x72\xee\x75\xff\x7b\x19\xf9\x65\xe5\xbb\x3d\x42\xfb\xe5\xce\x1a\xcd\x3c\xf3\xe8\xd1\x8c\x66\xbf\x31\x05\x15\x4a\xc3\x2a\xc8\xa7\x8f\xc7\x47\x78\x38\x64\xab\xd5\xb0\x8c\x5b\xd6\x69\xc8\x21\xdc\xf7\xa2\x30\x8d\x14\x61\xf4\xb0\x5a\x01\x00\x50\x2c\x33\x46\x35\xbb\x9e\xec\x90\x43\xf0\x8e\x41\xad\x70\x9f\xfb\xb5\x31\x9d\x4e\x93\xa4\x6a\x4c\xdd\xef\xe2\x42\xb6\x89\x61\xa2\x52\xac\xd5\x3e\x18\xa6\x2a\x34\xb9\xbf\xdd\x71\x26\xbe\xfa\xef\xbf\x0c\x3b\xef\x12\xf6\x1e\x1e\x61\x81\x91\x26\xc9\xcd\xcd\x4d\x2c\x3b\x14\xda\x28\x44\xd3\xb2\x2e\x96\xaa\x4a\x0a\xd9\xdd\xa9\xa6\xaa\xcd\x73\xb8\x1f\x69\x2f\x83\x8b\x0e\xc5\x5f\x36\xe6\x4f\xd6\x41\x21\xc5\xc0\x54\x2a\x7d\x2a\x91\x1e\x33\x09\xbc\x35\xf7\x28\x6c\x8e\xe7\xd0\x9f\x87\x5d\x02\x08\x32\x12\x61\x16\x62\x54\xc9\x0a\x37\xd9\x34\xf2\x3d\xd9\x56\x2b\x2f\xe8\x98\xd2\xf8\x07\xd3\x75\x90\xc2\xac\x64\xcd\x74\xbd\xd5\x46\x91\xa2\x1e\x45\x35\x7b\x98\x8d\x71\x23\x4a\xbc\xbd\xd8\x87\xc1\x0f\x41\x04\x79\x9e\xc3\x9b\x08\x1e\x56\x9e\x37\x39\x40\x0e\xb3\xaf\xee\x77\xda\xa8\xf0\x2c\xca\x06\xa0\xc3\x08\x48\x34\x38\x33\x99\xb3\x92\xc2\x59\xdd\x4b\xd9\x8e\xcb\xa3\x91\xa9\x4a\x2f\xc0\x3b\xde\x98\xd0\x4f\xfc\x68\xe1\x4a\x64\xc9\x35\xe6\x28\x2a\x53\xc3\xab\x1c\x7e\x8a\x88\x60\x21\x85\x96\x1c\x63\x2e\xab\xd0\xbf\x12\x0a\x0b\x59\x89\xe6\x1e\x4b\x0b\x09\xda\xa8\x46\x54\x16\xcc\x53\x68\x7a\x25\x40\xf4\x9c\xcf\xcc\x9d\x14\x44\x0f\x72\x4b\xe8\xfa\xcd\x66\xf4\xe0\xcc\x4c\xb6\xb3\xd9\x66\xeb\xcf\xda\xde\x6e\x32\xcf\xf3\x4e\xa0\xd8\x2b\xf8\x24\x4c\x48\xeb\x35\x9c\xbd\x89\x16\x80\x76\xfb\x77\x2e\x99\x09\x39\x33\xd1\x02\xd8\xdd\x93\x22\xca\x96\xf8\x24\x44\xa3\x3f\xb3\xcf\x16\x39\xa2\x06\x1a\x96\x84\xe3\xac\xa4\x88\x86\x1b\x5c\x08\xf4\x49\x7c\x63\xbc\x29\x2d\xcb\x84\x33\x93\x70\x29\xfc\x35\x0c\x24\x39\x33\x6b\xb0\x29\x5f\x10\xeb\x78\x6b\x96\x63\x69\xeb\xcd\xf3\x02\xce\x4c\x63\xfa\x12\x83\xd4\x62\x58\x93\x14\xd5\x6c\x93\xc2\xda\x28\x4b\x90\x0e\xc9\x46\xc8\x6c\xc4\x1c\xb3\x0d\xb0\xd9\xca\x3b\xac\xa9\x8e\x2b\xdb\x4e\x1f\xef\x3e\x95\x6e\x21\xb7\xac\xdb\x36\xe5\xda\x5e\xc0\x5c\xcc\x63\x5b\x6c\x91\x43\x0e\xa5\x2c\xfa\x16\x85\x89\x2b\x34\xbf\x71\xa4\x4f\x02\x09\xfd\x96\x75\x54\x0a\x47\x25\x5f\x8d\x41\xb6\x94\x5e\x3e\xf2\xb8\x43\x3d\x16\x0f\xa4\xc2\x31\x6e\xe0\xea\xcd\x64\x9f\x12\x45\xfe\x84\xe8\x90\x75\x34\x79\xde\x58\xfe\xd4\xca\xcf\x2b\x72\x3a\x54\x43\x4a\x0f\x68\x94\xfe\xc3\xf8\x0e\x62\xe8\x37\xe5\xa9\xf3\x34\xa5\x73\x9e\x53\xc0\xe4\x4a\x8f\xc8\xf5\xe0\xbd\x71\x8f\xef\xda\x9f\xc8\x70\xe4\x7e\xed\xb3\xae\xd9\x7e\xc5\x3b\x7f\xf3\x1d\xe9\x6c\x8f\x0f\xbe\x53\xd3\x1c\x63\x9f\xbd\x06\xe3\xbb\xbd\x95\x9d\x21\x59\x26\xc1\xc7\x47\xfb\xa2\x23\x5d\xb5\x7d\x08\xa2\xcc\x3b\x15\x79\xce\xee\x90\x1e\xab\x31\x22\xe6\xc8\xf6\x1c\x8d\x35\x87\x2e\xfa\x2c\xdb\x28\x32\xe4\x70\x4e\xe3\x67\x2a\x12\x87\x98\x0b\x1d\xb3\xb2\xfc\x22\x49\xbc\x05\xc0\x72\x34\x4d\xb4\x3f\x1c\xad\xe1\xd4\xe3\x34\x53\x1c\xef\x5f\x69\x54\x48\x4e\xb0\xae\xb7\xe3\x31\xe7\x71\x2f\x66\xa8\x87\x6c\x51\x9d\x83\x61\xee\x9c\xa5\x66\x6e\x5d\xfe\xb7\x7a\xfc\xff\x57\x9f\xbc\x3e\x0a\xa4\x0b\x14\xf8\x05\x6f\xed\x6b\x7a\xc3\x1a\x03\x7b\x34\x45\x1d\x0a\xbc\x81\x4b\xfc\xa7\x47\x6d\xc2\x60\x9a\x8d\x5a\xb6\x37\x58\xa3\x42\x3b\xcb\x6d\x68\x7c\xdf\x74\xc1\x1a\x1e\xa0\x46\x56\xa2\xd2\x29\x3c\x40\xf0\xa1\x28\xb0\x33\x41\x0a\x01\xeb\x3a\xde\x14\x8c\x8e\x97\x90\x27\x1c\xe0\x10\x45\xb1\xa9\x51\x84\x0a\xf2\xf7\xa0\x62\x83\xb7\x26\x8c\xa2\xec\x09\xa7\xab\xcb\x73\xc8\xe1\xea\xf2\x3c\x2e\x14\x32\x83\x17\xbb\xbf\xb1\x30\x57\x97\xe7\x96\xda\x47\x2e\x77\xe1\xf5\x4c\x7e\x33\xc7\x0f\xa4\xb8\x64\x65\x38\xc1\x10\xbb\x1d\xd3\xb8\xed\x98\xa9\x53\x78\xe1\x34\x01\x1c\x26\x8c\xd7\xc9\xb3\x61\xf8\x9d\x46\xb1\xac\xcd\x1d\xc7\x6d\xaf\xf8\xec\x35\x5b\x16\x7e\xa6\x79\xe2\x36\x19\xfc\x8d\x5b\xdc\x2f\xf6\x1e\xdd\xad\x3d\x5a\x6a\x3f\x6d\x11\xb4\x9d\x54\x26\x85\xeb\x95\xe7\x79\x73\xd6\xf5\xb8\xbb\x99\x3e\xb4\xec\x55\x81\x7a\x88\xf3\x5a\xd6\xdd\xa3\x18\x17\xb4\xdd\x2b\x9e\xce\xec\xd6\x8e\x75\xab\xfb\x5d\x29\x5b\xd6\x08\x9d\xc2\x75\xc0\x82\x35\x04\x3b\xfa\x53\xd0\x9f\x32\xd8\x2c\x9c\x3b\x46\x3f\xf0\x52\x78\x18\x45\x4a\x27\xf5\x0e\xb3\x9b\x4d\xa2\x9b\x7b\x4c\xe1\xe7\xb3\xb7\xb3\xb9\x65\xb7\x5b\x1a\x42\x29\x9c\xfd\x42\xb6\xc3\xc8\x9b\xfe\x1f\x4e\x0f\x26\x57\x1b\xb7\xc9\x9c\x9e\x75\x3b\x8c\x1a\xc3\x89\x76\xba\x79\x08\x9e\x52\x3c\x99\x2d\xf6\x57\x1e\xac\x0e\x11\xbd\x16\xff\x06\x00\x00\xff\xff\xd8\x2c\x6c\x47\x2e\x0b\x00\x00")

func staticJavascriptGeotagMapsJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagMapsJs,
		"static/javascript/geotag.maps.js",
	)
}

func staticJavascriptGeotagMapsJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagMapsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.maps.js", size: 2862, mode: os.FileMode(420), modTime: time.Unix(1586280485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagOembedInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcb\xce\xd3\x30\x10\x85\xd7\xf1\x53\x0c\x5e\x25\x52\x70\xf6\x44\x59\x80\xd4\x05\x02\xf4\x4b\xdd\xb1\x42\xc6\x1e\xbb\x16\x8e\xdd\xfa\xd2\x0a\x55\x7d\x77\xe4\xa4\xa4\x69\x85\x68\x37\xb9\xd8\x73\xce\x7c\x73\xe6\x64\x9c\xf4\x27\xc6\xa5\xdc\x1c\xd1\xa5\xaf\x26\x26\x74\x18\x6a\x6a\x3d\x97\xb4\x05\x95\x9d\x48\xc6\x3b\x28\xff\x35\x96\x9a\xe6\x4c\x00\x00\xa6\xc7\x91\x07\x38\xc0\x00\xd2\x8b\x3c\xa2\x4b\x4c\x63\xda\x58\x2c\x9f\x9f\x7e\x7f\x96\x35\xf5\x38\xfe\x44\xf9\x3e\x07\x4b\x9b\x9e\x4c\x1a\xa3\xa0\x7e\x07\x87\xe6\x4c\x2a\xe1\x5d\xf4\x16\x99\xf5\xba\xa6\xdf\x4c\x8c\xc6\x69\xb8\x49\x00\x67\x2b\x56\xb4\x55\xc0\x94\x83\xeb\x27\x8f\x0b\x59\xda\x9b\x17\xda\x9b\x91\x6b\x7c\x00\x30\xcf\x00\x26\xd1\x4b\x08\xea\x05\x04\x85\x49\xec\x1e\x10\xd4\x33\x84\x49\xf4\x3f\x84\x65\x11\x8a\x79\x27\xac\x11\xbf\x60\x58\x76\x56\x63\x73\x26\xa4\x2a\x84\x68\x61\x00\x64\x89\x07\x8d\xa9\x9f\xcf\x4a\xc0\x03\xa0\x65\x47\x6e\x33\xf6\xd7\x4a\xef\x7e\xc4\x2c\x04\xc6\xb8\x76\x0a\x71\x5f\x58\x4b\xa7\x3b\xde\xb7\x2f\x1f\xbf\xd3\x16\xca\x75\x4f\xaa\xcb\xca\x04\x43\xf0\xe1\x0e\x26\x84\x7f\x5a\x6c\xb6\xdb\xb7\x2d\x6d\xa1\xdc\xaf\x3d\xd0\xc9\xbd\x37\x2e\xc1\x00\x74\x97\xd2\x3e\x7e\xe8\xba\xd1\x58\x1b\x95\x41\x2b\x59\x54\x7e\xcc\x11\xf3\xc8\x7c\xd0\xdd\x9c\x57\x47\x7b\x52\x91\x4a\xa3\x4f\x5c\xb3\xf9\x8c\x4d\x19\xd6\x7f\xdd\xda\x32\x76\xbb\x9a\xb2\x5d\x60\x6f\xe9\x82\xe2\x36\xe2\x35\xe3\xf9\x45\x2e\x4d\x4f\xfe\x04\x00\x00\xff\xff\xa7\x0e\x1a\x5c\x2d\x03\x00\x00")

func staticJavascriptGeotagOembedInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagOembedInitJs,
		"static/javascript/geotag.oembed.init.js",
	)
}

func staticJavascriptGeotagOembedInitJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagOembedInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.oembed.init.js", size: 813, mode: os.FileMode(420), modTime: time.Unix(1586293262, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagOembedJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5d\x8b\xdb\x30\x10\x7c\x96\x7e\xc5\xe2\x87\x4a\x26\xc1\x3f\x20\xc6\xf4\xe1\x68\xb9\x96\x7e\xc0\xf5\x0a\x85\xa3\x04\x9f\xb3\x76\x7c\x71\x24\x65\x25\xa5\x3d\x72\xf9\xef\x45\x72\x14\xbb\xed\xf9\x49\x5a\xcf\x0c\xa3\x99\x3d\xd6\x04\x1d\x6a\x57\x77\x50\xa5\xc3\xcb\x0b\x9c\xce\x25\xe7\xe3\xb5\xd0\xb8\x7f\xc4\x0d\x54\x20\x5b\xaf\x1a\xd7\x6b\x25\xf3\x13\xe7\x00\x00\x81\x6d\x71\x68\xa1\x82\x13\xe7\x4c\xb4\xe8\x9a\xad\x58\xc1\x15\x88\x6a\x63\x74\xaf\xdc\x12\x3c\x0d\x4b\xd0\x6a\x6d\x7d\xd3\xa0\xb5\xf1\x8c\x44\x9a\xf2\x48\x4d\x6a\x87\x28\xc5\x98\xf0\x34\x88\x55\x64\x85\x5b\xab\x69\x5f\x3b\xb1\x02\xf1\x64\xb5\x12\xcb\x91\x10\x4c\x4e\xcc\xb5\x75\x04\x55\xf4\x53\x3c\xfa\x7e\xd8\xac\x0f\x1e\xe9\x39\x8c\x7b\xd5\xc9\x43\x5e\x4e\x60\x4f\x03\x54\x90\xdc\xc1\x02\xc4\x5b\x01\x8b\x51\x23\x89\x36\x5a\x59\x3d\x60\x31\xe8\x4e\x66\xef\xdf\xdd\xdf\xdc\x66\xf1\x19\x49\x67\x52\x23\x0c\xb6\x15\xfe\x82\x1f\x9f\x3f\xdd\x3a\x67\xee\xf0\xe0\xd1\x3a\xf9\x37\x94\xf0\x50\x68\x35\xe8\x3a\x84\x39\xcf\x92\x31\xce\x58\x14\xb2\xa6\x1c\x6f\x8e\x9e\x63\x0e\x91\x67\x0d\x54\xf0\xf1\xdb\xd7\x2f\x85\xa9\xc9\xa2\x74\xdb\xde\x16\x84\xd6\x68\x65\xf1\x1e\x7f\xbb\xbc\x8c\x75\xa4\x8f\x9d\x47\x91\xa6\x76\xcd\x16\x24\xe6\x49\x29\x65\x2e\x31\x38\xbb\x98\x72\x9e\x14\xb4\xf5\x60\x31\xcc\x2e\xd4\xa9\x29\x49\xd6\x4c\xfa\x29\xf7\x7f\x9f\x65\x50\xc9\xac\x43\x97\x5d\x9a\x76\xe4\x31\xbd\x3e\x00\x2c\xaa\x8d\xcc\x4b\x76\xd5\xe0\xec\xbc\xe4\x8c\x33\xf1\x7f\x55\xf3\x05\xaa\xa9\xb3\xf9\x7c\x41\x4c\xdd\x93\x85\x0a\x1e\x7e\xa6\x9e\x5a\x4d\x20\xc3\xaf\x1d\xf4\x0a\xae\x84\x18\xe8\x11\xaa\x38\x79\xd8\x45\x78\x9c\xa1\x6a\xd6\xbb\xd8\x7e\xa3\x37\xf8\xfd\xee\xc3\x8d\xde\x1b\xad\x50\x39\xb9\x8b\xb1\x24\xd0\xf1\x75\xd0\x31\x2f\xa7\xc6\x82\x9d\x11\xb6\xde\xc1\x02\xb2\x2a\x83\xc5\x48\x0e\xa0\x68\xb6\x30\xde\x6e\x65\x38\xa6\x40\xce\x9c\xcf\xb3\x1f\x51\x4f\xba\x57\x32\x7b\x93\x05\x50\x88\x86\xf1\x4b\xd4\x31\xad\x19\x3c\x2c\xf8\x65\x7a\xce\xc3\x8a\xfd\x09\x00\x00\xff\xff\xbd\x9b\x19\x70\xc4\x03\x00\x00")

func staticJavascriptGeotagOembedJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagOembedJs,
		"static/javascript/geotag.oembed.js",
	)
}

func staticJavascriptGeotagOembedJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagOembedJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.oembed.js", size: 964, mode: os.FileMode(420), modTime: time.Unix(1586293282, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagPlaceholderInitJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xc1\x6e\xdc\x36\x10\x3d\x4b\x5f\x31\xe5\x49\xc2\xae\x85\x14\x39\x2a\x7b\x48\xd3\x00\x2d\x60\xf7\x52\xa3\x17\x23\x30\x68\x71\xac\x25\x4c\x91\x32\x49\x59\x59\x18\xfb\xef\xc5\x50\x14\x57\x5a\x6f\x9c\x5c\x24\x8a\x33\x6f\xe6\xf1\xcd\x70\x34\x4a\x2d\xcc\x58\x71\x21\xbe\xbe\xa0\xf6\xd7\xd2\x79\xd4\x68\x0b\xa6\x0c\x17\x6c\x0b\x8f\x83\x6e\xbc\x34\x1a\xe8\xbb\x40\xf2\x29\x5f\xf3\x1c\x00\xe0\x85\x5b\x40\x2d\x7a\x23\xb5\x87\x1d\x08\xd3\x0c\x1d\x6a\x5f\x3d\x18\x71\xa8\x5a\xf4\x9f\xbd\xb7\xf2\x61\xf0\x58\x30\xc1\x3d\xbf\x72\xc8\x6d\xb3\xbf\x9a\x21\xac\xac\xa7\x38\xf2\x11\x8a\xdf\x52\xa4\xf2\x35\xcf\x2c\xfa\xc1\xea\x3a\x58\x8f\xe1\x99\x12\x3e\x2f\x33\xb5\xe8\xbf\x2a\xa4\xe5\x1f\x87\xbf\x45\xc1\x7a\xc5\x1b\xdc\x1b\x25\xd0\x5e\x3d\x0f\x68\x0f\x67\x39\x9e\x29\x78\x63\xb4\x33\x0a\x2b\x65\xda\x82\xdd\x48\xe7\xa4\x6e\xe1\x0d\x12\x70\x0a\x5c\x51\x88\x33\x42\x89\x8c\xfd\x55\x32\x16\xdd\xa0\xbc\x3b\xa3\x63\x7f\x89\x4e\xc4\xfe\x94\x50\x8b\xc6\xf3\xb6\x5a\x40\x2b\x87\xfe\x7e\xd6\xb5\x48\x02\x07\x54\x80\x3c\x57\x46\x3f\xe1\x61\xe8\x61\x97\x0a\x5d\x20\xd5\x37\x0b\xc5\x55\xb0\x03\xac\x3c\xb7\x2d\xfa\x7a\xda\xf3\xf8\x9d\x8a\x8d\xaa\x7a\xe1\x6a\xc0\x3a\xcf\xf2\x8c\x4e\x43\xfb\x95\x42\xdd\xfa\x3d\x7c\x82\x8f\x74\x32\xca\x30\xf3\xcc\x8e\x31\xa6\x45\x2d\xd0\xde\x4f\xbd\x70\x3f\x9f\x6d\x91\xde\xba\xbe\x7c\x25\x6c\x3e\x45\x20\x50\x63\x86\xd0\x62\xd6\xf5\x31\x47\x1d\xad\x94\x3a\x5a\x77\xf0\x81\xb2\x26\x69\xb2\xa8\x4d\x8a\x32\x5a\xde\xf7\xb8\xaa\x59\x63\x91\x7b\x8c\x65\x2b\x98\x90\x2f\x41\xdd\xd0\x71\xe1\xf9\x68\x2c\x14\x04\x96\xbb\x0f\x35\x48\xf8\x34\x71\xa9\x41\x6e\x36\x41\xa7\xe9\x50\x66\x9c\xd8\xdd\xc9\x6f\x75\xdc\x93\x82\xb6\xcc\x78\xc7\xa4\x60\x69\xb7\xf7\xf3\x6e\xa8\x93\x3f\xf4\x78\x32\x6a\xde\xe1\x6c\xa6\x35\x59\xa2\xa9\x45\xd3\xcd\x26\x5a\x9f\x40\x8a\x53\x48\xda\xbb\x63\x8a\xfb\x85\xc1\xe8\x93\xc1\xe8\xc9\xb0\x22\xfc\x33\x19\x32\x6b\x46\xea\xa1\xc5\x45\x6e\x14\x77\x8e\x6d\x61\xdd\xdd\x66\xfc\x91\x7f\xb8\xf8\xe3\xde\x38\xa3\x1f\xa5\x75\xfe\x4a\xd2\x50\x91\xe2\x3d\x77\xc5\xbd\xf4\x83\x40\xb6\xa5\xc3\xbd\xeb\x69\x74\x9b\x5c\x8d\x2e\x93\x5c\xa1\x37\x61\x37\x29\xba\x01\x06\x05\x83\x0d\x95\x64\x03\xac\x04\x5a\xf7\x3e\x38\x53\x60\xea\x0a\x2d\xbe\xec\xa5\x12\xc5\x99\x24\xb7\xf8\xdd\xff\x63\x04\x16\x21\x60\x59\x26\x8c\xd1\x8d\x92\xcd\xd3\xf9\xbd\xc9\x52\xb3\xad\xef\x4e\x16\x5b\x2a\x99\x43\x77\xa0\xba\x34\x26\xcf\xd4\x9a\x72\xce\xb8\xa9\xda\x97\x81\x49\xb7\x20\x59\x42\x84\x36\xf8\x01\x22\xe9\x57\xce\x14\x17\xc8\x8e\xf7\x53\x03\xd1\x58\xe9\x78\xef\x28\xc4\x0d\xef\xa7\xf1\xd6\xf1\x7e\xc1\x8d\xee\x61\xc7\xfb\x20\x41\xd6\xf1\x9e\xaa\xf5\x9f\xc4\xb1\xb8\x53\xdc\x87\xe2\x7c\xdb\xc2\xef\x1f\x13\xb3\xe3\x32\x51\xc3\x3b\xb4\xfc\x94\x6b\xfa\xa6\x6c\x5f\xc2\xaa\x48\xb0\xf8\x8a\x0e\x6e\x76\xb8\xe6\xfe\x5a\xb7\x8b\x5c\x09\x70\xf2\xbc\x0d\x95\xb8\xe0\xb9\x88\x1b\x5f\xb6\x72\xfe\xa0\xb0\x12\xd2\xf5\x8a\x1f\x60\x07\x4c\x1b\x8d\xac\x4e\x76\xa9\x35\xda\xbf\x6e\x6f\xae\xc9\x16\xf6\x8f\xf1\x76\xc5\x21\xb3\x6a\x2a\x6b\xc6\xf2\x6c\x1e\x4d\x43\x6a\x1e\x49\x34\x1f\xc3\x24\x25\x35\x8c\xbe\x77\x43\xd3\xa0\x7b\x3b\x16\x17\xc3\x2c\xf4\xd7\xc5\x61\x1a\x5c\x57\x43\x6c\xcd\x06\xd5\x6c\xbd\x70\xce\x07\x65\x9a\x27\x56\x47\xf0\xb1\xce\x13\x27\xb4\xd6\xd8\x55\xbf\x5b\x3b\x4f\xf9\xd5\x3f\xec\xdf\xcf\x7f\xb2\x2d\x90\xb5\x8e\x11\x2e\xfe\x99\x88\x72\xf8\x6b\x6c\x17\x27\xde\xa6\x4c\x65\xd0\xe3\x9d\x4a\xbc\x2d\x42\xd0\xb7\xce\x8f\x65\x9d\xff\x1f\x00\x00\xff\xff\xa1\xa2\x7d\x99\xd6\x08\x00\x00")

func staticJavascriptGeotagPlaceholderInitJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagPlaceholderInitJs,
		"static/javascript/geotag.placeholder.init.js",
	)
}

func staticJavascriptGeotagPlaceholderInitJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagPlaceholderInitJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.placeholder.init.js", size: 2262, mode: os.FileMode(420), modTime: time.Unix(1586281336, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJavascriptGeotagPlaceholderJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x4d\x8f\xd3\x30\x10\x3d\xdb\xbf\x62\x94\x03\xb1\xd5\x2a\xdc\x37\xb2\x38\x70\x01\xc4\x87\xb4\x80\x84\xb4\x42\x51\x48\xa6\x6d\xa8\xb1\x5d\x7f\x94\x5d\x75\xfb\xdf\x91\xed\x38\x8d\x80\x1e\x1a\x67\xfc\xde\xcb\x9b\x37\xf6\xb9\xb7\xb0\x47\xed\xfb\x3d\x88\xb2\x78\x7e\x86\xcb\xb5\xa5\x34\xbf\x36\x46\xf6\x03\x1e\xb4\x1c\xd1\x82\x00\xb6\x0b\x6a\xf0\x93\x56\x8c\x5f\x28\x05\x00\x88\x12\x1d\xaa\xd1\xe8\x49\x79\x10\xa0\x82\x94\x6d\xda\x59\xb6\x1d\xca\x1d\x08\xb8\x50\x4a\x6a\x87\x7e\x41\xd7\x77\xb0\xc8\x05\x2b\xf9\x85\x92\x48\x59\xab\x05\x2b\x5b\x4a\xae\xdb\x4c\xed\xed\x70\x58\x93\x3c\x3e\xfa\x2d\x68\xd5\xb9\x30\x0c\xe8\x5c\x5a\xa3\xb5\xda\xf2\xf4\xb5\x62\xe0\x94\xbe\x4e\x48\x1d\x09\xf5\x1d\x24\x5e\xde\x8e\x9d\xde\x70\x9d\xf3\xb1\xcb\x68\xb8\xf9\x11\x26\x39\x76\xa7\x80\xf6\x29\x96\x27\xb5\x67\x27\xde\xde\xc0\xc1\x4a\x10\x2b\xb3\x1b\x98\x1d\xbe\x7c\x55\xc3\x26\x6b\xad\xc5\x2d\x46\x1b\x0a\x7f\xc3\xb7\x0f\xef\xdf\x78\x6f\xee\xf1\x14\xd0\x79\x56\x44\xf3\xbf\xc5\x53\xa3\x95\xd4\xfd\x08\x02\xd6\x69\x13\x42\x09\x49\x42\xce\xb4\xf9\xcd\xdb\xa7\xd4\x57\xe2\x39\x03\x02\xde\x7d\xfe\xf4\xb1\x31\xbd\x75\xc8\xfc\x61\x72\x8d\x45\x67\xb4\x72\xf8\x05\x1f\x3d\xcf\x63\x29\x3f\x72\xcd\x22\x43\xef\x87\x03\x30\xe4\x45\xa9\x64\xc8\x30\x3a\x9b\x4d\xf9\x60\x15\xec\x7a\xe9\x30\xd6\x66\xea\x2d\x79\x66\x9d\xb9\xe9\x97\x64\xff\x6e\xcb\xa0\x62\xd5\x1e\x7d\xb5\x8d\xe9\x6d\xc1\xdb\x80\xa5\xfb\x08\x70\xa8\x46\xc6\x5b\xb2\x68\xa4\xd1\x13\x4a\xea\x7f\x87\xb1\x3e\x07\xbd\xdd\x3b\xbe\x1e\xb8\xe9\x27\xeb\x40\xc0\xc3\xf7\x32\x81\x9d\xb6\xc0\xe2\xd6\x11\x26\x05\x0b\x21\x05\x7a\x06\x91\x2a\x0f\xc7\x04\x4f\x35\x54\x43\x77\x04\x11\x9f\x7a\xc4\xaf\xf7\x6f\x5f\xeb\x5f\x46\x2b\x54\x9e\x1d\x53\x2c\x05\x74\xfe\x3f\xe8\xcc\xdb\xdb\xc4\xa2\x9d\x0c\xeb\x8e\xb0\x81\x4a\x54\xb0\xc9\xe4\x08\x4a\x66\x1b\x13\xdc\x81\xc5\x65\x09\xe4\x4a\xe9\x3a\xfb\x8c\xfa\xa9\x27\xc5\xaa\x17\x15\xcf\xb7\x82\x10\x3a\x47\xbd\x5c\xb8\x19\x1e\x8f\xf0\x5c\xbd\xf2\x78\xc4\xfe\x04\x00\x00\xff\xff\x31\xe9\xbe\xbb\xeb\x03\x00\x00")

func staticJavascriptGeotagPlaceholderJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptGeotagPlaceholderJs,
		"static/javascript/geotag.placeholder.js",
	)
}

func staticJavascriptGeotagPlaceholderJs() (*asset, error) {
	bytes, err := staticJavascriptGeotagPlaceholderJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/geotag.placeholder.js", size: 1003, mode: os.FileMode(420), modTime: time.Unix(1586217606, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticCssIndexCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xcb\x31\x0a\x42\x31\x0c\x00\xd0\xb9\x39\x45\x40\x1c\x95\xfe\xc1\x25\x3d\x4d\xbf\x29\x3f\x01\xd3\x48\x0c\x3a\x88\x77\x97\xce\x8f\x77\xb2\xfe\xc4\x2f\x94\xdd\x83\x47\x10\xbe\xfc\xa1\x8c\x29\x3a\x1b\x94\x8f\x72\x0a\xe1\x56\xeb\xb9\x41\x91\xa1\x87\x24\xe1\xad\xbe\xa5\xc1\x0f\xe0\x7a\xef\xc1\x2b\x5b\x8f\x43\xe7\x65\xf7\x4c\x37\xda\x62\xd8\xf2\x7f\x00\x00\x00\xff\xff\x55\x57\xc1\x4a\x5b\x00\x00\x00")

func staticCssIndexCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssIndexCss,
		"static/css/index.css",
	)
}

func staticCssIndexCss() (*asset, error) {
	bytes, err := staticCssIndexCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/index.css", size: 91, mode: os.FileMode(420), modTime: time.Unix(1586292072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticCssPlaceholderCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\x41\x8e\xb4\x20\x10\x05\xe0\x35\x9c\x82\xc4\x35\x86\xdf\x4e\x6f\x8a\xd3\x20\xa0\x90\xbf\xa4\x0c\x94\xa3\x33\x9d\xb9\xfb\xa4\x7b\xa6\x17\x26\x2e\xdf\xa2\xea\x7d\xaf\x5b\xd1\xf9\x98\x08\x43\xac\xba\xc6\xb6\x21\x37\xf5\x90\x62\xa5\x96\x39\x53\x01\xe5\xc6\x46\xb8\x71\xb4\x52\x7c\xe9\x5c\x42\x3c\x40\xdd\x8d\x31\x56\x8a\xc5\x1d\x3a\xc5\x3c\x27\x06\x35\x18\xb3\x1e\x56\x8a\x3d\x07\x4e\xa0\x6e\x7f\x71\xa4\x1a\x62\x05\xd5\x08\x73\x50\x9c\x72\xb1\x52\x30\xad\xa0\x6e\xfd\x60\x6a\x5c\xac\x14\x18\x27\x06\xf5\xaf\x1f\xee\xbf\x79\x74\xfe\xff\x5c\x69\x2b\x41\x7b\x42\xaa\xa0\xba\x69\x9a\xac\x14\xf4\x11\xeb\x84\xb4\x43\xf3\x95\x10\xad\x14\x21\xb7\x15\xdd\x27\x14\x2a\xd1\xca\x6f\x29\xfb\xd3\x1e\xda\x5f\x5b\x5c\x08\xb9\xcc\xa0\xfa\x77\xc1\x0b\xa5\x47\x62\xa6\xe5\x6c\xbb\x78\x01\xe9\xd9\xfb\xb8\x70\x75\xde\xfb\xeb\x13\x74\x8d\xb5\x4f\x19\xc3\x13\x70\xee\x7b\x5b\x7f\x02\x00\x00\xff\xff\x64\x3b\xa0\xb7\x7b\x01\x00\x00")

func staticCssPlaceholderCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssPlaceholderCss,
		"static/css/placeholder.css",
	)
}

func staticCssPlaceholderCss() (*asset, error) {
	bytes, err := staticCssPlaceholderCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/placeholder.css", size: 379, mode: os.FileMode(420), modTime: time.Unix(1586218580, 0)}
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
	"static/javascript/geotag.camera.init.js": staticJavascriptGeotagCameraInitJs,
	"static/javascript/geotag.camera.js": staticJavascriptGeotagCameraJs,
	"static/javascript/geotag.maps.init.js": staticJavascriptGeotagMapsInitJs,
	"static/javascript/geotag.maps.js": staticJavascriptGeotagMapsJs,
	"static/javascript/geotag.oembed.init.js": staticJavascriptGeotagOembedInitJs,
	"static/javascript/geotag.oembed.js": staticJavascriptGeotagOembedJs,
	"static/javascript/geotag.placeholder.init.js": staticJavascriptGeotagPlaceholderInitJs,
	"static/javascript/geotag.placeholder.js": staticJavascriptGeotagPlaceholderJs,
	"static/css/index.css": staticCssIndexCss,
	"static/css/placeholder.css": staticCssPlaceholderCss,
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
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"index.css": &bintree{staticCssIndexCss, map[string]*bintree{}},
			"placeholder.css": &bintree{staticCssPlaceholderCss, map[string]*bintree{}},
		}},
		"javascript": &bintree{nil, map[string]*bintree{
			"geotag.camera.init.js": &bintree{staticJavascriptGeotagCameraInitJs, map[string]*bintree{}},
			"geotag.camera.js": &bintree{staticJavascriptGeotagCameraJs, map[string]*bintree{}},
			"geotag.maps.init.js": &bintree{staticJavascriptGeotagMapsInitJs, map[string]*bintree{}},
			"geotag.maps.js": &bintree{staticJavascriptGeotagMapsJs, map[string]*bintree{}},
			"geotag.oembed.init.js": &bintree{staticJavascriptGeotagOembedInitJs, map[string]*bintree{}},
			"geotag.oembed.js": &bintree{staticJavascriptGeotagOembedJs, map[string]*bintree{}},
			"geotag.placeholder.init.js": &bintree{staticJavascriptGeotagPlaceholderInitJs, map[string]*bintree{}},
			"geotag.placeholder.js": &bintree{staticJavascriptGeotagPlaceholderJs, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
