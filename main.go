//+build debug

package debug

import (
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"runtime"

	"github.com/fatih/color"
	isatty "github.com/mattn/go-isatty"
)

var (
	extractLastPkgName   = regexp.MustCompile(`(?i)([a-z0-9_-]+)\.([a-z0-9_-]+)\.func1$`)
	debugPatternMatch, _ = regexp.Compile("^" + os.Getenv("DEBUG") + "$")
	isTty                = isatty.IsTerminal(os.Stdin.Fd())
	colorMap             = []color.Attribute{color.FgRed, color.FgYellow, color.FgGreen, color.FgHiMagenta}
)

func Debug(message interface{}) {
	if lastPkgName, ok := active(2); ok {
		if isTty {
			color := color.New(colorMap[hashToBucket(lastPkgName, uint32(len(colorMap)))])
			color.Printf("%s %s\n", lastPkgName, message)
		} else {
			fmt.Printf("%s %s", lastPkgName, message)
		}
	}
}

func hashToBucket(keyName string, bucketCount uint32) uint32 {
	hash := fnv.New32()
	hash.Write([]byte(keyName))
	return hash.Sum32() % bucketCount
}

func Debugf(format string, args ...interface{}) {
	if lastPkgName, ok := active(2); ok {
		if isTty {
			color := color.New(colorMap[hashToBucket(lastPkgName, uint32(len(colorMap)))])
			color.Printf(lastPkgName+" "+format+"\n", args...)
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
	if len(matches) == 3 {
		return matches[1]
	}
	return ""
}
