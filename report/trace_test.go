package report

import (
	"regexp"
	"testing"

	"github.com/bayashi/actually/trace"
)

func TestInfo(t *testing.T) {
	skipRule := func(file string) bool { return false }
	trace := trace.Info(skipRule)
	if len(trace) != 1 {
		t.Error("trace length should be 1.")
	}
	var traceRegexp = regexp.MustCompile(`actually/report/trace_test\.go:\d+$`)
	if !traceRegexp.MatchString(trace[0]) {
		t.Errorf("trace was not match Regexp:`%s`, Got:`%s`", traceRegexp.String(), trace[0])
	}
}
