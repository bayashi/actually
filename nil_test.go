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

	// fail now
	//actually.Got("").FailNow().Nil(t)

	// fail
	// actually.Got("").Nil(t)
	// actually.Got(nil).NotNil(t)
}
