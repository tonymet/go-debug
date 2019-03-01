package debug

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/go-stack/stack"
)

// set namespace
func Name(name string) {

}

func Debug(message interface{}) {
	if active(2) {
		fmt.Print(message)
		//color.Red("%s", message)
	}
}

func Debugf(format string, args ...interface{}) {
	if active(2) {
		fmt.Printf(format, args...)
		//color.Red(format, args...)
	}
}

// see if active based on calling method
func active(level int) bool {
	// if calling methodg
	c := stack.Caller(level)
	log.Print("FUNCTION: ", c.Frame().Function)
	theLastPkgName := lastPkgName(c.Frame().Function)
	log.Print("lastPkgName: ", theLastPkgName)
	return os.Getenv("DEBUG") == theLastPkgName
}

func lastPkgName(fullPkgName string) string {
	var re = regexp.MustCompile(`(?msi)(?:.*\/)?([a-z_-]+)\..*$`)
	matches := re.FindStringSubmatch(fullPkgName)
	log.Printf("%v", matches)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
