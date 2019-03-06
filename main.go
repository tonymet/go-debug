// +build debug

package debug

import (
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/fatih/color"
	isatty "github.com/mattn/go-isatty"
)

var (
	extractLastPkgName = regexp.MustCompile(`(?msi)(?:.*\/)?([a-z_-]+)\..*$`)
	isTty              = isatty.IsTerminal(os.Stdin.Fd())
)

func Debug(message interface{}) {
	if lastPkgName, ok := active(2); ok {
		if isTty {
			color.Red("%s %s", lastPkgName, message)
		} else {
			fmt.Printf("%s %s", lastPkgName, message)
		}
	}
}

func Debugf(format string, args ...interface{}) {
	if lastPkgName, ok := active(2); ok {
		if isTty {
			color.Red(lastPkgName+" "+format, args...)
		} else {
			fmt.Printf(lastPkgName+" "+format, args...)
		}
	}
}

// see if active based on calling method
func active(level int) (string, bool) {
	// if calling methodg
	//c := stack.Caller(level)
	pc, _, _, ok := runtime.Caller(level)
	if !ok {
		return "", false
	}
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	//log.Printf("function: %s", function)
	theLastPkgName := lastPkgName(frame.Function)
	//log.Printf("lastPkgName: %s", function)
	return theLastPkgName, os.Getenv("DEBUG") == theLastPkgName
}

func lastPkgName(fullPkgName string) string {
	matches := extractLastPkgName.FindStringSubmatch(fullPkgName)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
