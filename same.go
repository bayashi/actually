package actually

import (
	"testing"

	"github.com/bayashi/actually/report"
)

func (a *testingA) SameString(t *testing.T) *testingA {
	a.t = t

	// validate
	if isFuncType(a.got) {
		a.t.Helper()
		r := report.New().
			Reason("`Got` value is type of function. It cannot be used in Same*() method").
			Expectf("Type:%T ,%#v", a.expect, a.expect).
			Gotf("Type:%T ,%#v", a.got, a.got).
			Message("`Got` value should be string")
		return a.fail(r)
	}
	if isFuncType(a.expect) {
		a.t.Helper()
		r := report.New().
			Reason("`Expect` value is type of function. It cannot be used in Same*() method").
			Expectf("Type:%T ,%#v", a.expect, a.expect).
			Gotf("Type:%T ,%#v", a.got, a.got).
			Message("`Expect` value should be string")
		return a.fail(r)
	}

	// Got and Expect should be string
	if !isStringType(a.got) {
		a.t.Helper()
		r := report.New().
			Reason("`Got` value is Not string").
			Expectf("Type:%T ,%#v", a.expect, a.expect).
			Gotf("Type:%T ,%#v", a.got, a.got).
			Message("`Got` value should be string")
		return a.fail(r)
	}
	if !isStringType(a.expect) {
		a.t.Helper()
		r := report.New().
			Reason("`Expect` value is Not string").
			Expectf("Type:%T ,%#v", a.expect, a.expect).
			Gotf("Type:%T ,%#v", a.got, a.got).
			Message("`Expect` value should be string")
		return a.fail(r)
	}

	// compare
	if !objectsAreSame(a.expect, a.got) {
		a.t.Helper()
		r := report.New().
			Reason("Not same string").
			Expectf("Type:%T ,%#v", a.expect, a.expect).
			Gotf("Type:%T ,%#v", a.got, a.got)
		return a.fail(r)
	}

	return a
}
