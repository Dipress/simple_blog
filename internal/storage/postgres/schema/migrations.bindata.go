// Code generated by go-bindata.
// sources:
// migrations/1549100465_posts.down.sql
// migrations/1549100465_posts.up.sql
// migrations/1557063976_users.down.sql
// migrations/1557063976_users.up.sql
// DO NOT EDIT!

package schema

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

var __1549100465_postsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xc8\x2f\x2e\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x95\xd2\xda\x4b\x1c\x00\x00\x00")

func _1549100465_postsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1549100465_postsDownSql,
		"1549100465_posts.down.sql",
	)
}

func _1549100465_postsDownSql() (*asset, error) {
	bytes, err := _1549100465_postsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1549100465_posts.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1549101120, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1549100465_postsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8e\x31\x4b\xc4\x40\x10\x46\xeb\xd9\x5f\x31\xe5\xdd\x71\x70\x20\x5c\x65\x35\xc6\x39\x5c\xdc\xac\xc7\xec\x44\x4c\x15\xa2\xbb\xc5\x82\x21\x21\xbb\x16\xfe\x7b\xd1\x22\x5a\xdb\x7e\xef\x3d\xf8\x1a\x61\x52\x46\xa5\x3b\xc7\x68\x2f\xe8\x9f\x14\xf9\xc5\x06\x0d\xb8\xcc\xa5\x16\xdc\x19\xc8\x11\x02\x8b\x25\x87\x57\xb1\x2d\x49\x8f\x8f\xdc\x1f\x0d\x7c\x94\xb4\x0e\x39\xa2\xf5\xfa\xd3\xf9\xce\xb9\xa3\x81\x9a\xeb\x7b\x82\x67\x92\xe6\x81\x04\x77\x37\xe7\xf3\xfe\x2f\x7e\x9d\xe3\xe7\x46\x7f\x77\x03\xa7\x43\xcd\x53\x2a\x75\x9c\x96\xc3\xc9\xc0\xdb\x9a\xc6\x9a\xe2\x30\x56\x50\xdb\x72\x50\x6a\xaf\x9b\x8f\xf7\x7c\xa1\xce\x29\x36\x9d\x08\x7b\x1d\x36\xe5\xfb\xd7\x12\xff\x53\x9a\xfd\xad\xf9\x0a\x00\x00\xff\xff\x38\x7a\x61\x2c\x0f\x01\x00\x00")

func _1549100465_postsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1549100465_postsUpSql,
		"1549100465_posts.up.sql",
	)
}

func _1549100465_postsUpSql() (*asset, error) {
	bytes, err := _1549100465_postsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1549100465_posts.up.sql", size: 271, mode: os.FileMode(420), modTime: time.Unix(1558630843, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1557063976_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x2c\x02\x3d\xa7\x1c\x00\x00\x00")

func _1557063976_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1557063976_usersDownSql,
		"1557063976_users.down.sql",
	)
}

func _1557063976_usersDownSql() (*asset, error) {
	bytes, err := _1557063976_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1557063976_users.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1557065056, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1557063976_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x4b\xc3\x30\x14\x87\xcf\xc9\x5f\xf1\x8e\xcb\x18\x4c\x04\xf1\xd0\x53\xec\xde\x30\x98\x66\x33\x7d\x91\xed\x14\x82\x09\xac\x60\xb5\x34\x15\xfd\xf3\x65\xc5\x4e\x11\x61\xb0\xeb\xe3\xfb\x7e\x3c\xbe\xd2\xa2\x24\x04\x92\x77\x1a\x41\xad\xc1\x6c\x08\x70\xa7\x6a\xaa\xe1\x3d\xa7\x3e\xc3\x8c\xb3\x26\xb2\x1a\xad\x92\x1a\xb6\x56\x55\xd2\xee\xe1\x01\xf7\x0b\xce\x8e\xc0\x6b\x68\x13\x3c\x49\x5b\xde\x4b\x0b\xb3\x9b\x2b\x01\xce\xa8\x47\x87\xe3\x90\x71\x5a\x2f\x38\x4b\x6d\x68\x5e\xd8\x19\xa8\x0b\x39\x7f\xbc\xf5\xd1\x1f\x42\x3e\xfc\xc0\xb7\xd7\xe2\x17\xc5\xd9\x72\x0e\x43\xd3\xa6\x3c\x84\xb6\x83\xf9\x92\xb3\xe7\x3e\x85\x21\x45\x1f\x06\x46\xaa\xc2\x9a\x64\xb5\x3d\x19\xb0\xc2\xb5\x74\x9a\xa0\x74\xd6\xa2\x21\x7f\x42\x8e\xef\x77\xf1\x12\x93\x8b\x82\xf3\xef\x6c\xca\xac\x70\xf7\x5f\x36\x3f\xb5\xf1\x4d\xfc\x84\x8d\x99\x62\x4e\x67\x51\x9c\x9f\x18\xb3\xfd\xf1\xc7\x9b\x28\xbe\x02\x00\x00\xff\xff\xbd\x97\x39\x8f\xb6\x01\x00\x00")

func _1557063976_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1557063976_usersUpSql,
		"1557063976_users.up.sql",
	)
}

func _1557063976_usersUpSql() (*asset, error) {
	bytes, err := _1557063976_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1557063976_users.up.sql", size: 438, mode: os.FileMode(420), modTime: time.Unix(1558164651, 0)}
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
	"1549100465_posts.down.sql": _1549100465_postsDownSql,
	"1549100465_posts.up.sql": _1549100465_postsUpSql,
	"1557063976_users.down.sql": _1557063976_usersDownSql,
	"1557063976_users.up.sql": _1557063976_usersUpSql,
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
	"1549100465_posts.down.sql": &bintree{_1549100465_postsDownSql, map[string]*bintree{}},
	"1549100465_posts.up.sql": &bintree{_1549100465_postsUpSql, map[string]*bintree{}},
	"1557063976_users.down.sql": &bintree{_1557063976_usersDownSql, map[string]*bintree{}},
	"1557063976_users.up.sql": &bintree{_1557063976_usersUpSql, map[string]*bintree{}},
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

