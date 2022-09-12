package config

import (
	"flag"
	"fmt"
	"net/url"
)

func ParseArgs() (*Config, error) {
	cfg := &Config{
		Parallel: 10,
		URLs:     make([]url.URL, 0),
	}

	flag.IntVar(&cfg.Parallel, "parallel", cfg.Parallel, "limit the number of parallel requests")

	if !flag.Parsed() {
		flag.Parse()
	}

	for i, arg := range flag.Args() {
		u, err := url.Parse(arg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse argument at position %d: %w", i, err)
		}

		if u.Scheme == "" {
			u.Scheme = "http"
		}

		cfg.URLs = append(cfg.URLs, *u)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}
