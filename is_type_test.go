package actually

import (
	"testing"
)

func TestIsFuncType(t *testing.T) {
	tts := []struct {
		name   string
		arg    any
		expect bool
	}{
		{name: "Actual Func", arg: func() {}, expect: true},
		{name: "nil", arg: nil, expect: false},
		{name: "string", arg: "foo", expect: false},
	}
	for _, tt := range tts {
		if got := isFuncType(tt.arg); got != tt.expect {
			t.Errorf("isFuncType is wrong. Test Name:%v expect:%v, got:%v", tt.name, tt.expect, got)
		}
	}
}

func TestIsPointerType(t *testing.T) {
	i := 7
	tts := []struct {
		name   string
		arg    any
		expect bool
	}{
		{name: "Actual Pointer", arg: &i, expect: true},
		{name: "nil", arg: nil, expect: false},
		{name: "string", arg: "foo", expect: false},
	}
	for _, tt := range tts {
		if got := isPointerType(tt.arg); got != tt.expect {
			t.Errorf("isPointerType is wrong. Test Name:%v arg:%#v, expect:%v, got:%v", tt.name, tt.arg, tt.expect, got)
		}
	}
}

func TestIsTypeNil(t *testing.T) {
	if !isTypeNil(nil) {
		t.Errorf("<nil> should be nil type")
	}
	if isTypeNil(int(3)) {
		t.Errorf("int(3) should NOT be nil type")
	}
	// more cases?
}

func TestIsValidValue(t *testing.T) {
	if !isValidValue(1) {
		t.Errorf("1 should be valid value")
	}
	// more cases!
}
