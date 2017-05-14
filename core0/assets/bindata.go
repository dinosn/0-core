// Code generated by go-bindata.
// sources:
// scripts/stat.lua
// text/logo.txt
// DO NOT EDIT!

package assets

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

var _scriptsStatLua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6b\xe3\xc6\x13\x7f\xaf\x4f\x31\x08\x42\xec\x44\x76\xe4\xdc\xfd\xf3\xe2\xf8\xab\x50\xda\x70\x3d\xda\xc2\x91\x5c\x4b\xcb\x11\xc2\x46\x1a\x6b\xb7\x91\x76\x7d\xbb\x2b\x3b\xa6\xd7\xef\x5e\xf6\x41\x2b\xad\x65\x5f\x4b\xa1\x50\x21\xf0\x6a\x66\x67\x76\x66\x7e\xf3\xb0\x5e\x2c\x9e\x71\x9f\x6d\x49\xd3\x61\xa6\xb4\x9c\x71\xb1\x9b\x67\x7a\xbf\xc1\x4c\x93\x5a\x25\x49\x23\x4a\xd2\xc0\x33\xee\xa1\x80\xef\x6f\x7f\xbd\xff\xb8\x7a\xf0\x34\x2b\x03\x05\x68\xc1\xbb\xf6\x09\xe5\xec\xeb\xbb\xb7\x3f\x7f\x5c\x3d\xcc\x3d\x9f\x8b\xdd\x84\x7b\x1d\xb8\xe6\x08\x28\xc0\x52\x5f\xf5\x2a\xcd\x91\x3d\xf1\xf5\x43\xd0\x53\x85\x9d\xff\x7b\xe8\x2d\x52\x9a\x68\x74\x66\x29\x2d\x19\xaf\x97\x6b\x21\x5b\xa2\x67\xa9\xe1\xa8\x37\x67\xe6\x4d\x33\x2b\x9d\x19\xfb\xe7\x49\xb2\x58\x80\x13\xa6\x0a\xf5\x17\x65\xbd\x60\x6f\xec\x16\x0a\xf8\xfd\x0f\xff\x51\x42\x01\x69\x3a\x32\xc3\x2f\x37\x12\xcd\x3e\x89\x15\x53\xcb\x92\x34\xcd\xec\xfc\xed\xed\x87\xf3\x2c\x98\x3a\x4f\x86\xc8\x3c\x2a\x2a\xa4\x7e\x6c\xa1\x80\x59\x4b\x34\x5d\xae\x1b\x21\x6c\xf4\xe1\x0a\x5e\xe5\xf9\x1c\x2e\xdc\xcf\xa5\xf9\x99\xc8\xd1\xe3\x72\x37\x5e\xf0\xc6\x4b\xde\xe4\x79\x7f\x66\xc5\xd6\x6b\x94\xc8\x35\x23\x8d\x81\xc5\x86\xbf\x80\xf4\xdb\x34\x49\xd8\xda\x19\xaf\x29\xf2\x04\x00\x60\xb1\x80\x1a\xb5\x25\x32\xd1\x29\x07\x75\x06\x4c\x03\xbe\x30\xa5\x15\x30\x0e\x04\xe8\x33\xee\xed\x76\xe3\x76\xf9\x9b\x12\x7c\x59\x61\x29\x2a\x9c\x19\xc1\xb9\x65\x0d\x87\x1f\x7c\x6a\xd6\x62\xd2\x1f\x56\x92\xa6\xec\x1a\xa2\xd1\x98\x60\xd8\xb0\xa3\xc8\x81\x3c\x29\xd1\x74\x1a\x81\x4b\x05\x44\x22\x34\xa2\xae\xb1\x02\x5c\xd6\x4b\x78\xda\x6b\x84\x52\x74\x5c\xa3\x84\xb5\x90\xc0\x51\xef\x84\x7c\xb6\x3a\xd9\x3a\x76\x38\x78\xe6\x0f\x0c\x06\x99\xc7\x7c\x40\xe1\xf3\x79\x01\xdb\xe5\x96\x34\x11\xd3\xd8\x0a\x85\x4d\x68\xc3\xc6\x8d\x28\xe9\x58\xdb\x60\xbe\x55\xb5\x41\x09\x0a\x4b\xc1\xab\x65\xd8\x65\x72\x00\x0a\xc7\xbf\x1a\x02\x60\x58\xd8\x28\x3c\xdc\x67\x6d\x71\x5c\x5e\x85\x30\x7d\x43\xb1\x7c\x36\xbe\xed\x10\x4a\xc2\x61\xdd\x74\x8a\xda\x90\x05\xa4\x48\x5d\x4b\xac\x89\xc6\xca\xe9\x50\xcb\xa4\x0f\xc8\x76\xd9\x3e\x5a\xd3\xe1\xff\x51\x06\x1e\xc6\x66\x05\xad\x81\xd7\x2b\x62\x62\x60\x3a\xf4\xa4\xad\xeb\x83\xc2\x39\x53\x9f\xcd\xdb\x7d\x3e\x5b\x87\x37\xcd\x82\xa4\x79\x42\x25\x66\x83\x29\xae\x38\x1c\x81\x6c\x6b\xb7\x68\xc9\x8b\x5b\x68\xa1\x49\x33\x4f\x82\x96\x51\x69\xa5\x77\xef\x7f\xba\xff\x2e\xcd\x20\xfd\xd4\x61\x87\xea\x8d\xab\xdc\x96\xf1\x34\x33\x16\x8e\xa4\x82\x26\x28\x20\x8f\xa8\x5c\x4e\x48\x0d\x51\x36\xfe\xce\x9e\x09\xcb\x07\xb0\x18\x3c\x88\xb6\xb4\xe4\x25\x02\xaf\xa7\xf7\x52\xa3\xb0\x07\x70\x03\x3a\xf4\x08\x3a\xf4\x08\x3a\x54\x74\xf2\xdf\x86\x87\x1e\xc2\x43\x7b\x78\x68\x0f\x0f\xfd\x07\xf0\x18\xd3\x8f\xe0\x43\x8f\xe2\x43\xa7\xf8\xd0\x01\x1f\x7a\x80\x0f\x3d\xc4\x87\x4e\xf0\xa1\x27\xf0\xa1\x47\xf0\xa1\x93\xe2\x93\xd8\xa2\x19\x64\xb6\xde\xca\x4e\x9a\xd6\x32\x52\x65\xbb\x46\xa4\xdc\x37\x0a\xa7\xb6\x57\x13\xb4\xf4\x1a\x5a\x24\xaa\x33\x54\xae\x33\x20\xbc\x3a\xe8\x85\x64\x5b\x5f\x19\xab\x4d\x83\x6b\x19\x37\xcd\x70\x7c\xc2\x90\xd9\xc3\xfa\xd2\x0d\xa5\x28\xc9\xfd\xe2\x12\x56\x81\x4e\xb6\x75\x24\x76\xe5\x37\xf5\x09\x69\x3b\xd1\x57\x21\xad\xa3\x3c\x1c\x72\x3d\x1c\x35\x0e\x95\x69\xc3\xd6\x62\xe3\x81\x4d\xd7\x0d\x4a\x26\x2a\xcf\xf7\x28\x32\x05\x5c\x68\x90\xf8\xa9\x63\x12\x2b\xd8\x89\xae\xa9\x2c\x69\x23\xc5\x96\x55\x08\xa4\xaa\x98\xc9\x70\xd2\x00\xe3\x6b\x91\x1c\x66\xcb\xb0\x8e\x7c\xa6\xbd\xcf\x34\xf6\x99\x06\x9f\xe9\xc8\x67\x7a\xcc\x67\x7a\xcc\x67\xfa\x25\x9f\x49\xb3\x23\x7b\x05\x12\x15\x6a\x77\x8f\x61\x1c\x4a\xa2\x10\xc4\x1a\x4a\x4a\x78\x8d\x4b\x6f\x87\xbf\xe5\xd8\xfb\xd5\x61\x72\x31\xee\xca\x68\x3c\x2d\x89\x26\x61\xc0\x22\xb7\x03\xd6\x4f\xd7\xf1\x5d\xe3\x3e\xba\x6b\x64\x56\x6c\xba\xeb\xf6\x97\xf7\xef\xee\x6e\xa3\x8d\xd7\xaf\x2f\x6e\xf2\x8b\x9b\x7c\x6e\x0c\xc1\x97\x0d\x93\xe8\xe6\x7b\x45\xf6\xc1\xbd\x4a\xf0\x73\x0d\xb5\x69\x2e\x62\x8b\x12\xae\xf3\x3c\xcb\xf3\x1c\x24\x96\x42\x56\xea\xf0\x9c\xf4\x87\x0f\x77\xef\x7e\x3c\xd1\x9e\x17\xd7\xb9\x79\x32\x58\xac\x26\x06\x9e\x10\xf4\x8d\x23\x92\xf4\xa2\xba\x93\xdc\x3a\x9b\x84\x41\xfa\x17\xd3\x3f\xe2\x59\xfd\x51\x62\xbb\x24\xc9\xff\xc6\xfc\x70\xe9\x70\x74\xa6\x44\x69\x33\xd5\x78\xaa\xe3\x4d\x35\x46\x5d\x30\xba\x2a\x2c\x16\x35\xe9\x48\x8d\xea\xa4\x0f\xd3\x41\x14\x2a\x67\xc2\x39\x35\xba\xec\xe9\xab\x63\xfe\x4c\xdb\xe8\x29\xed\xa7\x1b\xaf\xd7\x1e\x15\xd3\x68\x0c\xe7\x81\x70\x6a\x82\x9e\xea\xde\x47\x3a\x30\x9c\xe8\xd2\xee\x8f\x40\xb8\xcb\x46\xe5\xf9\x5f\x29\xc3\x49\xaa\xf3\x2a\xf9\x33\x00\x00\xff\xff\x6a\x25\xe9\x63\xba\x0d\x00\x00")

func scriptsStatLuaBytes() ([]byte, error) {
	return bindataRead(
		_scriptsStatLua,
		"scripts/stat.lua",
	)
}

func scriptsStatLua() (*asset, error) {
	bytes, err := scriptsStatLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/stat.lua", size: 3514, mode: os.FileMode(493), modTime: time.Unix(1494334185, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _textLogoTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x09\xc0\x50\x08\x05\xfb\x4c\x71\xa5\x16\x89\x0b\x09\xb7\xff\x16\xe1\x4b\x20\x9d\x16\xef\x44\x38\x7c\xac\xa3\xfe\xf1\xe1\x5a\x8d\x42\x33\x90\x0c\x10\x72\x0e\xab\x13\x84\xf7\x53\x48\x9f\x35\xc9\x56\x7a\xff\xd3\x6a\x4d\xc4\x54\xcb\x83\x7a\x03\x00\x00\xff\xff\x68\x0a\xe9\xbb\xce\x00\x00\x00")

func textLogoTxtBytes() ([]byte, error) {
	return bindataRead(
		_textLogoTxt,
		"text/logo.txt",
	)
}

func textLogoTxt() (*asset, error) {
	bytes, err := textLogoTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "text/logo.txt", size: 206, mode: os.FileMode(420), modTime: time.Unix(1494767085, 0)}
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
	"scripts/stat.lua": scriptsStatLua,
	"text/logo.txt": textLogoTxt,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"stat.lua": &bintree{scriptsStatLua, map[string]*bintree{}},
	}},
	"text": &bintree{nil, map[string]*bintree{
		"logo.txt": &bintree{textLogoTxt, map[string]*bintree{}},
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

