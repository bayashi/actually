package actually

import (
	"fmt"
	"testing"
)

type testCase struct {
	expected           any
	actuallyGot        any
	expectedFailReason string
}

type testName = string

func TestSame(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"nil": {
			expected: nil, actuallyGot: nil,
		},
		"blank string": {
			expected: "", actuallyGot: "",
		},
		"string": {
			expected: "a", actuallyGot: "a",
		},
		"zero": {
			expected: 0, actuallyGot: 0,
		},
		"int8": {
			expected: int8(12), actuallyGot: int8(12),
		},
		"array": {
			expected: [2]int{1, 2}, actuallyGot: [2]int{1, 2},
		},
		"blank slice": {
			expected: []string{}, actuallyGot: []string{},
		},
		"slice": {
			expected: []string{"a"}, actuallyGot: []string{"a"},
		},
		"map": {
			expected: map[string]int{"foo": 12}, actuallyGot: map[string]int{"foo": 12},
		},
		"struct": {
			expected: struct{ bar string }{bar: "foo"}, actuallyGot: struct{ bar string }{bar: "foo"},
		},
	} {
		t.Run(tn, func(t *testing.T) {
			Got(tt.actuallyGot).Expect(tt.expected).Same(t)
		})
	}

	// `foo` and `bar` are same value. But these pointer addresses are different in fact.
	foo := struct {
		bar string
	}{
		bar: "foo",
	}
	bar := struct {
		bar string
	}{
		bar: "foo",
	}
	Got(foo).Expect(bar).Same(t)                                     // Pass
	Got(fmt.Sprintf("%p", &foo) == fmt.Sprintf("%p", &bar)).False(t) // Different

	// test name
	Got(0).Expect(0).Same(t, "zero")

	// NOT SUPPORT chan YET
	// ch1 := make(chan string, 1)
	// ch1 <- "foo"
	// ch2 := make(chan string, 1)
	// ch2 <- "foo"
	// actually.Got(ch1).Expect(ch2).Same(t)
}

func TestSame_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"different strings": {
			expected: "a", actuallyGot: "b", expectedFailReason: reason_NotSame,
		},
		"different type": {
			expected: int16(12), actuallyGot: int32(12), expectedFailReason: reason_WrongType,
		},
		"func type is not supported": {
			expected: func() {}, actuallyGot: func() {}, expectedFailReason: reason_GotIsFunc,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).Same(t)
			}, tt.expectedFailReason)
		})
	}
}

func TestSamePointer(t *testing.T) {
	i := 7
	ptr := &i
	ptr2 := ptr
	Got(ptr).Expect(ptr2).SamePointer(t)
	Got(ptr).Expect(&i).SamePointer(t)

	// test name
	Got(ptr).Expect(ptr2).SamePointer(t, "Same Pointer")

	// fail
	// actually.Got("").Expect(ptr).SamePointer(t)
	// actually.Got(ptr).Expect("").SamePointer(t)
	// j := 7
	// actually.Got(ptr).Expect(&j).SamePointer(t)
}

func TestSamePointer_Fail(t *testing.T) {
	i := 7
	ptr := &i
	stubConfirm(t, func() {
		Got(ptr).Expect(1).SamePointer(t)
	}, reason_ExpectIsNotPointer)
	stubConfirm(t, func() {
		Got(1).Expect(ptr).SamePointer(t)
	}, reason_GotIsNotPointer)

	j := 7
	stubConfirm(t, func() {
		Got(ptr).Expect(&j).SamePointer(t)
	}, reason_WrongPointerAddress)
}

func TestSameConvertibleNumber(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"same number": {
			expected: 1, actuallyGot: 1,
		},
		"converted same": {
			expected: int8(12), actuallyGot: int64(12),
		},
		"float and int are convertible": {
			expected: float32(1.0), actuallyGot: int64(1),
		},
		// "complex number": {
		// 	expected: complex64(1e+10 + 1e+10i), actuallyGot: complex128(1e+10 + 1e+10i),
		// },
	} {
		t.Run(tn, func(t *testing.T) {
			Got(tt.actuallyGot).Expect(tt.expected).SameConvertibleNumber(t, tn)
		})
	}

	// test name
	Got(1).Expect(1).SameConvertibleNumber(t, "Same Number")
}

func TestSameConvertibleNumber_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"expected is not number": {
			expected: "1", actuallyGot: 1, expectedFailReason: reason_ExpectIsNotNumber,
		},
		"got is not number": {
			expected: 1, actuallyGot: "1", expectedFailReason: reason_GotIsNotNumber,
		},
		"expected is nil": {
			expected: nil, actuallyGot: 1, expectedFailReason: reason_ExpectIsNilType,
		},
		"got is nil": {
			expected: 1, actuallyGot: nil, expectedFailReason: reason_GotIsNilType,
		},
		"should not be same": {
			expected: int8(14), actuallyGot: int(270), expectedFailReason: reason_NotSame,
		},
		"should not be same, too": {
			expected: int(270), actuallyGot: int8(14), expectedFailReason: reason_NotSame,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).SameConvertibleNumber(t, tn)
			}, tt.expectedFailReason)
		})
	}
}

func TestChain(t *testing.T) {
	Got(7).NotNil(t).
		Expect(7).SameConvertibleNumber(t).Same(t)
}

func TestSameType(t *testing.T) {
	Got(nil).Expect(nil).SameType(t)
	Got(true).Expect(false).SameType(t) // both are boolean
	a := Got(t).Expect(t).SameType(t)
	Got(a).Expect(a).SameType(t)
}

func TestSameType_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"string - int": {
			expected: "1", actuallyGot: 1, expectedFailReason: reason_WrongType,
		},
		"nil - int": {
			expected: nil, actuallyGot: 0, expectedFailReason: reason_WrongType,
		},
		"int32 - int64": {
			expected: int32(5), actuallyGot: int64(5), expectedFailReason: reason_WrongType,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).SameType(t, tn)
			}, tt.expectedFailReason)
		})
	}
}
