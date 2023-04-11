package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestSameString(t *testing.T) {
	// pass
	actually.Got("").Expect("").SameString(t)
	actually.Got("a").Expect("a").SameString(t)

	// fail
	//actually.Got("a").Expect("b").SameString(t)
}
