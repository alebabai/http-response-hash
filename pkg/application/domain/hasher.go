package domain

import (
	"context"
	"fmt"
	"hash"
	"io"
	"net/http"

	"github.com/alebabai/http-response-hash/pkg/application/domain/hasher"
)

type HasherService struct {
	client HTTPClient
	hash   hash.Hash
}

func NewHasherService(h hash.Hash, opts ...HasherServiceOption) *HasherService {
	svc := &HasherService{
		client: http.DefaultClient,
		hash:   h,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

type HasherServiceOption func(*HasherService)

func HasherServiceWithClient(c *http.Client) HasherServiceOption {
	return func(svc *HasherService) {
		svc.client = c
	}
}

func (svc *HasherService) HashURLContent(ctx context.Context, in hasher.HashURLContentInput) (*hasher.HashURLContentOutput, error) {
	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate input: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, in.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize http request: %w", err)
	}

	resp, err := svc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &hasher.HashURLContentOutput{
		URL:  in.URL,
		Sum:  svc.hash.Sum(data),
		Size: svc.hash.Size(),
	}, nil
}
