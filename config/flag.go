package config

import (
	"flag"
	"fmt"
)

func ParseArgs() (*Config, error) {
	cfg := &Config{
		Parallel: 10,
		Inputs:   make([]string, 0),
	}

	flag.IntVar(&cfg.Parallel, "parallel", cfg.Parallel, "limit the number of parallel requests")

	if !flag.Parsed() {
		flag.Parse()
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}
