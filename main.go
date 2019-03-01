package debug

import (
	"fmt"
)

// set namespace
func Name(name string) {

}

func Debug(message interface{}) {
	fmt.Print(message)
}

func Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)

}
