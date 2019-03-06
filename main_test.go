package debug

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Setenv("DEBUG", "go-debug")
	t.Run("Test debug.Debug ", func(t *testing.T) {
		Debug("hi")
		w.Close()
		out, _ := ioutil.ReadAll(r)
		assert.Equal(t, "go-debug hi", string(out))
	})
	os.Stdout = old
}

func TestLastPkgName(t *testing.T) {
	assert.Equal(t, lastPkgName("_/Users/tonym/sotion/go-debug.TestReflect"), "go-debug")
}

func TestActive(t *testing.T) {
	os.Setenv("DEBUG", "go-debug")
	_, ok := active(1)
	assert.Equal(t, true, ok)
}

func TestReflect(t *testing.T) {

	pc, myFunction, _, _ := runtime.Caller(1)
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	myType := reflect.TypeOf(frame.Func)
	fmt.Printf("typeof myfunc: %v", myType)
	fmt.Printf("typeof myfunc: %v", myType.PkgPath())
	fmt.Printf("myFunction: %v", myFunction)

}

func TestDebugf(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Setenv("DEBUG", "go-debug")
	t.Run("Test debug.Debugf ", func(t *testing.T) {
		Debugf("hey %s", "tony")
		w.Close()
		out, _ := ioutil.ReadAll(r)
		assert.Equal(t, "go-debug hey tony", string(out))
	})
	os.Stdout = old
}

// from fib_test.go
func BenchmarkDebug(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Debug("hi bench")
	}
}

func BenchmarkDebugf(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Debugf("hi %s", "bob")
	}
}

func BenchmarkActive(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		active(1)
	}
}
