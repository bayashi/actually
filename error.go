package actually

import (
	"testing"

	"github.com/bayashi/actually/report"
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
func (a *TestingA) NoError(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if a.got.RawValue() != nil {
		r := report.New()
		if !a.isTypeOfError() {
			r.Reason(reason_WrongType).Expect("It should be type of error")
		} else {
			r.Reason("Error happened").Expect("No error")
		}
		return a.fail(r.Gotf("Type:%Y, %#v", a.got, a.got))
	}

	return a
}

func (a *TestingA) isTypeOfError() bool {
	_, ok := a.got.RawValue().(error)

	return ok
}

// GotError sets the error value you actually got. GotError creates `*TestingA` and returns it.
func GotError(g error) *TestingA {
	return Got(g)
}

// GotError on *TestingA sets the error value you actually got.
func (a *TestingA) GotError(g error) *TestingA {
	return a.Got(g)
}
