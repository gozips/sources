package sources

import "io"
import "os"
import "github.com/gozips/filepath"

// FS returns ReadCloser for an fs source
func FS(pathStr string) (string, io.ReadCloser, error) {
	name := filepath.Base(pathStr)

	r, err := os.Open(pathStr)
	if err != nil {
		return Errorize(name, err)
	}

	return name, r, nil
}
