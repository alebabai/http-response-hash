package hasher

import (
	"fmt"
	"io"
	"net/http"
)

type Hasher struct {
	client httpClient
	hash   hash
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type hash interface {
	Sum(b []byte) []byte
	Size() int
}

func New(c httpClient, hsh hash) (*Hasher, error) {
	h := &Hasher{
		client: c,
		hash:   hsh,
	}

	if err := h.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate a hasher: %w", err)
	}

	return h, nil
}

func (h *Hasher) Process(url string) (*Result, error) {
	resp, err := h.client.Get(url)
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
		Sum:  h.hash.Sum(respBytes),
		Size: h.hash.Size(),
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
