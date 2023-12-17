package actually

import (
	"reflect"
	"testing"
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
		w := reportForSame(a).Message("Notice", "It should be boolean")
		return a.fail(w, reason_WrongType)
	}

	if a.got != true {
		return a.fail(reportForSame(a), message_ExpectTrue)
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
		w := reportForSame(a).Message("Notice", "It should be boolean")
		return a.fail(w, reason_WrongType)
	}

	if a.got != false {
		return a.fail(reportForSame(a), message_ExpectFalse)
	}

	return a
}

func (a *TestingA) isBool() bool {
	v := reflect.ValueOf(a.got)
	k := v.Kind()

	return k == reflect.Bool
}
