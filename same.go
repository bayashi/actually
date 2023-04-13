package actually

import (
	"reflect"
	"testing"
)

// `Same` verifies that two objects are same in value and type.
// Function type value is not acceptable. And not verify pointer address.
// It will be fail, int(1) and uint(1), because of type.
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
