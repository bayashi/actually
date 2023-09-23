package actually

import (
	"testing"
)

// NotSamePointer method verifies that two objets are not same pointer.
func (a *TestingA) NotSamePointer(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got.RawValue()
	expect := a.expect.RawValue()

	if !isPointerType(got) {
		return a.fail(reportForSame(a).Reason(reason_GotIsNotPointer).Notice(notice_NotSamePointer_ShouldPointer))
	}
	if !isPointerType(expect) {
		return a.fail(reportForSame(a).Reason(reason_ExpectIsNotPointer).Notice(notice_NotSamePointer_ShouldPointer))
	}

	if got == expect {
		return a.fail(reportForNotSameType(a).Reason(reason_SamePointerAddress))
	}

	return a
}

// NotSameNumber method verifies that two objects are not same number.
func (a *TestingA) NotSameNumber(t *testing.T, testNames ...string) *TestingA {
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

	if convert2float64(expect) == convert2float64(got) {
		return a.fail(reportForSameWithDiff(a).Reason(reason_Same))
	}

	return a
}

// NotSameType method verifies that each pair of values are NOT same type.
// Not care about actual value of these. Just verify the type.
func (a *TestingA) NotSameType(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !objectsAreSameType(a.expect.RawValue(), a.got.RawValue()) {
		return a
	}

	return a.fail(reportForNotSameType(a))
}
