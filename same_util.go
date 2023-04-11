package actually

import (
	"bytes"
	"reflect"
)

func isFuncType(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func isStringType(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.String
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
