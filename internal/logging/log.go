package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Init establishes a global logger instance that writes
// to stdout for now.
func Init() {
	defaultLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Logger = defaultLogger
}
