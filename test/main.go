package main

import (
	. "github.com/tonymet/go-debug"
	. "github.com/tonymet/go-debug/test/lib"
)

func main() {
	Debug("heythere")
	Debugf("dude! your name is %s", "bob")
	DoStuff()
}
