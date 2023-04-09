package actually

import (
	"fmt"
	"reflect"
	"testing"
)

func (a *testingA) Nil(t *testing.T) *testingA {
	a.t = t

	if !a.isNil() {
		a.t.Helper()
		a.fail(fmt.Sprintf("Expect `nil`. But, got %#+v.", a.got))
	}

	return a
}

func (a *testingA) NotNil(t *testing.T) *testingA {
	a.t = t

	if a.isNil() {
		a.t.Helper()
		a.fail("Expect not `nil`. But, got `nil`.")
	}

	return a
}

func (a *testingA) isNil() bool {
	if a.got == nil {
		return true
	}

	return isSpecialNil(a.got)
}

func isSpecialNil(got any) bool {
	v := reflect.ValueOf(got)
	k := v.Kind()

	return isSpecialKind(k) && v.IsNil()
}

func isSpecialKind(k reflect.Kind) bool {
	// Special Kind is either one: Chan || Func || Interface || Map || Pointer || Slice
	// See https://github.com/golang/go/blob/8d68b388d4d1debec8d349adac58dd9f1cb03d25/src/reflect/type.go#L262-L267
	return k >= reflect.Chan && k <= reflect.Slice
}
