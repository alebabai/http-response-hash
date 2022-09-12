package hasher

import (
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	URL  string
	Sum  []byte
	Size int
}

func (res Result) String() string {
	return fmt.Sprintf("%s %x", res.URL, res.Sum[:res.Size])
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type hashSum interface {
	Sum(b []byte) []byte
	Size() int
}

type Hasher struct {
	client httpClient
	hash   hashSum
}

func New(client httpClient, hash hashSum) *Hasher {
	return &Hasher{
		client: client,
		hash:   hash,
	}
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
