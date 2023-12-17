package actually

import (
	"reflect"
	"testing"

	w "github.com/bayashi/witness"
)

// Nil asserts that a test data you got is <nil>
/*
	actually.Got(a).Nil(t) // If `a` is <nil>, then pass.
*/
func (a *TestingA) Nil(t *testing.T, testNames ...string) *TestingA {
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !a.isNil() {
		w := w.Expect("<nil>").Got(a.got)
		return a.fail(w, "Expected <nil>, but it was NOT <nil>")
	}

	return a
}

// NotNil asserts that a test data you got is NOT <nil>
/*
	actually.Got(a).NotNil(t) // If `a` is NOT <nil>, then pass.
*/
func (a *TestingA) NotNil(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if a.isNil() {
		w := w.Expect("Not <nil>").Got("<nil>")
		return a.fail(w, reason_ExpectIsNotNil)
	}

	return a
}

func (a *TestingA) isNil() bool {
	if a.got == nil {
		return true
	}

	return isSpecialNil(a.got)
}

func isSpecialNil(gotv any) bool {
	v := reflect.ValueOf(gotv)
	k := v.Kind()

	return isSpecialKind(k) && v.IsNil()
}

func isSpecialKind(k reflect.Kind) bool {
	// Special Kind is either one: Chan || Func || Interface || Map || Pointer || Slice || UnsafePointer
	// See https://github.com/golang/go/blob/8d68b388d4d1debec8d349adac58dd9f1cb03d25/src/reflect/type.go#L262-L267
	return (k >= reflect.Chan && k <= reflect.Slice) || k == reflect.UnsafePointer
}
