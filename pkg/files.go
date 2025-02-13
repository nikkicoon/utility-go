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
	"path/filepath"
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

func SymlinkFiles(logger *log.Logger, origin string, target string) error {
	origin, err := filepath.Abs(origin)
	if err != nil {
		return err
	}
	target, err = filepath.Abs(target)
	if err != nil {
		return err
	}
	if _, err = os.Lstat(target); err == nil {
		if err = os.Remove(target); err != nil {
			return fmt.Errorf("failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
		// warning, the file could just not exist because it was not yet created
		switch logger {
		case nil:
			_, _ = fmt.Fprintf(os.Stderr, "WARNING: failed to check symlink: %+v\n", err)
		default:
			logger.Warn().Msgf("failed to check symlink: %+v", err)
		}
	}
	if err = os.Symlink(origin, target); err != nil {
		return fmt.Errorf("failed to symlink: %+v", err)
	} else {
		// no error, but is it a broken symlink?
		_, err = os.Stat(target)
		if os.IsNotExist(err) {
			switch logger {
			case nil:
				_, _ = fmt.Fprintf(os.Stderr, "broken symlink: %+v\n", err)
			default:
				logger.Info().Msgf("broken symlink: %+v", err)
			}
			if err = os.Remove(target); err != nil {
				return fmt.Errorf("failed to unlink: %+v", err)
			}
			return fmt.Errorf("broken symlink: %+v", err)
		} else {
			switch logger {
			case nil:
				_, _ = fmt.Fprintf(os.Stderr, "symlinked file %s to %s\n", origin, target)
			default:
				logger.Info().Msgf("symlinked file %s to %s", origin, target)
			}
		}
	}
	return nil
}
