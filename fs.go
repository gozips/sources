package source

import "os"
import "github.com/gozips/filepath"

// FS returns ReadCloser for an fs source
func FS(pathStr string) (string, interface{}) {
	name := filepath.Base(pathStr)

	r, err := os.Open(pathStr)
	if err != nil {
		return Errorize(name, err)
	}

	return name, r
}
