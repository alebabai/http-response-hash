package hasher

import (
	"fmt"
	"io"
	"net/http"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type hashSum interface {
	Sum(b []byte) []byte
	Size() int
}

type Service struct {
	client httpClient
	hash   hashSum
}

func NewService(client httpClient, hash hashSum) *Service {
	return &Service{
		client: client,
		hash:   hash,
	}
}

func (svc *Service) Process(url string) (*Output, error) {
	resp, err := svc.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	out := &Output{
		URL:  url,
		Sum:  svc.hash.Sum(respBytes),
		Size: svc.hash.Size(),
	}

	return out, nil
}
