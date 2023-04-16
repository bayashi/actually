package actually

import (
	"reflect"
	"testing"

	"github.com/bayashi/actually/report"
)

func (a *testingA) Nil(t *testing.T) *testingA {
	a.t = t

	if !a.isNil() {
		a.t.Helper()
		r := report.New().
			Expect("nil").
			Gotf("Type:%T, %#v", a.got, a.got)
		return a.fail(r)
	}

	return a
}

func (a *testingA) NotNil(t *testing.T) *testingA {
	a.t = t

	if a.isNil() {
		a.t.Helper()
		r := report.New().
			Expect("Not nil").
			Got("nil")
		return a.fail(r)
	}

	return a
}

func (a *testingA) isNil() bool {
	if a.got.RawValue() == nil {
		return true
	}

	return isSpecialNil(a.got.RawValue())
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
