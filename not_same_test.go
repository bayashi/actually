package actually

import (
	"testing"
)

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

	stub()
	Got("").Expect(ptr).NotSamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNotPointer {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNotPointer, stubRes)
	}

	stub()
	Got(ptr).Expect("").NotSamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNotPointer {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNotPointer, stubRes)
	}

	stub()
	Got(ptr).Expect(&i).NotSamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_SamePointerAddress {
		t.Errorf("expected `%s`, but got `%s`", reason_SamePointerAddress, stubRes)
	}
}

func TestNotSameNumber(t *testing.T) {
	Got(1).Expect(2).NotSameNumber(t, "these should be different number")
	Got(int8(1)).Expect(int32(2)).NotSameNumber(t)
	Got(float32(1.1)).Expect(int64(1)).NotSameNumber(t)

	// NOTE: Be aware of a result of test to compare int value with float value
	Got(1).Expect(float64(1.000000000000001)).NotSameNumber(t)
}

func TestNotSameNumber_Fail(t *testing.T) {
	stub()
	Got(1).Expect(1).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_Same {
		t.Errorf("expected `%s`, but got `%s`", reason_Same, stubRes)
	}

	stub()
	Got(nil).Expect(0).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNilType {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNilType, stubRes)
	}
	stub()
	Got(0).Expect(nil).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNilType {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNilType, stubRes)
	}

	stub()
	Got("1").Expect(1).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNotNumber {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNotNumber, stubRes)
	}
	stub()
	Got(0).Expect([]byte("0")).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNotNumber {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNotNumber, stubRes)
	}

	// NOTE: Be aware of a result of test to compare int value with float value
	stub()
	Got(1).Expect(float64(1.0000000000000001)).NotSameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_Same {
		t.Errorf("expected `%s`, but got `%s`", reason_Same, stubRes)
	}
}

func TestNotSameType(t *testing.T) {
	Got(nil).Expect(0).NotSameType(t)
	Got("1").Expect(1).NotSameType(t)
	Got(t).Expect(&testing.B{}).NotSameType(t)
}

func TestNotSameType_Fail(t *testing.T) {
	stub()
	Got(nil).Expect(nil).NotSameType(t) // Both are nil, it will be failed. regarded same.
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_SameType {
		t.Errorf("expected `%s`, but got `%s`", reason_SameType, stubRes)
	}

	stub()
	Got(true).Expect(false).NotSameType(t) // Both are boolean.
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_SameType {
		t.Errorf("expected `%s`, but got `%s`", reason_SameType, stubRes)
	}

	stub()
	Got(&testing.T{}).Expect(t).NotSameType(t) // Both are same type.
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_SameType {
		t.Errorf("expected `%s`, but got `%s`", reason_SameType, stubRes)
	}
}
