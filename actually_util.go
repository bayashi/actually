package actually

// internal utils

import (
	"os"
	"strings"

	"github.com/bayashi/actually/report"
	"github.com/bayashi/actually/trace"
)

func (a *TestingA) fail(r *report.Report) *TestingA {
	a.t.Helper()
	r.Trace(traceinfo()).Function(a.t.Name() + "()").Name(a.name)
	a.t.Errorf("\n%s", r.Put())
	if a.failNow != nil && !*a.failNow {
		a.t.Fail()
	} else if (a.failNow != nil && *a.failNow) || len(os.Getenv(envKey_FailNow)) > 0 {
		a.t.FailNow()
	} else {
		a.t.Fail()
	}

	return a
}

var skipTraceRule = func(filepath string) bool {
	// Skip myself
	return strings.Contains(filepath, "actually.go")
}

func traceinfo() string {
	return strings.Join(trace.Info(skipTraceRule), traceSeparator)
}

func (a *TestingA) naming(testNames ...string) string {
	if a.name != "" {
		if len(testNames) > 0 {
			n := []string{a.name}
			n = append(n, testNames...)
			return strings.Join(n, ", ")
		} else {
			return a.name
		}
	} else {
		return strings.Join(testNames, ", ")
	}
}

func invalidCall(a *TestingA) {
	if !a.setGot {
		panic("You called assertion method, but you forgot to call Got().")
	}
}
