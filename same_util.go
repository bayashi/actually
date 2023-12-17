package actually

import (
	"bytes"
	"reflect"
	"strings"

	w "github.com/bayashi/witness"
)

func reportForSame(a *TestingA) *w.Witness {
	r := w.Expect(a.expect).Got(a.got)
	if a.showRawData {
		r = r.ShowRaw()
	}

	return r
}

func reportForSameWithDiff(a *TestingA) *w.Witness {
	return reportForSame(a).ShowDiff()
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

func isTypeNumber(v any) bool {
	typ := reflect.TypeOf(v).Name()
	return strings.HasPrefix(typ, "int") || strings.HasPrefix(typ, "uint") || strings.HasPrefix(typ, "float")
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

func convert2float64(a any) float64 {
	var f float64

	b := reflect.TypeOf(a).Kind()

	// https://pkg.go.dev/reflect#Kind
	if b >= 2 && b <= 6 {
		f = float64(reflect.ValueOf(a).Int())
	} else if b >= 7 && b <= 11 {
		f = float64(reflect.ValueOf(a).Uint())
	} else if b == 13 || b == 14 {
		f = reflect.ValueOf(a).Float()
	}

	return f
}
