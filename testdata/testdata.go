// testdata package provides utilities to handle data of testing
package testdata

import (
	"bufio"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type TestData struct {
	value  any
	maxLen int
}

func NewTestData(v any, maxLen int) *TestData {
	max := bufio.MaxScanTokenSize
	if maxLen < 0 {
		maxLen = max + maxLen
	} else if maxLen == 0 || maxLen > bufio.MaxScanTokenSize {
		maxLen = bufio.MaxScanTokenSize
	}

	return &TestData{
		value:  v,
		maxLen: maxLen,
	}
}

func (td *TestData) RawValue() any {
	return td.value
}

func (td *TestData) IsStringType() bool {
	if v := td.RawValue(); v != nil {
		k := reflect.ValueOf(v).Kind()
		if k == reflect.String {
			return true
		}
	}

	return false
}

func (td *TestData) IsDumpableRawType() bool {
	if v := td.RawValue(); v != nil {
		k := reflect.ValueOf(v).Kind()
		if k == reflect.Struct || k == reflect.Map || k == reflect.Slice || k == reflect.Array {
			return true
		}
	}

	return false
}

func (td *TestData) Dump() string {
	return spew.Sdump(td.RawValue())
}

// For `fmt/print.go` Formatter interface
func (td *TestData) Format(s fmt.State, verb rune) {
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

func truncate(data interface{}, format string, maxLen int) string {
	v := fmt.Sprintf(format, data)
	if len(v) > maxLen {
		v = v[0:maxLen] + "<... truncated>"
	}

	return v
}
