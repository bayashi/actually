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
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !objectsAreSameType(expect, got) {
		return a.fail(reportForSame(a).Reason(reason_WrongType))
	}

	if isFuncType(got) {
		return a.fail(reportForSame(a).Reason(reason_GotIsFunc).Notice(notice_Same_NotAcceptable))
	}
	if isFuncType(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsFunc).Notice(notice_Same_NotAcceptable))
	}

	if !objectsAreSame(expect, got) {
		return a.fail(reportForSameWithDiff(a).Reason(reason_NotSame))
	}

	return a
}

// SamePointer method verifies that two objects point to the same object.
func (a *TestingA) SamePointer(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isPointerType(got) {
		return a.fail(reportForSame(a).Reason(reason_GotIsNotPointer).Notice(notice_SamePointer_ShouldPointer))
	}
	if !isPointerType(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotPointer).Notice(notice_SamePointer_ShouldPointer))
	}

	if !objectsAreSameType(expect, got) {
		return a.fail(reportForSame(a).Reason(reason_WrongType))
	}

	if got != expect {
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
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if isTypeNil(got) {
		return a.fail(reportForSame(a).Reason(reason_GotIsNilType).Notice(notice_SameNumber_ShouldNumber))
	}
	if isTypeNil(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNilType).Notice(notice_SameNumber_ShouldNumber))
	}

	if !isTypeNumber(got) {
		return a.fail(reportForSame(a).Reason(reason_GotIsNotNumber).Notice(notice_SameNumber_ShouldNumber))
	}
	if !isTypeNumber(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotNumber).Notice(notice_SameNumber_ShouldNumber))
	}

	if !isValidValue(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotValidValue))
	}

	if !objectsAreConvertible(expect, got) {
		return a.fail(reportForSame(a).Reason(reason_NotConvertibleTypes))
	}

	if !isSameConvertedValueAsOther(expect, got) {
		return a.fail(reportForSameWithDiff(a).Reason(reason_NotSame))
	}

	return a
}

// SameType method verifies that each pair of values are same type or not.
// Not care about actual value of these. Just verify the type.
/*
	Pass: actually.Got("foo").Expect("bar").SameType(t)
	Fail: actually.Got("1").Expect(1).SameType(t)
	      actually.Got(1).Expect(1.0).SameType(t)
*/
func (a *TestingA) SameType(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !objectsAreSameType(a.expect.RawValue(), a.got.RawValue()) {
		return a.fail(reportForSameType(a))
	}

	return a
}
