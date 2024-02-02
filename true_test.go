package actually

import (
	"testing"
)

func TestTrue(t *testing.T) {
	// pass
	Got(1 == 1).True(t)         //lint:ignore SA4000 this is test
	Got("foo" == "foo").True(t) //lint:ignore SA4000 this is test
	Got(1 == 2).False(t)
	Got("foo" == "hoo").False(t)

	// test name
	Got(2 == 2).True(t, "True test") //lint:ignore SA4000 this is test
	Got(2 == 3).False(t, "False test")

	// fail now
	//actually.Got(1==2).FailNow().True(t)

	// fail
	// actually.Got(12).True(t)
	// actually.Got(1==2).True(t)
	// actually.Got(12).False(t)
	// actually.Got(1==1).False(t)
	// actually.Got("foo"=="foo").False(t)
}

func TestTrue_Fail(t *testing.T) {
	stub()
	Got(false).True(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != message_ExpectTrue {
		t.Errorf("expected `%s`, but got `%s`", message_ExpectTrue, stubRes)
	}
	stub()
	Got(12).True(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongType {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongType, stubRes)
	}

	stub()
	Got(true).False(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != message_ExpectFalse {
		t.Errorf("expected `%s`, but got `%s`", message_ExpectFalse, stubRes)
	}
	stub()
	Got(12).False(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongType {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongType, stubRes)
	}
}
