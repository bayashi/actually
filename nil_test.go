package actually

import (
	"testing"
	"unsafe"
)

func TestNil(t *testing.T) {
	Got(nil).Nil(t)
	Got("").NotNil(t)

	var n unsafe.Pointer = nil
	Got(n).Nil(t)

	// test name
	Got(nil).Nil(t, "nil test")
	Got(1).NotNil(t, "not nil test")
	Got(nil).Name("nil test").Nil(t)
	Got(nil).Name("Not nil").Nil(t, "Not nil", "Not nil")
}

func TestNil_Fail(t *testing.T) {
	stub()
	Got("").Nil(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectNilButNotNil {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectNilButNotNil, stubRes)
	}

	stub()
	Got(nil).NotNil(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNotNil {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectNilButNotNil, stubRes)
	}
}
