package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/logrusorgru/aurora"
)

// Log : Log message to stdout with LOG level
func Log(format string, args ...interface{}) (int, error) {
	return fmt.Printf(("[" + time.Now().String() + "][LOG] " + format + "\n"), args...)
}

// LogError : Log error to stderr with ERROR level
func LogError(err error) (int, error) {
	return fmt.Fprintf(
		os.Stderr,
		aurora.Red("["+time.Now().Format("Mon Jan 2 15:04:05")+"][ERROR] %s\n").String(),
		err,
	)
}
