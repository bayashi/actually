package actually

// helper functions

import (
	"testing"

	"github.com/bayashi/witness"
)

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

// X turns on a flag to show test values as raw in a fail report.
func (a *TestingA) X() *TestingA {
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
func (a *TestingA) Name(n string) *TestingA {
	a.name = n

	return a
}

// Diff is a helper function to get a diff string of 2 objects for debugging
func Diff(a any, b any) string {
	return witness.Diff(a, b)
}
