package source

import "io"
import "regexp"
import "testing"
import "github.com/nowk/assert"

func TestHTTPURLParseError(t *testing.T) {
	name, v := HTTP("thisisabadurl")
	assert.Equal(t, "thisisabadurl.txt", name)

	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	assert.Equal(t, "parse thisisabadurl: invalid URI for request", string(b[:n]))

	r.Close()
}

func TestHTTPClientError(t *testing.T) {
	// fails if ISP picks up and redirects to search, which TWC does
	name, v := HTTP("http://unreachable")
	assert.Equal(t, "unreachable.txt", name)

	reg := regexp.MustCompile(`Get http:\/\/unreachable:( dial tcp:)? lookup unreachable: no such host`)
	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	if str := string(b[:n]); !reg.MatchString(str) {
		t.Errorf("Expected %s, got %s", reg.String(), str)
	}

	r.Close()
}

// func TestHTTPAppendsExtFromContentType(t *testing.T) {
// 	t.Skip("")
// }
