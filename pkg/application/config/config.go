package config

import (
	"fmt"
	"net/url"
)

type Config struct {
	Parallel uint
	URLs     []url.URL
}

func New(parsers ...ParseFunc) (*Config, error) {
	out := Default()

	for _, p := range parsers {
		cfg, err := p(out)
		if err != nil {
			return nil, fmt.Errorf("failed to parse config: %w", err)
		}

		out = *cfg
	}

	return &out, nil
}

type ParseFunc func(Config) (*Config, error)
