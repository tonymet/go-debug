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
	extractLastPkgName   = regexp.MustCompile(`(?i)([a-z_-]+)\..*$`)
	debugPatternMatch, _ = regexp.Compile("^" + os.Getenv("DEBUG") + "$")
	isTty                = isatty.IsTerminal(os.Stdin.Fd())
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
	if debugPatternMatch == nil {
		return "", false
	}
	pc, _, _, ok := runtime.Caller(level)
	if !ok {
		return "", false
	}
	theLastPkgName := lastPkgName(runtime.FuncForPC(pc).Name())
	return theLastPkgName, debugPatternMatch.MatchString(theLastPkgName)
}

func lastPkgName(fullPkgName string) string {
	matches := extractLastPkgName.FindStringSubmatch(fullPkgName)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
