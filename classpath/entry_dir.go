// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

// newDirEntry create a new DirEntry
// path: The absolute path of the directory
// return: DirEntry
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// readClass read the class file from the directory
// className: The relative path of the class file to be loaded (for example: java/lang/Object.class)
// return: The byte array of the read class file, the Entry used to read the class file, and the error message
func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, de, err
}

// String
// return: The absolute path of the Entry
func (de *DirEntry) String() string {
	return de.absDir
}
