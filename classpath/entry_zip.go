// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

// newZipEntry create a new ZipEntry
// path: The absolute path of the zip file
// return: ZipEntry
func newZipEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs}
}
func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, ze, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// String Similar to the toString method in Java
// return: The absolute path of the Entry
func (ze *ZipEntry) String() string {
	return ze.absPath
}
