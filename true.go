package actually

import (
	"reflect"
	"testing"

	"github.com/bayashi/actually/report"
)

// True method asserts that a test data you got is true value of boolean type.
/*
	actually.Got(true).True(t) // Truly pass
*/
func (a *TestingA) True(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !a.isBool() {
		r := report.New().
			Reason(reason_WrongType).
			Expect(message_ExpectTrue).
			Gotf("Type:%Y, %#v", a.got, a.got)
		return a.fail(r)
	}

	if a.got.RawValue() != true {
		r := report.New().
			Expect(message_ExpectTrue).
			Gotf("%#v", a.got)
		return a.fail(r)
	}

	return a
}

// False method asserts that a test data you got is false value of boolean type.
/*
	actually.Got(false).False(t) // pass
*/
func (a *TestingA) False(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !a.isBool() {
		r := report.New().
			Reason(reason_WrongType).
			Expect(message_ExpectFalse).
			Gotf("Type:%Y, %#v", a.got, a.got)
		return a.fail(r)
	}

	if a.got.RawValue() != false {
		r := report.New().
			Expect(message_ExpectFalse).
			Gotf("%#v", a.got)
		return a.fail(r)
	}

	return a
}

func (a *TestingA) isBool() bool {
	v := reflect.ValueOf(a.got.RawValue())
	k := v.Kind()

	return k == reflect.Bool
}
