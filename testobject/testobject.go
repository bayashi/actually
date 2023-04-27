// testobject package provides utilities to handle data of testing
package testobject

import (
	"bufio"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type TestObject struct {
	value  any
	maxLen int
}

func NewTestObject(v any, maxLen int) *TestObject {
	max := bufio.MaxScanTokenSize
	if maxLen < 0 {
		maxLen = max + maxLen
	} else if maxLen == 0 || maxLen > bufio.MaxScanTokenSize {
		maxLen = bufio.MaxScanTokenSize
	}

	return &TestObject{
		value:  v,
		maxLen: maxLen,
	}
}

func (td *TestObject) RawValue() any {
	return td.value
}

func (td *TestObject) IsStringType() bool {
	if v := td.RawValue(); v != nil {
		k := reflect.ValueOf(v).Kind()
		if k == reflect.String {
			return true
		}
	}

	return false
}

func (td *TestObject) IsDumpableRawType() bool {
	if v := td.RawValue(); v != nil {
		k := reflect.ValueOf(v).Kind()
		if k == reflect.Struct || k == reflect.Map || k == reflect.Slice || k == reflect.Array {
			return true
		}
	}

	return false
}

func (td *TestObject) Dump() string {
	return spew.Sdump(td.RawValue())
}

// For `fmt/print.go` Formatter interface
func (td *TestObject) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		flag := ""
		if s.Flag('+') {
			flag += "+"
		}
		if s.Flag('#') {
			flag += "#"
		}
		d := truncate(td.RawValue(), fmt.Sprintf("%%%sv", flag), td.maxLen)
		fmt.Fprint(s, d)
	case 'Y':
		d := fmt.Sprintf("%T", td.RawValue())
		fmt.Fprint(s, d)
	case 's':
		d := truncate(td.RawValue(), "%s", td.maxLen)
		fmt.Fprint(s, d)
	}
}

func truncate(data any, format string, maxLen int) string {
	v := fmt.Sprintf(format, data)
	if len(v) > maxLen {
		v = v[0:maxLen] + "<... truncated>"
	}

	return v
}
