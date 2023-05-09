// Actual trace_test.go is in ../report directory.
package trace

import (
	"strings"
	"testing"
)

func TestTraceInfo(t *testing.T) {
	trace := strings.Join(Info(func(filepath string) bool { return false }), "\n\t")
	if !strings.Contains(trace, "trace/trace_test.go:10") {
		t.Errorf("Wrong trace: %#v", trace)
	}
}

func TestIsGolangTestFunc(t *testing.T) {
	if !isGolangTestFunc("TestFoo") {
		t.Errorf("Expect that TestFoo is test func")
	}
	if !isGolangTestFunc("Benchmark") {
		t.Errorf("Expect that Benchmark is test func")
	}
	if isGolangTestFunc("JustFoo") {
		t.Errorf("Expect that JustFoo is Not test func")
	}
}
