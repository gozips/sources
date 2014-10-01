package sources

import "testing"
import "github.com/nowk/assert"

func TestFSPathError(t *testing.T) {
	errmsg := "open path/to/doesnotexist.txt: no such file or directory"
	name, r, err := FS("path/to/doesnotexist.txt")
	defer r.Close()
	assert.Equal(t, "doesnotexist.txt.txt", name)
	assert.Equal(t, errmsg, err.Error())

	b := make([]byte, 32*1024)
	n, _ := r.Read(b)
	assert.Equal(t, errmsg, string(b[:n]))
}
