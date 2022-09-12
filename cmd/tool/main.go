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

func genericSlice(in []url.URL) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}

	return out
}

func main() {
	cfg, err := config.ParseArgs()
	if err != nil {
		fatal(fmt.Errorf("failed to parse config: %w", err))
	}

	svc := hasher.NewService(
		http.DefaultClient,
		md5.New(),
	)

	action := func(in interface{}) string {
		u := in.(url.URL)
		out, err := svc.Process(u.String())
		if err != nil {
			fatal(fmt.Errorf("failed to process %s: %w", u.String(), err))
		}

		return out.String()
	}
	consumer := func(out interface{}) {
		fmt.Println(out)
	}
	wp := pool.NewWorkerPool(
		action,
		consumer,
		pool.WithSize(cfg.Parallel),
	)
	wp.Process(genericSlice(cfg.URLs)...)
}
