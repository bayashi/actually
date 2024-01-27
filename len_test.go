package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestLen(t *testing.T) {
	actually.Got("aiko").Expect(4).Len(t)
	actually.Got(`LLR`).Expect(3).Len(t)
	actually.Got([]int{1, 2}).Expect(2).Len(t)
	actually.Got(&[3]int{1, 2, 3}).Expect(3).Len(t) // Ptr of array can be applied builtin `len`
	actually.Got(map[int]int{123:255}).Expect(1).Len(t)

	// fail
	// actually.Got("aiko").Expect(5).Len(t)
	// actually.Got(1234).Expect(5).Len(t)
	// actually.Got(&[]int{1, 2, 3}).Expect(3).Len(t) // Ptr of slice cannot be applied builtin `len`
	// actually.Got(struct{i int}{i: 12}).Expect(1).Len(t)
	// actually.Got("foo").Expect(int64(3)).Len(t)
}
