package actually

import (
	"reflect"
	"strings"
)

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
