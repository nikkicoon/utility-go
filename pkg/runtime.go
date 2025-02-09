package pkg

import (
	"fmt"
	"github.com/phuslu/log"
	"os"
	"runtime"
	"time"
)

// GetCurrentFuncName returns a string holding the current
// function name. Pass any boolean value to it to get the
// name of the function which called the function.
func GetCurrentFuncName(parent ...bool) string {
	var pc uintptr
	if len(parent) == 1 {
		pc, _, _, _ = runtime.Caller(2)
	} else {
		pc, _, _, _ = runtime.Caller(1)
	}
	return runtime.FuncForPC(pc).Name()
}

// TrackExecutionTime takes a [time.Time] value and an optional message, and returns
// a log message (or [fmt.Fprintf] if logger was nil) with details on the time spent
// executing a function.
func TrackExecutionTime(logger *log.Logger, pre time.Time, msg ...string) {
	switch logger {
	case nil:
		switch len(msg) {
		case 0:
			_, _ = fmt.Fprintf(os.Stderr, "%s\texecution time of %s: %s\n", time.Now(), GetCurrentFuncName(true), time.Since(pre).String())
		default:
			_, _ = fmt.Fprintf(os.Stderr, "%s\texecution time of %s (%s): %s\n", time.Now(), GetCurrentFuncName(true), msg[0], time.Since(pre).String())
		}
	default:
		switch len(msg) {
		case 0:
			logger.Trace().Str("function name", GetCurrentFuncName(true)).Str("execution time", time.Since(pre).String()).Msg("")
		default:
			logger.Trace().Str("function name", GetCurrentFuncName(true)).Str("execution time", time.Since(pre).String()).Str("context message", msg[0]).Msg("")
		}
	}
}

func TrackTimeSeconds(pre time.Time, fn func(float64)) {
	fn(time.Since(pre).Seconds())
}