package actually

// helper functions

import (
	"fmt"
	"testing"

	"github.com/bayashi/actually/witness"
	"github.com/bayashi/actually/witness/obj"

	"github.com/yassinebenaid/godump"
)

func failNowPtr(v bool) *bool {
	return &v
}

// `FailNow` turns a flag on to stop further test execution immediately if one test fails.
// NOTE that FailNow method should be called after `Got` or `Expect`.
/*
	actually.Got(something).FailNow().Nil(t) // Fail now for only this test
*/
func (a *testingA) FailNow() *testingA {
	a.failNow = failNowPtr(true)

	return a
}

// FailNowOn function turns an ENV flag on to stop further test execution immediately if one test fails.
// This switch is enabled within the test. Not only during function.
/*
	func Test(t *testing.T) {
		actually.FailNowOn(t)
		actually.Got(something).Nil(t)                    // Fail Now
		actually.Got(something).Expect(something).Same(t) // Fail Now
	}
*/
// Warning: Do not use FailNowOn along with t.Parallel.
func FailNowOn(t *testing.T) {
	t.Setenv(envKey_FailNow, envKey_FailNow)
}

// FailNotNowOn function turns an ENV flag off to stop further test execution immediately if one test fails.
// If you want to turn the ENV flag on, then you should call `FailNowOn`.
/*
	func Test(t *testing.T) {
		// turn on to fail right now
		actually.FailNowOn(t)
		actually.Got(something).Nil(t)                    // Fail Now
		actually.Got(something).Expect(something).Same(t) // Fail Now

		// turn off
		actually.FailNotNowOn(t)
		actually.Got(something).Nil(t)                    // NOT Fail Now
		actually.Got(something).Expect(something).Same(t) // NOT Fail Now

		// Fail Now by FailNow() in the chain
		actually.Got(something).FailNow().Nil(t)          // Fail Now

		// Again, turn on to fail right now
		actually.FailNowOn(t)
		actually.Got(something).Nil(t)                    // Fail Now
	}
*/
// This switch is enabled within the test. Not only during function.
func FailNotNowOn(t *testing.T) {
	t.Setenv(envKey_FailNow, "")
}

// FailNotNow turns a flag off, so that even if the test fails, execution does not stop immediately.
/*
   It behaves this way by default. If you want the opposite behavior, call `FailNow` method.
   NOTE that FailNotNow method should be called after `Got` or `Expect`.
*/
// Deprecated: Anyone uses? This method will be removed in the near future.
func (a *testingA) FailNotNow() *testingA {
	a.failNow = failNowPtr(false)

	return a
}

// X turns on a flag to show test values as raw in a fail report.
func (a *testingA) X() *testingA {
	a.showRawData = true

	return a
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
func Skip(t *testing.T, skipReasons ...any) {
	if testing.Short() {
		t.Skip(skipReasons...)
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

// Fi doesn NOT return `*testingA`, returns whether a test failed instead.
// If even there is only one test fails in a one chain to test, then it will be `true` anyway.
/*
	func Test(t *testing.T) {
		love := true
		// Fi returns `true` if either `NotNil(t)`, `True(t)` or 'Same(t)' failed
		if res := a.Got(love).NotNil(t).True(t).Expect(true).Same(t).Fi(); !res {
			// your own some action on fail
		}
	}
*/
func (a *testingA) Fi() bool {
	return a.failed
}

// Diff is a helper function to get a diff string of 2 objects for debugging
func Diff(a any, b any) string {
	return witness.Diff(a, b)
}

var defaultDumper godump.Dumper

// Dump is a helper function to get a dumped string of objects for debugging
func Dump(a ...any) string {
	if len(a) == 0 {
		return ""
	}

	if obj.DUMPER == nil {
		obj.DUMPER = func(d any) string {
			return defaultDumper.Sprint(d)
		}
	}

	if len(a) == 1 {
		return witness.Dump(a[0]) + "\n"
	}

	result := ""
	for i, v := range a {
		result += fmt.Sprintf("[%d]\n%s\n", i, witness.Dump(v))
	}

	return result
}

// Debug is a helper function to show debug info on fail
func (a *testingA) Debug(label string, info ...any) *testingA {
	a.debugInfo.mutex.Lock()
	defer a.debugInfo.mutex.Unlock()
	a.debugInfo.debugInfo = append(a.debugInfo.debugInfo, map[string][]any{label: info})
	return a
}

// Fail is to show decorated fail report. (Actual shortcut to witness.Fail)
/*
	if g != e {
		actually.Fail(t, "Not same", g, e)
	}
*/
func Fail(t *testing.T, reason string, got any, expect ...any) {
	t.Helper()
	witness.Fail(t, reason, got, expect...)
}

// FailNow is to show decorated fail report by t.Fatal. (Actual shortcut to witness.FailNow)
/*
	if g != e {
		actually.FailNow(t, "Not same", g, e)
	}
*/
func FailNow(t *testing.T, reason string, got any, expect ...any) {
	t.Helper()
	witness.FailNow(t, reason, got, expect...)
}
