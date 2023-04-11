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

func Got(g any) *testingA {
	return &testingA{
		got: g,
	}
}

func (a *testingA) Got(g any) *testingA {
	a.got = g

	return a
}

func Expect(e any) *testingA {
	return &testingA{
		expect: e,
	}
}

func (a *testingA) Expect(e any) *testingA {
	a.expect = e

	return a
}

func Want(e any) *testingA {
	return Expect(e)
}
func (a *testingA) Want(e any) *testingA {
	return a.Expect(e)
}

func (a *testingA) FailNotNow() *testingA {
	a.failNow = false

	return a
}

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
	return strings.Join(trace.Info(SkipTraceRule), TraceSeparator)
}
