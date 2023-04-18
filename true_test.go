package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestTrue(t *testing.T) {
	// pass
	actually.Got(1 == 1).True(t)         //lint:ignore SA4000 this is test
	actually.Got("foo" == "foo").True(t) //lint:ignore SA4000 this is test
	actually.Got(1 == 2).False(t)
	actually.Got("foo" == "hoo").False(t)

	// fail now
	//actually.Got(1==2).FailNow().True(t)

	// fail
	// actually.Got(12).True(t)
	// actually.Got(1==2).True(t)
	// actually.Got(12).False(t)
	// actually.Got(1==1).False(t)
	// actually.Got("foo"=="foo").False(t)
}
