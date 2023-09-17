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
