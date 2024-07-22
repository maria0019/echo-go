package app

import (
	slogecho "github.com/samber/slog-echo"
	slogformatter "github.com/samber/slog-formatter"
	"log/slog"
	"os"
	"time"
)

func InitLogger() *slog.Logger {
	return slog.New(
		slogformatter.NewFormatterHandler(
			slogformatter.TimezoneConverter(time.UTC),
			slogformatter.TimeFormatter(time.DateTime, nil),
		)(
			slog.NewTextHandler(os.Stdout, nil),
		),
	)
}

func InitLoggerConfig() slogecho.Config {
	return slogecho.Config{
		WithSpanID:         true,
		WithTraceID:        true,
		DefaultLevel:       slog.LevelInfo,
		ClientErrorLevel:   slog.LevelWarn,
		ServerErrorLevel:   slog.LevelError,
		WithRequestBody:    true,
		WithResponseBody:   true,
		WithRequestHeader:  true,
		WithResponseHeader: true,
	}
}
