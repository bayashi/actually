package actually

import (
	"testing"

	w "github.com/bayashi/witness"
)

// NoError(t) method asserts that an error you got is NOt kind of error.
// Actually, this assertion is mostly same as Nil(t). But fail report of NoError(t) is helpful for a type of error.
/*
	g, err := someFunc()
	actually.GotError(err).NoError(t)
*/
// These are almost same as above code and below one.
/*
	g, err := someFunc()
	actually.Got(err).NoError(t)
*/
// Got(any) can accept any type of value, but GotError(error) can accept ONLY a type of error.
// It's more strict when you use GotError(error) to test a type of error.
func (a *testingA) NoError(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if a.got != nil {
		var reason string
		w := w.Got(a.got)
		if !a.isTypeOfError() {
			reason = reason_WrongType
			w.Message(notice_Label, "It should be type of error")
		} else {
			reason = reason_UnexpectedlyError
			w.Message(notice_Label, "No error")
		}
		return a.fail(w, reason)
	}

	return a
}

func (a *testingA) isTypeOfError() bool {
	_, ok := a.got.(error)

	return ok
}

// GotError sets the error value you actually got. GotError creates `*testingA` and returns it.
func GotError(g error) *testingA {
	return Got(g)
}

// GotError on *testingA sets the error value you actually got.
func (a *testingA) GotError(g error) *testingA {
	return a.Got(g)
}
