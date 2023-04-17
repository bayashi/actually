// Actual trace_test.go is in ../report directory.
package trace

import "testing"

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
