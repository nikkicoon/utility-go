package pkg

import (
	"fmt"
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
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

func TrackExecutionTime(pre time.Time, msg ...string) {
	if len(msg) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%s\texecution time of %s: %s\n", time.Now(), GetCurrentFuncName(true), time.Since(pre).String())
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "%s\texecution time of %s (%s): %s\n", time.Now(), GetCurrentFuncName(true), msg[0], time.Since(pre).String())
	}
}
