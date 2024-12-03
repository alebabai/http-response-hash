package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/alebabai/http-response-hash/cmd"
	"github.com/alebabai/http-response-hash/pkg/application/config"
	"github.com/alebabai/http-response-hash/pkg/application/domain"
	"github.com/alebabai/http-response-hash/pkg/application/domain/hasher"
)

func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	{
		ctx = context.Background()
		ctx, cancel = context.WithCancel(ctx)
		defer func() {
			cancel()
			time.Sleep(3 * time.Second)
		}()
	}

	cmd.InitLogger()

	slog.Info("initialization of the application")

	slog.Info("initializing configuration")

	cfg, err := config.New(config.ParseArgs)
	if err != nil {
		cmd.Fatal(fmt.Errorf("failed to parse config: %w", err))
	}

	slog.Info("initializing services")

	svc := domain.NewHasherService(
		md5.New(),
	)

	wp := &cmd.WorkerPool[url.URL, hasher.HashURLContentOutput]{
		Action: func(u url.URL) hasher.HashURLContentOutput {
			res, err := svc.HashURLContent(ctx, hasher.HashURLContentInput{
				URL: u.String(),
			})
			if err != nil {
				cmd.Fatal(fmt.Errorf("failed to process %s: %w", u.String(), err))
			}

			return *res
		},
		Consumer: func(in hasher.HashURLContentOutput) {
			slog.Info("hashing results", slog.String("input_url", in.URL), slog.String("content_hash", fmt.Sprintf("%x", in.Sum[:in.Size])))
		},
		Size: cfg.Parallel,
	}

	slog.Info("processing urls")

	wp.Process(cfg.URLs...)
}
