package actually

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var l []int
	Got(l).Expect(0).Len(t)

	var l2 [2]int
	Got(l2).Expect(2).Len(t) // actually, it's 2 length. Not zero length.

	Got("").Expect(0).Len(t)

	Got("aiko").Expect(4).Len(t)
	Got(`LLR`).Expect(3).Len(t)
	Got([]int{1, 2}).Expect(2).Len(t)
	Got(map[int]int{123: 255}).Expect(1).Len(t)

	// This test passes Go1.19 or later
	// actually.Got(&[3]int{1, 2, 3}).Expect(3).Len(t) // Ptr of array can be applied builtin `len`
}

func TestLen_Fail(t *testing.T) {
	expectRes := fmt.Sprintf(reason_ShouldHaveItems, 5, 4)
	stubConfirm(t, func() {
		Got("aiko").Expect(5).Len(t)
	}, expectRes)

	expectRes = fmt.Sprintf(reason_ExpectvalueNotInt, "string")
	stubConfirm(t, func() {
		Got("LLR").Expect("3").Len(t)
	}, expectRes)

	stubConfirm(t, func() {
		Got(nil).Expect(0).Len(t)
	}, reason_CouldNotBeAppliedLen)

	// other fail cases
	// actually.Got(1234).Expect(5).Len(t)
	// actually.Got(&[]int{1, 2, 3}).Expect(3).Len(t) // Ptr of slice cannot be applied builtin `len`
	// actually.Got(struct{i int}{i: 12}).Expect(1).Len(t)
	// actually.Got("foo").Expect(int64(3)).Len(t)
}
