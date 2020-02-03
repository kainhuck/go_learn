package requests

import (
	"io/ioutil"
	"net/http"
)

// Requests ...
type Requests struct {
	StatusCode int
	Text       string
}

// Get ...
func Get(url string, headers map[string]string) (r *Requests, err error) {
	client := &http.Client{}
	r = new(Requests)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	// Add headers
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r.StatusCode = resp.StatusCode

	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r.Text = string(pageBytes)

	return
}
