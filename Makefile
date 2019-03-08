.PHONY:test
test:
	DEBUG='.*' go test -tags debug .
bench:
	DEBUG='.*' go test -cpuprofile cpu.prof -benchmem -tags debug -run=^$ github.com/tonymet/go-debug -bench '^(BenchmarkActive)'
