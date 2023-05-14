package actually

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/bayashi/actually/diff"
	"github.com/bayashi/actually/report"
)

func reportForSame(a *TestingA) *report.Report {
	r := report.New().
		Expectf(template_Dump, a.expect, a.expect).
		Gotf(template_Dump, a.got, a.got)

	if a.expect.IsStringType() {
		r = r.Expectf(template_DumpStringType, a.expect)
		if a.showRawData {
			r = r.ExpectAsRaw(fmt.Sprintf(template_DumpAsRawString, a.expect))
		}
	} else if a.expect.IsDumpableRawType() {
		r = r.Expectf(template_DumpStringType, a.expect)
		if a.showRawData {
			r = r.ExpectAsDump(a.expect.Dump())
		}
	}

	if a.got.IsStringType() {
		r = r.Gotf(template_DumpStringType, a.got)
		if a.showRawData {
			r = r.GotAsRaw(fmt.Sprintf(template_DumpAsRawString, a.got))
		}
	} else if a.got.IsDumpableRawType() {
		r = r.Gotf(template_DumpStringType, a.got)
		if a.showRawData {
			r = r.GotAsDump(a.got.Dump())
		}
	}

	return r
}

func reportForSameWithDiff(a *TestingA) *report.Report {
	d := diff.Diff(a.expect.RawValue(), a.got.RawValue())
	return reportForSame(a).Diff(d)
}

func reportForSameType(a *TestingA) *report.Report {
	return report.New().
		Expectf("Type: %Y", a.expect).
		Gotf("Type: %Y", a.got).
		Reason(reason_WrongType).
		Notice("SameType() just verifies the type. It doesn't care about the actual value")
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

func invalidCallForSame(a *TestingA) {
	if !a.setExpect {
		panic("You called kind of Same() method, but you forgot to call Expect().")
	}
	if !a.setGot {
		panic("You called kind of Same() method, but you forgot to call Got().")
	}
}
