// Package classpath
// @author: Ronan
// @since: 2023/8/4
// @desc:
package classpath

type WildcardEntry struct {
	absDir string
}

func newWildcardEntry(path string) *WildcardEntry {
	return nil
}
func (we *WildcardEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func (we *WildcardEntry) String() string {
	return ""
}
