// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := NewEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	for _, entry := range ce {
		data, form, err := entry.readClass(className)
		if err == nil {
			return data, form, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (ce CompositeEntry) String() string {
	stirs := make([]string, len(ce))
	for i, entry := range ce {
		stirs[i] = entry.String()
	}
	return strings.Join(stirs, pathListSeparator)
}
