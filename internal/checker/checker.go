package checker

import (
	"net/http"
)

type LinkResult struct {
	URL        string
	StatusCode int
	Alive      bool
	Error 	   string
}

func CheckLink(url string) LinkResult {
	resp, err := http.Head(url)
	if err != nil {
		return LinkResult {
			URL: url,
			StatusCode: 0,
			Alive: false,
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	aliveLink := resp.StatusCode >= 200 && resp.StatusCode < 400
	return LinkResult {
		URL: url,
		StatusCode: resp.StatusCode,
		Alive: aliveLink,
		Error: "",
	}
}