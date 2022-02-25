package config

import (
	"flag"
	"fmt"
)

func ParseArgs() (*Config, error) {
	cfg := &Config{
		Parallel: 10,
	}

	flag.IntVar(&cfg.Parallel, "parallel", cfg.Parallel, "limit the number of parallel requests")

	if !flag.Parsed() {
		flag.Parse()
	}

	cfg.Inputs = flag.Args()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}