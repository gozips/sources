package sources

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"
import "github.com/gozips/filepath"

// HTTP returns a ReadCloser from an http source
func HTTP(urlStr string) (string, interface{}) {
	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return fmt.Sprintf("%s.txt", urlStr), errReadCloser(err)
	}

	name := filepath.Base(u)
	req := &http.Request{Method: "GET", URL: u}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("%s.txt", name), errReadCloser(err)
	}

	return name, resp.Body
}

func errReadCloser(err error) (r io.ReadCloser) {
	if err != nil {
		r = ioutil.NopCloser(strings.NewReader(err.Error()))
	}

	return
}
