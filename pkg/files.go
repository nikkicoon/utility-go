package pkg

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/phuslu/log"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"time"
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
		switch logger {
		case nil:
			_, _ = fmt.Fprint(os.Stderr, "%w\n", err)
		default:
			logger.Debug("", slog.Any("error", err))
		}
	}
	switch logger {
	case nil:
		_, _ = fmt.Fprintf(os.Stderr, "file found: %q at path %q\n", s, path)
	default:
		logger.Debug("file found", slog.String("file", s), slog.String("path", path))
	}
	cmd := exec.Command(s)
	if err := cmd.Run(); err != nil {
		return false, err
	}
	switch logger {
	case nil:
		_, _ = fmt.Fprintf(os.Stderr, "file is executable: %q\n", s)
	default:
		logger.Debug("file is executable", slog.String("file", s))
	}
	return true, nil
}

// LineCounter counts the number of lines in a file in a sufficiently fast way.
func LineCounter(logger *log.Logger, r io.Reader, tFlag ...bool) (int, error) {
	start := time.Now()
	if len(tFlag) > 0 {
		switch logger {
		case nil:
			defer TrackExecutionTime(nil, start)
		default:
			defer TrackExecutionTime(logger, start)
		}
	}
	var c int
	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			c++
		}
		if err == io.EOF {
			break
		}
	}
	return c, nil
}
