package actually_test

import (
	"testing"
	"unsafe"

	"github.com/bayashi/actually"
)

func TestNil(t *testing.T) {
	// pass
	actually.Got(nil).Nil(t)
	actually.Got("").NotNil(t)

	var n unsafe.Pointer = nil
	actually.Got(n).Nil(t)

	// test name
	actually.Got(nil).Nil(t, "nil test")
	actually.Got(1).NotNil(t, "not nil test")
	actually.Got(nil).Name("nil test").Nil(t)
	actually.Got(nil).Name("Not nil").Nil(t, "Not nil", "Not nil")

	// fail now
	//actually.Got("").FailNow().Nil(t)

	// fail
	// actually.Got("").Nil(t)
	// actually.Got(nil).NotNil(t)
}
