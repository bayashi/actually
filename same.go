package actually

import (
	"reflect"
	"testing"
)

// `Same` verifies that two objects are same in value and type.
// Function type value is not acceptable. And not verify pointer address.
// It will be fail, int(1) and uint(1), because of type.
/*
	Pass: actually.Got(12).Expect(12).Same(t)
	Fail: actually.Got(int16(12)).Expect(int32(12)).Same(t)
*/
func (a *testingA) Same(t *testing.T) *testingA {
	a.t = t

	if reflect.TypeOf(a.got) != reflect.TypeOf(a.expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_WrongType))
	}

	if isFuncType(a.got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_GotIsFunc))
	}
	if isFuncType(a.expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_ExpectIsFunc))
	}

	if !objectsAreSame(a.expect, a.got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_NotSame))
	}

	return a
}

// `SamePointer` verifies that two objects point to the same object.
func (a *testingA) SamePointer(t *testing.T) *testingA {
	a.t = t

	if !isPointerType(a.got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_GotIsNotPointer))
	}
	if !isPointerType(a.expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_ExpectIsNotPointer))
	}

	if reflect.TypeOf(a.got) != reflect.TypeOf(a.expect) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_WrongType))
	}

	if a.got != a.expect {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_WrongPointerAddress))
	}

	return a
}

// `SameNumber` verifies that each pair of numbers are same or
// convertible to the same types and converted objects are equal. (i.e. int* and float*)
/*
	Pass: actually.Got(float32(1.0)).Expect(int64(1)).SameNumber(t)
	Fail: actually.Got("1").Expect(1).SameNumber(t) // string cannot convert to int
*/
func (a *testingA) SameNumber(t *testing.T) *testingA {
	a.t = t

	if !isFuncType(a.got) && !isFuncType(a.expect) && objectsAreSame(a.expect, a.got) {
		return a // Pass
	}

	gotType := reflect.TypeOf(a.got)
	if gotType == nil {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_GotIsNilType))
	}
	expectType := reflect.TypeOf(a.expect)
	if expectType == nil {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_ExpectIsNilType))
	}

	expectValue := reflect.ValueOf(a.expect)
	if !expectValue.IsValid() {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_ExpectIsNotValidValue))
	}

	if !reflect.ValueOf(a.got).Type().ConvertibleTo(expectType) ||
		!expectValue.Type().ConvertibleTo(gotType) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_NotConvertibleTypes))
	}

	if !reflect.DeepEqual(expectValue.Convert(gotType).Interface(), a.got) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(FailReason_NotSame))
	}

	return a
}
