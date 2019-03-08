//+build !debug

package debug

// Debug print message when debugging
func Debug(message interface{}) {}

// Debugf printf version of Debug
func Debugf(format string, args ...interface{}) {}

func lastPkgName(s string) string                            { return "" }
func active(i int) (string, bool)                            { return "", false }
func hashToBucket(keyName string, bucketCount uint32) uint32 { return 0 }
