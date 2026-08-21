package config

import (
	"io"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Host string
	Port int

	ShutdownTime time.Duration

	LoggerWriter io.Writer
	LoggerOpts   *slog.HandlerOptions
}

func Default() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,

		ShutdownTime: 5 * time.Second,

		LoggerWriter: os.Stdout,
		LoggerOpts: &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		},
	}
}
