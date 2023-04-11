package actually

import (
	"reflect"
	"testing"

	"github.com/bayashi/actually/report"
)

func (a *testingA) True(t *testing.T) *testingA {
	a.t = t

	if !a.isBool() {
		a.t.Helper()
		r := report.New().
			Reason(FailReason_WrongType).
			Expect("Boolean type true").
			Gotf("Type:%T, %#v", a.got, a.got)
		return a.fail(r)
	}

	if a.got != true {
		a.t.Helper()
		r := report.New().
			Expect("true").
			Gotf("%#v", a.got)
		return a.fail(r)
	}

	return a
}

func (a *testingA) False(t *testing.T) *testingA {
	a.t = t

	if !a.isBool() {
		a.t.Helper()
		r := report.New().
			Reason(FailReason_WrongType).
			Expect("Boolean type false").
			Gotf("Type:%T ,%#v", a.got, a.got)
		return a.fail(r)
	}

	if a.got != false {
		a.t.Helper()
		r := report.New().
			Expect("false").
			Gotf("%#v", a.got)
		return a.fail(r)
	}

	return a
}

func (a *testingA) isBool() bool {
	v := reflect.ValueOf(a.got)
	k := v.Kind()

	return k == reflect.Bool
}
