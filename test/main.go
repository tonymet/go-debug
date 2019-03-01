package main

import (
	. "github.com/tonymet/go-debug"
)

func main() {
	Debug("heythere")
	Debugf("dude! your name is %s", "bob")
}
