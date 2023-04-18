package actually

import (
	"testing"
	"time"
)

// func TestReportForSame(t *testing.T) {
// 	a := Got("Hello.\nThanks!").Expect("Hello.\nThank you!").ShowRawData()
// 	out := reportForSame(a).Put()
// 	fmt.Print(out)
// }

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

func TestObjectsAreConvertible(t *testing.T) {
	if !objectsAreConvertible(int(2), float32(2.0)) {
		t.Errorf("Int and float32 should be Convertible")
	}
	if !objectsAreConvertible("foo", `foo`) {
		t.Errorf("String and rune should be Convertible")
	}
	if objectsAreConvertible(int(2), complex128(2)) {
		t.Errorf("Int and complex128 are NOT Convertible")
	}
}

func TestIsSameConvertedValueAsOther(t *testing.T) {
	if !isSameConvertedValueAsOther(int(2), float32(2.0)) {
		t.Errorf("Int(2) and float32(2.0) should be Convertible and same")
	}
	if isSameConvertedValueAsOther(int(2), float32(2.1)) {
		t.Errorf("Int(2) and float32(2.1) are Convertible, but not same")
	}
	if !isSameConvertedValueAsOther("foo", `foo`) {
		t.Errorf("String:\"foo\" and rune:`foo` should be Convertible and same")
	}
}

func TestObjectsAreSameType(t *testing.T) {
	tts := []struct {
		name   string
		a      any
		b      any
		expect bool
	}{
		{name: "Same", a: "a", b: "b", expect: true},
		{name: "Number", a: int(7), b: float32(7.0), expect: false},
		{name: "map", a: map[string]int{}, b: map[string]string{}, expect: false},
	}
	for _, tt := range tts {
		if got := objectsAreSameType(tt.a, tt.b); got != tt.expect {
			t.Errorf("objectsAreSameType is wrong. Test Name:%v args:%#v %#v, expect:%v, got:%v", tt.name, tt.a, tt.b, tt.expect, got)
		}
	}
}

func TestObjectsAreSame(t *testing.T) {
	tts := []struct {
		a      any
		b      any
		expect bool
	}{
		{a: "aiko", b: "aiko", expect: true},
		{a: 123, b: 123, expect: true},
		{a: 123.5, b: 123.5, expect: true},
		{a: []byte("Hello World"), b: []byte("Hello World"), expect: true},
		{a: nil, b: nil, expect: true},

		// cases that are expected not to be equal
		{a: map[int]int{5: 10}, b: map[int]int{10: 20}, expect: false},
		{a: 'x', b: "x", expect: false},
		{a: "x", b: 'x', expect: false},
		{a: 0, b: 0.1, expect: false},
		{a: 0.1, b: 0, expect: false},
		{a: time.Now, b: time.Now, expect: false},
		{a: func() {}, b: func() {}, expect: false},
		{a: uint32(10), b: int32(10), expect: false},
	}
	for _, tt := range tts {
		if got := objectsAreSame(tt.a, tt.b); got != tt.expect {
			t.Errorf("objectsAreSame is wrong. args:%#v %#v, expect:%v, got:%v", tt.a, tt.b, tt.expect, got)
		}
	}
}
