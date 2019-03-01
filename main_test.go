package debug

import (
	"io/ioutil"
	"os"
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
		assert.Equal(t, "hi", string(out))
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

func TestDebugf(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Setenv("DEBUG", "go-debug")
	t.Run("Test debug.Debugf ", func(t *testing.T) {
		Debugf("hey %s", "tony")
		w.Close()
		out, _ := ioutil.ReadAll(r)
		assert.Equal(t, "hey tony", string(out))
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
