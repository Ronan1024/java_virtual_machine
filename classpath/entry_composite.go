// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

type CompositeEntry struct {
	absDir string
}

func newCompositeEntry(path string) *CompositeEntry {
	return nil
}

func (ce *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}
func (ce *CompositeEntry) String() string {
	return ""
}
