package actually

import (
	"testing"

	"github.com/bayashi/actually/witness"
)

// Actual is an alias of Got.
func Actual(g any) *testingA {
	return Got(g)
}

// Actual is an alias of Got.
func (a *testingA) Actual(g any) *testingA {
	return a.Got(g)
}

// Want is an alias of Expect.
func Want(e any) *testingA {
	return Expect(e)
}

// Want is an alias of Expect.
func (a *testingA) Want(e any) *testingA {
	return a.Expect(e)
}

// Fatal is an alias of FailNow.
func (a *testingA) Fatal() *testingA {
	return a.FailNow()
}

// FatalOn is an alias of FailNowOn.
func FatalOn(t *testing.T) {
	FailNowOn(t)
}

// Fatal is an alias of FailNow
func Fatal(t *testing.T, reason string, got any, expect ...any) {
	witness.FailNow(t, reason, got, expect...)
}
