package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/alebabai/http-response-hash/pkg/config"
	"github.com/alebabai/http-response-hash/pkg/hasher"
	"github.com/alebabai/http-response-hash/pkg/pool"
)

func fatal(err error) {
	fmt.Println(fmt.Errorf("fatal: %w", err))
	os.Exit(1)
}

func main() {
	cfg, err := config.ParseArgs()
	if err != nil {
		fatal(fmt.Errorf("failed to parse config: %w", err))
	}

	h, err := hasher.New(
		http.DefaultClient,
		md5.New(),
	)
	if err != nil {
		fatal(fmt.Errorf("failed to init hasher: %w", err))
	}

	action := func(u url.URL) string {
		res, err := h.Process(u.String())
		if err != nil {
			fatal(fmt.Errorf("failed to process %s: %w", u.String(), err))
		}

		return res.String()
	}
	consumer := func(in string) {
		fmt.Println(in)
	}
	p, err := pool.New(
		action,
		consumer,
		cfg.Parallel,
	)
	if err != nil {
		fatal(fmt.Errorf("failed to init pool: %w", err))
	}

	p.Process(cfg.URLs...)
}
