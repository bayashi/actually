package actually

import (
	"testing"
)

// Same method verifies that two objects are same in value and type.
// Function type value is not acceptable. And not verify pointer address.
// It will be fail, int(1) and uint(1), because of type.
/*
	Pass: actually.Got(12).Expect(12).Same(t)
	Fail: actually.Got(int16(12)).Expect(int32(12)).Same(t) // not same type
*/
func (a *TestingA) Same(t *testing.T, testNames ...string) *TestingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !objectsAreSameType(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_WrongType))
	}

	if isFuncType(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_GotIsFunc).Notice(notice_Same_NotAcceptable))
	}
	if isFuncType(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_ExpectIsFunc).Notice(notice_Same_NotAcceptable))
	}

	if !objectsAreSame(expect, got) {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(reason_NotSame))
	}

	return a
}

// SamePointer method verifies that two objects point to the same object.
func (a *TestingA) SamePointer(t *testing.T, testNames ...string) *TestingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isPointerType(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_GotIsNotPointer).Notice(notice_SamePointer_ShouldPointer))
	}
	if !isPointerType(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotPointer).Notice(notice_SamePointer_ShouldPointer))
	}

	if !objectsAreSameType(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_WrongType))
	}

	if got != expect {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(reason_WrongPointerAddress))
	}

	return a
}

// SameNumber method verifies that each pair of numbers are same or
// convertible to the same types and converted objects are equal. (i.e. int* and float*)
/*
	Pass: actually.Got(float32(1.0)).Expect(int64(1)).SameNumber(t)
	Fail: actually.Got("1").Expect(1).SameNumber(t) // string cannot convert to int
	      actually.Got(nil).Expect(0).SameNumber(t) // <nil> is not acceptable
*/
func (a *TestingA) SameNumber(t *testing.T, testNames ...string) *TestingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isFuncType(got) && !isFuncType(expect) && objectsAreSame(expect, got) {
		return a // Pass
	}

	if isTypeNil(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_GotIsNilType).Notice(notice_SameNumber_ShouldNumber))
	}
	if isTypeNil(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNilType).Notice(notice_SameNumber_ShouldNumber))
	}

	if !isValidValue(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotValidValue))
	}

	if !objectsAreConvertible(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(reason_NotConvertibleTypes))
	}

	if !isSameConvertedValueAsOther(expect, got) {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(reason_NotSame))
	}

	return a
}
