package sources

import "fmt"
import "net/http"
import "net/url"
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
