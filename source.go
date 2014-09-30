package sources

import "fmt"
import "io"
import "io/ioutil"
import "strings"

// Errorize is a helper to append txt to file names and return the error message
// as a ReadCloser to be written in as an entry
func Errorize(name string, e error) (string, io.ReadCloser, error) {
	if e != nil {
		r := strings.NewReader(e.Error())
		return fmt.Sprintf("%s.txt", name), ioutil.NopCloser(r), e
	}

	return "", nil, nil
}
