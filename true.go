package actually

import (
	"reflect"
	"testing"

	w "github.com/bayashi/witness"
)

// True method asserts that a test data you got is true value of boolean type.
/*
	actually.Got(true).True(t) // Truly pass
*/
func (a *testingA) True(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !a.isBool() {
		w := w.Got(a.got).Message(notice_Label, "It should be boolean")
		return a.fail(w, reason_WrongType)
	}

	if a.got != true {
		return a.fail(w.Got(a.got), message_ExpectTrue)
	}

	return a
}

// False method asserts that a test data you got is false value of boolean type.
/*
	actually.Got(false).False(t) // pass
*/
func (a *testingA) False(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !a.isBool() {
		w := w.Got(a.got).Message(notice_Label, "It should be boolean")
		return a.fail(w, reason_WrongType)
	}

	if a.got != false {
		return a.fail(w.Got(a.got), message_ExpectFalse)
	}

	return a
}

func (a *testingA) isBool() bool {
	v := reflect.ValueOf(a.got)
	k := v.Kind()

	return k == reflect.Bool
}
