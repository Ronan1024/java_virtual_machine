// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

import (
	"os"
	"strings"
)

// Define constants storing path separator
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// readClass Find and load class file
	// className: The relative path of the class file to be loaded (for example: java/lang/Object.class)
	// return: The byte array of the read class file, the Entry used to read the class file, and the error message
	readClass(className string) ([]byte, Entry, error)

	// String Similar to the toString method in Java
	// return: The absolute path of the Entry
	String() string
}

// NewEntry
// return: Entry
func NewEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
