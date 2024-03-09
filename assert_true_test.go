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
	stubConfirm(t, func() {
		Got(false).True(t)
	}, message_ExpectTrue)

	stubConfirm(t, func() {
		Got(12).True(t)
	}, reason_WrongType)

	stubConfirm(t, func() {
		Got(true).False(t)
	}, message_ExpectFalse)

	stubConfirm(t, func() {
		Got(12).False(t)
	}, reason_WrongType)
}
