package pkg

import (
	"fmt"
	"runtime"
)

// GetCurrentFuncName returns a string holding the current
// function name.
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
