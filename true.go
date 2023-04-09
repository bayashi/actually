package actually

import (
	"fmt"
	"reflect"
	"testing"
)

func (a *testingA) True(t *testing.T) *testingA {
	a.t = t

	if !a.isBool() {
		a.t.Helper()
		a.fail(fmt.Sprintf("Expect boolean for `True()`, but got %T `%#+v`.", a.got, a.got))
	}

	if a.got != true {
		a.t.Helper()
		a.fail(fmt.Sprintf("Expect true, but got %#+v.", a.got))
	}

	return a
}

func (a *testingA) False(t *testing.T) *testingA {
	a.t = t

	if !a.isBool() {
		a.t.Helper()
		a.fail(fmt.Sprintf("Expect boolean for `False()`, but got %T `%#+v`.", a.got, a.got))
	}

	if a.got != false {
		a.t.Helper()
		a.fail(fmt.Sprintf("Expect false, but got %#+v.", a.got))
	}

	return a
}

func (a *testingA) isBool() bool {
	v := reflect.ValueOf(a.got)
	k := v.Kind()

	return k == reflect.Bool
}
