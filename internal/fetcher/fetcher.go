package fetcher

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func FetchHTML(url string) (io.Reader, error){
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error - %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)

	return reader, nil
}