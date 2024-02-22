package actually

import (
	"testing"
)

func TestNotSame(t *testing.T) {
	Got("bar").Expect("baz").NotSame(t)
	Got([]int{1, 2, 3}).Expect([]int{1, 2, 3, 4}).NotSame(t)

	// test name
	Got("aiko").Expect("eiko").NotSame(t, "Not Same")
}

func TestNotSame_Fail(t *testing.T) {
	i := &[]int{1, 2, 3}
	j := &[]int{1, 2, 3}

	Got(i).Expect(j).NotSamePointer(t) // pass. pointer addresses are different

	for tn, tt := range map[testName]testCase{
		"same string": {
			expected: "foo", actuallyGot: "foo", expectedFailReason: reason_Same,
		},
		"same slice": {
			expected: []int{1, 2, 3}, actuallyGot: []int{1, 2, 3}, expectedFailReason: reason_Same,
		},
		"Not same pointer address, but same values. These are same. So, expected to fail": {
			expected: j, actuallyGot: i, expectedFailReason: reason_Same,
		},
		"expect value is func": {
			expected: func() {}, actuallyGot: "", expectedFailReason: reason_ExpectIsFunc,
		},
		"got value is func": {
			expected: "", actuallyGot: func() {}, expectedFailReason: reason_GotIsFunc,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).NotSame(t)
			}, tt.expectedFailReason)
		})
	}
}

func TestNotSamePointer(t *testing.T) {
	i := 7
	j := 7
	ptr := &i
	ptr2 := &j
	Got(ptr).Expect(ptr2).NotSamePointer(t)
	Got(ptr).Expect(&j).NotSamePointer(t)
	Got(ptr2).Expect(&i).NotSamePointer(t)

	// test name
	Got(&i).Expect(&j).NotSamePointer(t, "Not Same Pointer")
}

func TestNotSamePointer_Fail(t *testing.T) {
	i := 7
	ptr := &i

	stubConfirm(t, func() {
		Got(ptr).Expect("").NotSamePointer(t)
	}, reason_ExpectIsNotPointer)
	stubConfirm(t, func() {
		Got("").Expect(ptr).NotSamePointer(t)
	}, reason_GotIsNotPointer)

	stubConfirm(t, func() {
		Got(ptr).Expect(&i).NotSamePointer(t)
	}, reason_SamePointerAddress)
}

func TestNotSameNumber(t *testing.T) {
	Got(1).Expect(2).NotSameNumber(t, "these should be different number")
	Got(int8(1)).Expect(int32(2)).NotSameNumber(t)
	Got(float32(1.1)).Expect(int64(1)).NotSameNumber(t)

	// NOTE: Be aware of a result of test to compare int value with float value
	Got(1).Expect(float64(1.000000000000001)).NotSameNumber(t)
}

func TestNotSameNumber_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"same number": {
			expected: 1, actuallyGot: 1, expectedFailReason: reason_Same,
		},
		"expect value is nil": {
			expected: nil, actuallyGot: 0, expectedFailReason: reason_ExpectIsNilType,
		},
		"got value is nil": {
			expected: 0, actuallyGot: nil, expectedFailReason: reason_GotIsNilType,
		},
		"expect value is not a kind of number": {
			expected: []byte("0"), actuallyGot: 0, expectedFailReason: reason_ExpectIsNotNumber,
		},
		"got value is not a kind of number": {
			expected: 1, actuallyGot: "1", expectedFailReason: reason_GotIsNotNumber,
		},
		"NOTE: Be aware of a result of test to compare int value with float value": {
			expected: float64(1.0000000000000001), actuallyGot: 1, expectedFailReason: reason_Same,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).NotSameNumber(t)
			}, tt.expectedFailReason)
		})
	}
}

func TestNotSameType(t *testing.T) {
	Got(nil).Expect(0).NotSameType(t)
	Got("1").Expect(1).NotSameType(t)
	Got(t).Expect(&testing.B{}).NotSameType(t)
}

func TestNotSameType_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got(nil).Expect(nil).NotSameType(t) // Both are nil, it will be failed. regarded same.
	}, reason_SameType)

	stubConfirm(t, func() {
		Got(true).Expect(false).NotSameType(t) // Both are boolean.
	}, reason_SameType)

	stubConfirm(t, func() {
		Got(&testing.T{}).Expect(t).NotSameType(t) // Both are same type.
	}, reason_SameType)
}
