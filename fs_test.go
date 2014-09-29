package source

import "io"
import "testing"
import "github.com/nowk/assert"

func TestFSPathError(t *testing.T) {
	name, v := FS("path/to/doesnotexist.txt")
	assert.Equal(t, "doesnotexist.txt.txt", name)

	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	assert.Equal(t, "open path/to/doesnotexist.txt: no such file or directory", string(b[:n]))

	r.Close()
}
