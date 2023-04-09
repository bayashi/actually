package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestNil(t *testing.T) {
	// pass
	actually.Got(nil).Nil(t)
	actually.Got("").NotNil(t)

	actually.Got(nil).Nil(t).Got("").NotNil(t)

	// fail now
	//actually.Got("").FailNow().Nil(t)

	// fail
	// actually.Got("").Nil(t)
	// actually.Got(nil).NotNil(t)
}
