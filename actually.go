// Yet another pithy testing framework `actually`
package actually

import (
	"os"
	"strings"
	"testing"

	"github.com/bayashi/actually/diff"
	"github.com/bayashi/actually/report"
	"github.com/bayashi/actually/testobject"
	"github.com/bayashi/actually/trace"
)

// TestingA is a context of the test
type TestingA struct {
	got         *testobject.TestObject
	setGot      bool
	expect      *testobject.TestObject
	setExpect   bool
	t           *testing.T
	failNow     *bool
	showRawData bool
	name        string
}

// Got sets the value you actually got. Got() creates *TestingA and returns it.
func Got(g any) *TestingA {
	return &TestingA{
		got:    testobject.NewTestObject(g, 0),
		setGot: true,
	}
}

// Got sets the value you actually got.
func (a *TestingA) Got(g any) *TestingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = testobject.NewTestObject(g, 0)
	a.setGot = true

	return a
}

// Expect sets the value you expect to be the same as the one you got.
// Expect creates *TestingA and returns it.
func Expect(e any) *TestingA {
	return &TestingA{
		expect:    testobject.NewTestObject(e, 0),
		setExpect: true,
	}
}

// Expect sets the value you expect to be the same as the one you got.
func (a *TestingA) Expect(e any) *TestingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = testobject.NewTestObject(e, 0)
	a.setExpect = true

	return a
}

func failNowPtr(v bool) *bool {
	return &v
}

// FailNotNow turns a flag so that even if the test fails, execution does not stop immediately.
/*
   It behaves this way by default. If you want the opposite behavior, call `FailNow` method.
   NOTE that FailNotNow method should be called after `Got` or `Expect`.
*/
func (a *TestingA) FailNotNow() *TestingA {
	a.failNow = failNowPtr(false)

	return a
}

// FailNotNowOn function turns off an ENV flag to stop further test execution immediately if one test fails.
/*
	func Test(t *testing.T) {
		actually.FailNowOn(t)
		actually.Got(something).Nil(t)                    // Fail Now
		actually.Got(something).Expect(something).Same(t) // Fail Now

		actually.FailNotNowOn(t)
		actually.Got(something).Nil(t)                    // NOT Fail Now
		actually.Got(something).Expect(something).Same(t) // NOT Fail Now

		actually.Got(something).FailNow().Nil(t)          // Fail Now
	}
*/
func FailNotNowOn(t *testing.T) {
	t.Setenv(envKey_FailNow, "")
}

// `FailNow` turns on a flag to stop further test execution immediately if one test fails
// NOTE that FailNow method should be called after `Got` or `Expect`.
/*
	actually.Got(something).FailNow().Nil(t) // Fail now for only this test
*/
func (a *TestingA) FailNow() *TestingA {
	a.failNow = failNowPtr(true)

	return a
}

// FailNowOn function turns on an ENV flag to stop further test execution immediately if one test fails.
/*
	func Test(t *testing.T) {
		actually.FailNowOn(t)
		actually.Got(something).Nil(t)                    // Fail Now
		actually.Got(something).Expect(something).Same(t) // Fail Now
	}
*/
func FailNowOn(t *testing.T) {
	t.Setenv(envKey_FailNow, envKey_FailNow)
}

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

// X turns on a flag to show test values as raw in a fail report.
func (a *TestingA) X() *TestingA {
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
func (a *TestingA) Name(n string) *TestingA {
	a.name = n

	return a
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

// Diff is a helper function to get a diff string of 2 objects for debugging
func Diff(a any, b any) string {
	return diff.DiffSimple(a, b)
}

func invalidCall(a *TestingA) {
	if !a.setGot {
		panic("You called assertion method, but you forgot to call Got().")
	}
}