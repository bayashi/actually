package actually

import (
	"bytes"
	"reflect"

	"github.com/bayashi/actually/report"
)

func reportForSame(a *testingA) *report.Report {
	return report.New().
		Expectf("Type:%T, %#v", a.expect, a.expect).
		Gotf("Type:%T, %#v", a.got, a.got)
}

func isFuncType(v any) bool {
	return v != nil && reflect.TypeOf(v).Kind() == reflect.Func
}

func isPointerType(v any) bool {
	return v != nil && reflect.TypeOf(v).Kind() == reflect.Pointer
}

func objectsAreSame(expectv any, gotv any) bool {
	if expectv == nil || gotv == nil {
		return expectv == gotv
	}

	exp, ok := expectv.([]byte)
	if !ok {
		return reflect.DeepEqual(expectv, gotv)
	}

	act, ok := gotv.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}

	return bytes.Equal(exp, act)
}
