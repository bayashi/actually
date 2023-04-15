package actually

import (
	"strings"
	"testing"

	"github.com/bayashi/actually/report"
	"github.com/bayashi/actually/trace"
)

type testingA struct {
	got       any
	setGot    bool
	expect    any
	setExpect bool
	t         *testing.T
	failNow   bool
}

// `Got` sets the value you actually got.
func Got(g any) *testingA {
	return &testingA{
		got: g,
		setGot: true,
	}
}

func (a *testingA) Got(g any) *testingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = g
	a.setGot = true

	return a
}

// `Expect` sets the value you expect to be the same as the one you got.
func Expect(e any) *testingA {
	return &testingA{
		expect: e,
		setExpect: true,
	}
}

func (a *testingA) Expect(e any) *testingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = e
	a.setExpect = true

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

// `Skip` provides shorthand to skip further tests within the same function for `-short` option.
/*
	func Test(t *testing.T) {
		Got(1).NotNil(t) // Run
		Skip(t)
		Got(2).NotNil(t) // Skip
		Got(3).NotNil(t) // Skip Also
	}
*/
func Skip(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}
