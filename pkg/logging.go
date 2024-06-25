package pkg

import (
	"io"
	"log"
	"log/slog"
	"os"
)

// InitializeSlogLogger initializes a new `log/slog` logger.
// level: loglevel (debug, info, warning, error)
// logfile: either stdout, stderr or file
// format: json or text, every other argument defaults to text
func InitializeSlogLogger(level string, logfile string, format string) *slog.Logger {
	var logFile io.Writer
	switch logfile {
	case "stdout":
		logFile = os.Stdout
	case "stderr":
		logFile = os.Stderr
	case "file":
		f, err := os.OpenFile("/var/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		logFile = f
	}
	var l = new(slog.LevelVar)
	switch level {
	case "debug":
		l.Set(slog.LevelDebug)
	case "info":
		l.Set(slog.LevelInfo)
	case "warning":
		l.Set(slog.LevelWarn)
	case "error":
		l.Set(slog.LevelError)
	}
	handlerOptions := &slog.HandlerOptions{Level: l}
	switch format {
	case "json":
		return slog.New(slog.NewJSONHandler(logFile, handlerOptions))
	case "text":
		return slog.New(slog.NewTextHandler(logFile, handlerOptions))
	default:
		return slog.New(slog.NewTextHandler(logFile, handlerOptions))
	}
}
