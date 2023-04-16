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

	if reflect.TypeOf(a.got.RawValue()) != reflect.TypeOf(a.expect.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_WrongType))
	}

	if isFuncType(a.got.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsFunc))
	}
	if isFuncType(a.expect.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsFunc))
	}

	if !objectsAreSame(a.expect.RawValue(), a.got.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_NotSame))
	}

	return a
}

// `SamePointer` verifies that two objects point to the same object.
func (a *testingA) SamePointer(t *testing.T) *testingA {
	a.t = t

	if !isPointerType(a.got.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsNotPointer))
	}
	if !isPointerType(a.expect.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNotPointer))
	}

	if reflect.TypeOf(a.got.RawValue()) != reflect.TypeOf(a.expect.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_WrongType))
	}

	if a.got.RawValue() != a.expect.RawValue() {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_WrongPointerAddress))
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

	if !isFuncType(a.got.RawValue()) && !isFuncType(a.expect.RawValue()) && objectsAreSame(a.expect.RawValue(), a.got.RawValue()) {
		return a // Pass
	}

	gotType := reflect.TypeOf(a.got.RawValue())
	if gotType == nil {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_GotIsNilType))
	}
	expectType := reflect.TypeOf(a.expect.RawValue())
	if expectType == nil {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNilType))
	}

	expectValue := reflect.ValueOf(a.expect.RawValue())
	if !expectValue.IsValid() {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_ExpectIsNotValidValue))
	}

	if !reflect.ValueOf(a.got.RawValue()).Type().ConvertibleTo(expectType) ||
		!expectValue.Type().ConvertibleTo(gotType) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_NotConvertibleTypes))
	}

	if !reflect.DeepEqual(expectValue.Convert(gotType).Interface(), a.got.RawValue()) {
		a.t.Helper()
		return a.fail(reportForSame(a).Reason(failReason_NotSame))
	}

	return a
}
