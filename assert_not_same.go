package actually

import (
	"testing"
)

// NotSame method verifies that two objects are not same value.
// NOTE1: Function type value is not acceptable.
// NOTE2: Not verify pointer address, just verify only values.
// NOTE3: If you would like to verify pointers, types or numbers, then please use NotSame* method instead for more accuracy.
/*
	// Pass
	actually.Got(12).Expect(34).NotSame(t)
	actually.Got("bar").Expect("baz").NotSame(t)
	// Fail
	actually.Got("foo").Expect("foo").NotSame(t)
*/
func (a *testingA) NotSame(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got
	expect := a.expect

	if isFuncType(got) {
		w := reportForSame(a).Message(notice_Label, notice_Same_NotAcceptable)
		return a.fail(w, reason_GotIsFunc)
	}
	if isFuncType(expect) {
		w := reportForSame(a).Message(notice_Label, notice_Same_NotAcceptable)
		return a.fail(w, reason_ExpectIsFunc)
	}

	if objectsAreSame(expect, got) {
		return a.fail(reportForSameWithDiff(a), reason_Same)
	}

	return a
}

// NotSamePointer method verifies that two objets are not same pointer.
func (a *testingA) NotSamePointer(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got
	expect := a.expect

	if !isPointerType(got) {
		w := reportForSame(a).Message(notice_Label, notice_NotSamePointer_ShouldPointer)
		return a.fail(w, reason_GotIsNotPointer)
	}
	if !isPointerType(expect) {
		w := reportForSame(a).Message(notice_Label, notice_NotSamePointer_ShouldPointer)
		return a.fail(w, reason_ExpectIsNotPointer)
	}

	if got == expect {
		return a.fail(reportForSame(a), reason_SamePointerAddress)
	}

	return a
}

// Deprecated: Use `NotSameConvertibleNumber` instead. The `NotSameNumber` method will be removed.
func (a *testingA) NotSameNumber(t *testing.T, testNames ...string) *testingA {
	return a.NotSameConvertibleNumber(t, testNames...)
}

// NotSameConvertibleNumber method verifies that two objects are not same number.
func (a *testingA) NotSameConvertibleNumber(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got
	expect := a.expect

	if isTypeNil(got) {
		w := reportForSame(a).Message(notice_Label, notice_SameConvertibleNumber_ShouldNumber)
		return a.fail(w, reason_GotIsNilType)
	}
	if isTypeNil(expect) {
		w := reportForSame(a).Message(notice_Label, notice_SameConvertibleNumber_ShouldNumber)
		return a.fail(w, reason_ExpectIsNilType)
	}

	if !isTypeNumber(got) {
		w := reportForSame(a).Message(notice_Label, notice_SameConvertibleNumber_ShouldNumber)
		return a.fail(w, reason_GotIsNotNumber)
	}
	if !isTypeNumber(expect) {
		w := reportForSame(a).Message(notice_Label, notice_SameConvertibleNumber_ShouldNumber)
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
func (a *testingA) NotSameType(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !objectsAreSameType(a.expect, a.got) {
		return a
	}

	return a.fail(reportForSame(a), reason_SameType)
}
