package cmd

import (
	"log/slog"
	"os"
	"time"
)

func InitLogger() {
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{

				Level: slog.LevelInfo,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key == slog.TimeKey {
						a.Value = slog.StringValue(time.Now().Format(time.RFC3339))
					}

					return a
				},
			}),
		),
	)
}

func Fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
