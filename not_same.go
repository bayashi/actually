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

	got := a.got
	expect := a.expect

	if !isPointerType(got) {
		w := reportForSame(a).Message("Notice", notice_NotSamePointer_ShouldPointer)
		return a.fail(w, reason_GotIsNotPointer)
	}
	if !isPointerType(expect) {
		w := reportForSame(a).Message("Notice", notice_NotSamePointer_ShouldPointer)
		return a.fail(w, reason_ExpectIsNotPointer)
	}

	if got == expect {
		return a.fail(reportForSame(a), reason_SamePointerAddress)
	}

	return a
}

// NotSameNumber method verifies that two objects are not same number.
func (a *TestingA) NotSameNumber(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got
	expect := a.expect

	if isTypeNil(got) {
		w := reportForSame(a).Message("Notice", notice_SameNumber_ShouldNumber)
		return a.fail(w, reason_GotIsNilType)
	}
	if isTypeNil(expect) {
		w := reportForSame(a).Message("Notice", notice_SameNumber_ShouldNumber)
		return a.fail(w, reason_ExpectIsNilType)
	}

	if !isTypeNumber(got) {
		w := reportForSame(a).Message("Notice", notice_SameNumber_ShouldNumber)
		return a.fail(w, reason_GotIsNotNumber)
	}
	if !isTypeNumber(expect) {
		w := reportForSame(a).Message("Notice", notice_SameNumber_ShouldNumber)
		return a.fail(w, reason_ExpectIsNotNumber)
	}

	if !isValidValue(expect) {
		return a.fail(reportForSame(a), reason_ExpectIsNotValidValue)
	}

	if !objectsAreConvertible(expect, got) {
		return a.fail(reportForSame(a), reason_NotConvertibleTypes)
	}

	if convert2float64(expect) == convert2float64(got) {
		return a.fail(reportForSameWithDiff(a), reason_Same)
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

	if !objectsAreSameType(a.expect, a.got) {
		return a
	}

	return a.fail(reportForSame(a), reason_SameType)
}
