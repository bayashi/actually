package obj

import (
	"bufio"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

var DUMPER func(d any) string

type Object struct {
	value  any
	touch  bool
	maxLen int
	kind   reflect.Kind
	dumper func(d any) string
}

func NewObject(v any) *Object {
	return &Object{
		touch:  true,
		value:  v,
		maxLen: bufio.MaxScanTokenSize,
		kind:   reflect.ValueOf(v).Kind(),
	}
}

func NewObjectWithDumper(v any, dumper func(d any) string) *Object {
	return &Object{
		touch:  true,
		value:  v,
		maxLen: bufio.MaxScanTokenSize,
		kind:   reflect.ValueOf(v).Kind(),
		dumper: dumper,
	}
}

func NewObjectWithMaxLen(v any, len int) *Object {
	return &Object{
		touch:  true,
		value:  v,
		maxLen: len,
		kind:   reflect.ValueOf(v).Kind(),
	}
}

func (o *Object) Touch() bool {
	return o.touch
}

// Return raw value as any type that should be same as original value
func (o *Object) AsRawValue() any {
	return o.value
}

func (o *Object) AsType() string {
	return fmt.Sprintf("%T", o.value)
}

// Return the value as string that was converted by fmt.Sprintf for each type
func (o *Object) AsString() string {
	if !o.touch {
		return ""
	}

	switch o.value.(type) {
	case string, []byte:
		return fmt.Sprintf("%q", o.value)
	case nil:
		return "<nil>"
	case bool:
		return fmt.Sprintf("<%t>", o.value)
	case *bool:
		return o.AsDumpString()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", o.value)
	case float32, float64, complex64, complex128:
		return fmt.Sprintf("%g", o.value)
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64,
		*float32, *float64, *complex64, *complex128:
		return o.AsDumpString()
	case error:
		return o.AsDumpString()
	default:
		if o.IsPointerType() {
			return fmt.Sprintf("%p, %#v", o.value, o.value)
		} else {
			return fmt.Sprintf("%#v", o.value)
		}
	}
}

func (o *Object) AsFmtString() string {
	return fmt.Sprintf("%#v", o.value)
}

// Return dump string even if the value would be any value
func (o *Object) AsDumpString() string {
	if o.dumper != nil {
		return o.dumper(o.value)
	}

	if DUMPER != nil {
		return DUMPER(o.value)
	}

	return spew.Sdump(o.value)
}

// Return boolean whether the value is string type
func (o *Object) IsStringType() bool {
	if o.touch && o.value != nil {
		if o.kind == reflect.String {
			return true
		}
	}

	return false
}

// Return boolean whether the value is dumpable
func (o *Object) IsDumpableRawType() bool {
	if o.touch && o.value != nil {
		if o.kind == reflect.Struct || o.kind == reflect.Map || o.kind == reflect.Slice || o.kind == reflect.Array {
			return true
		}
	}

	return false
}

// Return boolean whether the value is a pointer
func (o *Object) IsPointerType() bool {
	return o.value != nil && reflect.TypeOf(o.value).Kind() == reflect.Pointer
}

// Return boolean whether the value is a struct
func (o *Object) IsStructType() bool {
	return o.value != nil && reflect.TypeOf(o.value).Kind() == reflect.Struct
}

// format and truncate. For `fmt/print.go` Formatter interface
func (o *Object) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		flag := ""
		if s.Flag('+') {
			flag += "+"
		}
		if s.Flag('#') {
			flag += "#"
		}
		fmt.Fprint(s, truncate(o.value, fmt.Sprintf("%%%sv", flag), o.maxLen))
	case 's':
		fmt.Fprint(s, truncate(o.value, "%s", o.maxLen))
	}
}

// Truncate string as format
func truncate(v any, format string, maxLen int) string {
	str := fmt.Sprintf(format, v)
	if len(str) > maxLen {
		str = str[0:maxLen] + "<... truncated>"
	}

	return str
}
