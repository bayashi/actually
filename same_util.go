package actually

import (
	"bytes"
	"reflect"

	"github.com/bayashi/actually/diff"
	"github.com/bayashi/actually/report"
)

func reportForSame(a *testingA) *report.Report {
	r := report.New()
	if a.expect.IsStringType() {
		r.Expectf(template_DumpStringType, a.expect, a.expect, a.expect)
	} else {
		r.Expectf(template_Dump, a.expect, a.expect)
	}
	if a.got.IsStringType() {
		r.Gotf(template_DumpStringType, a.got, a.got, a.got)
	} else {
		r.Gotf(template_Dump, a.got, a.got)
	}

	return r
}

func reportForSameWithDiff(a *testingA) *report.Report {
	d := diff.Diff(a.expect.RawValue(), a.got.RawValue())
	return reportForSame(a).Diff(d)
}

func isFuncType(v any) bool {
	return v != nil && reflect.TypeOf(v).Kind() == reflect.Func
}

func isPointerType(v any) bool {
	return v != nil && reflect.TypeOf(v).Kind() == reflect.Pointer
}

func isTypeNil(v any) bool {
	return reflect.TypeOf(v) == nil
}

func isValidValue(v any) bool {
	return reflect.ValueOf(v).IsValid()
}

// Just confirming only types are convertible or not
func objectsAreConvertible(expectv any, gotv any) bool {
	return reflect.ValueOf(expectv).Type().ConvertibleTo(reflect.TypeOf(gotv)) &&
		reflect.ValueOf(gotv).Type().ConvertibleTo(reflect.TypeOf(expectv))
}

func isSameConvertedValueAsOther(expectv any, gotv any) bool {
	return reflect.DeepEqual(reflect.ValueOf(expectv).Convert(reflect.TypeOf(gotv)).Interface(), gotv)
}

func objectsAreSameType(expectv any, gotv any) bool {
	return reflect.TypeOf(gotv) == reflect.TypeOf(expectv)
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
