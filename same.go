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

	got := a.got
	expect := a.expect

	if !objectsAreSameType(expect, got) {
		return a.fail(reportForSame(a), reason_WrongType)
	}

	if isFuncType(got) {
		w := reportForSame(a).Message("Notice", notice_Same_NotAcceptable)
		return a.fail(w, reason_GotIsFunc)
	}
	if isFuncType(expect) {
		w := reportForSame(a).Message("Notice", notice_Same_NotAcceptable)
		return a.fail(w, reason_ExpectIsFunc)
	}

	if !objectsAreSame(expect, got) {
		return a.fail(reportForSameWithDiff(a), reason_NotSame)
	}

	return a
}

// SamePointer method verifies that two objects point to the same object.
func (a *TestingA) SamePointer(t *testing.T, testNames ...string) *TestingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	got := a.got
	expect := a.expect

	if !isPointerType(got) {
		w := reportForSame(a).Message("Notice", notice_SamePointer_ShouldPointer)
		return a.fail(w, reason_GotIsNotPointer)
	}
	if !isPointerType(expect) {
		w := reportForSame(a).Message("Notice", notice_SamePointer_ShouldPointer)
		return a.fail(w, reason_ExpectIsNotPointer)
	}

	if !objectsAreSameType(expect, got) {
		return a.fail(reportForSame(a), reason_WrongType)
	}

	if got != expect {
		return a.fail(reportForSameWithDiff(a), reason_WrongPointerAddress)
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

	if !isSameConvertedValueAsOther(expect, got) {
		return a.fail(reportForSameWithDiff(a), reason_NotSame)
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

	if !objectsAreSameType(a.expect, a.got) {
		w := reportForSame(a).Expect(a.expect).Got(a.got).
			Message("Notice", "SameType() just verifies the type. It doesn't care about the actual value")
		return a.fail(w, reason_WrongType)
	}

	return a
}
