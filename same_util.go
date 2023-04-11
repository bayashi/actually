package actually

import (
	"bytes"
	"reflect"

	"github.com/bayashi/actually/report"
)

func reportForSame(a *testingA) *report.Report {
	return report.New().
		Expectf("Type:%T ,%#v", a.expect, a.expect).
		Gotf("Type:%T ,%#v", a.got, a.got)
}

func isFuncType(v any) bool {
	return v != nil && reflect.TypeOf(v).Kind() == reflect.Func
}

func objectsAreSame(expect any, got any) bool {
	if expect == nil || got == nil {
		return expect == got
	}

	exp, ok := expect.([]byte)
	if !ok {
		return reflect.DeepEqual(expect, got)
	}

	act, ok := got.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}

	return bytes.Equal(exp, act)
}
