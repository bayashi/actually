package actually

import (
	"strings"
	"testing"

	"github.com/bayashi/actually/report"
	"github.com/bayashi/actually/trace"
)

type testingA struct {
	got any
	expect any
	t *testing.T
	failNow bool
}

// `Got` sets the value you actually got.
func Got(g any) *testingA {
	return &testingA{
		got: g,
	}
}

func (a *testingA) Got(g any) *testingA {
	a.got = g

	return a
}

// `Expect` sets the value you expect to be the same as the one you got.
func Expect(e any) *testingA {
	return &testingA{
		expect: e,
	}
}

func (a *testingA) Expect(e any) *testingA {
	a.expect = e

	return a
}

// `FailNotNow` turns a flag so that even if the test fails, execution does not stop immediately.
//
// It behaves this way by default. If you want the opposite behavior, call `FailNow` method.
func (a *testingA) FailNotNow() *testingA {
	a.failNow = false

	return a
}

// `FailNow` turns on a flag to stop further test execution immediately if one test fails
func (a *testingA) FailNow() *testingA {
	a.failNow = true

	return a
}

func (a *testingA) fail(r *report.Report) *testingA {
	a.t.Helper()
	r.Trace(traceinfo()).Name(a.t.Name() + "()")
	a.t.Errorf("\n%s", r.Put())
	if a.failNow {
		a.t.FailNow()
	} else {
		a.t.Fail()
	}

	return a
}

var SkipTraceRule = func(file string) bool {
	// Skip myself
	return strings.Contains(file, "actually.go")
}

func traceinfo() string {
	return strings.Join(trace.Info(SkipTraceRule), traceSeparator)
}
