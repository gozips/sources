package source

import "net/http"
import "net/url"
import "github.com/gozips/filepath"

// HTTP returns a ReadCloser from an http source
func HTTP(urlStr string) (string, interface{}) {
	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return Errorize(urlStr, err)
	}

	name := filepath.Base(u)
	resp, err := http.DefaultClient.Do(&http.Request{
		Method: "GET",
		URL:    u,
	})
	if err != nil {
		return Errorize(name, err)
	}

	return name, resp.Body
}
