package sources

import "fmt"
import "os"
import "github.com/gozips/filepath"

// FS returns ReadCloser for an fs source
func FS(pathStr string) (string, interface{}) {
	name := filepath.Base(pathStr)
	r, err := os.Open(pathStr)
	if err != nil {
		name = fmt.Sprintf("%s.txt", name) // force to .txt
		return name, errReadCloser(err)
	}

	return name, r
}
