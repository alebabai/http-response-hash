package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func ParseArgs(out Config) (*Config, error) {
	fs := flag.NewFlagSet("application", flag.ContinueOnError)

	fs.UintVar(&out.Parallel, "parallel", out.Parallel, "limit the number of parallel requests")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse command line arguments: %w", err)
	}

	for i, arg := range fs.Args() {
		u, err := url.Parse(arg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse argument at position %d: %w", i, err)
		}

		if u.Scheme == "" {
			u.Scheme = "http"
		}

		out.URLs = append(out.URLs, *u)
	}

	return &out, nil
}
