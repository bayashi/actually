package actually

import (
	"testing"

	"github.com/bayashi/witness"
)

// Actual is an alias of Got.
func Actual(g any) *TestingA {
	return Got(g)
}

// Actual is an alias of Got.
func (a *TestingA) Actual(g any) *TestingA {
	return a.Got(g)
}

// Want is an alias of Expect.
func Want(e any) *TestingA {
	return Expect(e)
}

// Want is an alias of Expect.
func (a *TestingA) Want(e any) *TestingA {
	return a.Expect(e)
}

// Fatal is an alias of FailNow.
func (a *TestingA) Fatal() *TestingA {
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
