package actually

import (
	"strings"
	"testing"

	"github.com/bayashi/actually/report"
	"github.com/bayashi/actually/testobject"
	"github.com/bayashi/actually/trace"
)

// testingA is a context of the test
type testingA struct {
	got         *testobject.TestObject
	setGot      bool
	expect      *testobject.TestObject
	setExpect   bool
	t           *testing.T
	failNow     bool
	showRawData bool
	name        string
}

// Got sets the value you actually got.
func Got(g any) *testingA {
	return &testingA{
		got:    testobject.NewTestObject(g, 0),
		setGot: true,
	}
}

func (a *testingA) Got(g any) *testingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = testobject.NewTestObject(g, 0)
	a.setGot = true

	return a
}

// Expect sets the value you expect to be the same as the one you got.
func Expect(e any) *testingA {
	return &testingA{
		expect:    testobject.NewTestObject(e, 0),
		setExpect: true,
	}
}

func (a *testingA) Expect(e any) *testingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = testobject.NewTestObject(e, 0)
	a.setExpect = true

	return a
}

// FailNotNow turns a flag so that even if the test fails, execution does not stop immediately.
/*
    It behaves this way by default. If you want the opposite behavior, call `FailNow` method.
    NOTE that FailNotNow method should be called after `Got` or `Expect`.
*/
func (a *testingA) FailNotNow() *testingA {
	a.failNow = false

	return a
}

// `FailNow` turns on a flag to stop further test execution immediately if one test fails
/*
	NOTE that FailNow method should be called after `Got` or `Expect`.
*/
func (a *testingA) FailNow() *testingA {
	a.failNow = true

	return a
}

func (a *testingA) fail(r *report.Report) *testingA {
	a.t.Helper()
	r.Trace(traceinfo()).Function(a.t.Name() + "()").Name(a.name)
	a.t.Errorf("\n%s", r.Put())
	if a.failNow {
		a.t.FailNow()
	} else {
		a.t.Fail()
	}

	return a
}

// X turns on a flag to show test values as raw in a fail report.
func (a *testingA) X() *testingA {
	a.showRawData = true

	return a
}

var skipTraceRule = func(filepath string) bool {
	// Skip myself
	return strings.Contains(filepath, "actually.go")
}

func traceinfo() string {
	return strings.Join(trace.Info(skipTraceRule), traceSeparator)
}

// Skip provides shorthand to skip further tests within the same function for `-short` option.
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

// Set test name spefically.
/*
	You can also set a test name on assertion methods.

	actually.Got(a).Expect(b).Same(t, "Test Name")
*/
func (a *testingA) Name(n string) *testingA {
	a.name = n

	return a
}

func (a *testingA) naming(testNames ...string) string {
	if a.name != "" {
		if  len(testNames) > 0 {
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
