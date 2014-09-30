package sources

import "io"
import "regexp"
import "testing"
import "github.com/nowk/assert"

func TestHTTPURLParseError(t *testing.T) {
	errmsg := "parse thisisabadurl: invalid URI for request"

	name, v, err := HTTP("thisisabadurl")
	assert.Equal(t, "thisisabadurl.txt", name)
	assert.Equal(t, errmsg, err.Error())

	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	assert.Equal(t, errmsg, string(b[:n]))

	r.Close()
}

func TestHTTPClientError(t *testing.T) {
	reg := regexp.MustCompile(`Get http:\/\/unreachable:( dial tcp:)? lookup unreachable: no such host`)

	// fails if ISP picks up and redirects to search, which TWC does
	name, v, err := HTTP("http://unreachable")
	assert.Equal(t, "unreachable.txt", name)
	if !reg.MatchString(err.Error()) {
		t.Errorf("Expected %s to match %s", err.Error(), reg.String())
	}

	b := make([]byte, 32*1024)
	r := v.(io.ReadCloser)
	n, _ := r.Read(b)
	if str := string(b[:n]); !reg.MatchString(str) {
		t.Errorf("Expected %s to match %s", str, reg.String())
	}

	r.Close()
}

// func TestHTTPAppendsExtFromContentType(t *testing.T) {
// 	t.Skip("")
// }
