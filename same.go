package actually

import (
	"testing"
)

// `Same` verifies that two objects are same in value and type.
// Function type value is not acceptable. And not verify pointer address.
// It will be fail, int(1) and uint(1), because of type.
/*
	Pass: actually.Got(12).Expect(12).Same(t)
	Fail: actually.Got(int16(12)).Expect(int32(12)).Same(t)
*/
func (a *testingA) Same(t *testing.T, testNames ...string) *testingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !objectsAreSameType(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_WrongType))
	}

	if isFuncType(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsFunc).Notice(failNotice_NotAcceptableSameMethod))
	}
	if isFuncType(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsFunc).Notice(failNotice_NotAcceptableSameMethod))
	}

	if !objectsAreSame(expect, got) {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(failReason_NotSame))
	}

	return a
}

// `SamePointer` verifies that two objects point to the same object.
func (a *testingA) SamePointer(t *testing.T, testNames ...string) *testingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isPointerType(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsNotPointer).Notice(failNotice_ShouldPointerSamePointer))
	}
	if !isPointerType(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNotPointer).Notice(failNotice_ShouldPointerSamePointer))
	}

	if !objectsAreSameType(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_WrongType))
	}

	if got != expect {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(failReason_WrongPointerAddress))
	}

	return a
}

// `SameNumber` verifies that each pair of numbers are same or
// convertible to the same types and converted objects are equal. (i.e. int* and float*)
/*
	Pass: actually.Got(float32(1.0)).Expect(int64(1)).SameNumber(t)
	Fail: actually.Got("1").Expect(1).SameNumber(t) // string cannot convert to int
	      actually.Got(nil).Expect(0).SameNumber(t) // <nil> is not acceptable
*/
func (a *testingA) SameNumber(t *testing.T, testNames ...string) *testingA {
	a.name = a.naming(testNames...)
	a.t = t

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isFuncType(got) && !isFuncType(expect) && objectsAreSame(expect, got) {
		return a // Pass
	}

	if isTypeNil(got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsNilType).Notice(failNotice_ShouldNumberSameNumber))
	}
	if isTypeNil(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNilType).Notice(failNotice_ShouldNumberSameNumber))
	}

	if !isValidValue(expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNotValidValue))
	}

	if !objectsAreConvertible(expect, got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_NotConvertibleTypes))
	}

	if !isSameConvertedValueAsOther(expect, got) {
		a.t.Helper()
		return a.fail(reportForSameWithDiff(a).Reason(failReason_NotSame))
	}

	return a
}
