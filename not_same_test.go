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

func TestNotSameNumber(t *testing.T) {
	actually.Got(1).Expect(2).NotSameNumber(t, "these should be different number")
	actually.Got(int8(1)).Expect(int32(2)).NotSameNumber(t)
	actually.Got(float32(1.1)).Expect(int64(1)).NotSameNumber(t)
	actually.Got(1).Expect(float64(1.000000000000001)).NotSameNumber(t)

	// fail
	// actually.Got(1).Expect(1).NotSameNumber(t)
	// actually.Got("1").Expect(1).NotSameNumber(t)
	// actually.Got(1).Expect("1").NotSameNumber(t)
	// actually.Got(nil).Expect(nil).NotSameNumber(t)
	// actually.Got(nil).Expect(0).NotSameNumber(t)
	// actually.Got(0).Expect(nil).NotSameNumber(t)
	// actually.Got([]byte("0")).Expect([]byte("0")).NotSameNumber(t)
	// actually.Got("0").Expect("0").NotSameNumber(t)
	// actually.Got(1).Expect(float64(1.0000000000000001)).NotSameNumber(t)
}

func TestNotSameType(t *testing.T) {
	actually.Got(nil).Expect(0).NotSameType(t)
	actually.Got("1").Expect(1).NotSameType(t)
	actually.Got(t).Expect(&testing.B{}).NotSameType(t)

	// fail
	// actually.Got(nil).Expect(nil).NotSameType(t)
	// actually.Got(true).Expect(false).NotSameType(t) // both are same boolean
	// actually.Got(&testing.T{}).Expect(t).NotSameType(t)
}
