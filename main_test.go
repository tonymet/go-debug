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
	assert.Equal(t, true, active(1))
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
