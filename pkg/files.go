package pkg

import (
	"errors"
	"log/slog"
	"os"
	"os/exec"
)

// CheckFile checks if the given file `s` exists.
func CheckFile(s string) bool {
	_, err := os.Stat(s)
	return !errors.Is(err, os.ErrNotExist)
}

// CheckExecutable checks if the given executable `s` exists
// on the PATH. It returns true and nil, else it returns
// false and an error.
func CheckExecutable(s string, logger *slog.Logger) (bool, error) {
	path, err := exec.LookPath(s)
	if err != nil {
		logger.Debug("", slog.Any("error", err))
	}
	logger.Debug("file found", slog.String("file", s), slog.String("path", path))
	cmd := exec.Command(s)
	if err := cmd.Run(); err != nil {
		return false, err
	}
	logger.Debug("file is executable", slog.String("file", s))
	return true, nil
}
