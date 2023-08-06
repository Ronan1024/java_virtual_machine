// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

import (
	"archive/zip"
	"errors"
	"io"
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
	reader, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == className {
			readClose, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer readClose.Close()
			data, err := io.ReadAll(readClose)
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
