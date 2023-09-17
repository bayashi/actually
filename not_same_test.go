package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestNotSamePointer(t *testing.T) {
	i := 7
	j := 7
	ptr := &i
	ptr2 := &j
	actually.Got(ptr).Expect(ptr2).NotSamePointer(t)
	actually.Got(ptr).Expect(&j).NotSamePointer(t)
	actually.Got(ptr2).Expect(&i).NotSamePointer(t)

	// test name
	actually.Got(&i).Expect(&j).NotSamePointer(t, "Not Same Pointer")

	// fail
	// actually.Got("").Expect(ptr).NotSamePointer(t)
	// actually.Got(ptr).Expect("").NotSamePointer(t)
	// actually.Got(ptr).Expect(&i).NotSamePointer(t)
}
