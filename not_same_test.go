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
	stubConfirm(t, func() {
		Got("foo").Expect("foo").NotSame(t)
	}, reason_Same)

	stubConfirm(t, func() {
		Got([]int{1, 2, 3}).Expect([]int{1, 2, 3}).NotSame(t)
	}, reason_Same)

	i := &[]int{1, 2, 3}
	j := &[]int{1, 2, 3}
	Got(i).Expect(j).NotSamePointer(t) // pass
	stubConfirm(t, func() {
		Got(i).Expect(j).NotSame(t) // Not same pointer address, but same values. So, expected fail. These are same.
	}, reason_Same)

	f := func() {}
	stubConfirm(t, func() {
		Got("").Expect(f).NotSame(t)
	}, reason_ExpectIsFunc)
	stubConfirm(t, func() {
		Got(f).Expect("").NotSame(t)
	}, reason_GotIsFunc)
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
	stubConfirm(t, func() {
		Got(1).Expect(1).NotSameNumber(t)
	}, reason_Same)

	stubConfirm(t, func() {
		Got(0).Expect(nil).NotSameNumber(t)
	}, reason_ExpectIsNilType)
	stubConfirm(t, func() {
		Got(nil).Expect(0).NotSameNumber(t)
	}, reason_GotIsNilType)

	stubConfirm(t, func() {
		Got(0).Expect(nil).NotSameNumber(t)
	}, reason_ExpectIsNilType)

	stubConfirm(t, func() {
		Got(0).Expect([]byte("0")).NotSameNumber(t)
	}, reason_ExpectIsNotNumber)
	stubConfirm(t, func() {
		Got("1").Expect(1).NotSameNumber(t)
	}, reason_GotIsNotNumber)

	stubConfirm(t, func() {
		// NOTE: Be aware of a result of test to compare int value with float value
		Got(1).Expect(float64(1.0000000000000001)).NotSameNumber(t)
	}, reason_Same)
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
