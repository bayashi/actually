package actually

import (
	"reflect"
	"testing"
)

// `Same` verifies that two objects are same in value and type.
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
