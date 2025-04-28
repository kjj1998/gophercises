package request

import (
	"bytes"
	"io"
	"net/http"
)

func ReadPageHtml(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return bytes.NewReader(body), nil
}
