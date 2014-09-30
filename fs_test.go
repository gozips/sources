package sources

import "io"
import "testing"
import "github.com/nowk/assert"

func TestFSPathError(t *testing.T) {
	errmsg := "open path/to/doesnotexist.txt: no such file or directory"
	name, v, err := FS("path/to/doesnotexist.txt")
	assert.Equal(t, "doesnotexist.txt.txt", name)
	assert.Equal(t, errmsg, err.Error())

	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	assert.Equal(t, errmsg, string(b[:n]))

	r.Close()
}
