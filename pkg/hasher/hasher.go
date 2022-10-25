package hasher

import (
	"fmt"
	"io"
	"net/http"
)

type Hasher struct {
	Client httpClient
	Hash   hash
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type hash interface {
	Sum(b []byte) []byte
	Size() int
}

func (h *Hasher) Process(url string) (*Result, error) {
	resp, err := h.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &Result{
		URL:  url,
		Sum:  h.Hash.Sum(respBytes),
		Size: h.Hash.Size(),
	}, nil
}

type Result struct {
	URL  string
	Sum  []byte
	Size int
}

func (res Result) String() string {
	return fmt.Sprintf("%s %x", res.URL, res.Sum[:res.Size])
}
