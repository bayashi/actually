package actually

import (
	"bytes"
	"reflect"

	w "github.com/bayashi/witness"
)

func reportForSame(a *testingA) *w.Witness {
	r := w.Expect(a.expect).Got(a.got).Name(a.name)
	if a.showRawData {
		r = r.ShowRaw()
	}

	return r
}

func reportForSameWithDiff(a *testingA) *w.Witness {
	return reportForSame(a).ShowDiff()
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

func invalidCallForSame(a *testingA) {
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
