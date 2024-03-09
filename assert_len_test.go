package actually

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var (
		l  []int
		l2 [2]int
	)

	for tn, tt := range map[testName]testCase{
		"blank string": {
			expected: 0, actuallyGot: "",
		},
		"string 'aiko'": {
			expected: 4, actuallyGot: "aiko",
		},
		"string 'LLR'": {
			expected: 3, actuallyGot: "LLR",
		},
		"map": {
			expected: 1, actuallyGot: map[int]int{12: 34},
		},
		"slice": {
			expected: 2, actuallyGot: []int{1, 2},
		},
		"slice 0 length": {
			expected: 0, actuallyGot: l,
		},
		"array: it's 2 length. Not zero length even if it's undefined.": {
			expected: 2, actuallyGot: l2,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			Got(tt.actuallyGot).Expect(tt.expected).Len(t)
		})
	}

	// This test passes Go1.19 or later
	// actually.Got(&[3]int{1, 2, 3}).Expect(3).Len(t) // Ptr of array can be applied builtin `len`
}

func TestLen_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"wrong length": {
			expected: 5, actuallyGot: "aiko", expectedFailReason: fmt.Sprintf(reason_ShouldHaveItems, 5, 4),
		},
		"expect value is not int": {
			expected: "3", actuallyGot: "LLR", expectedFailReason: fmt.Sprintf(reason_ExpectvalueNotInt, "string"),
		},
		"could not be applied Len": {
			expected: 0, actuallyGot: nil, expectedFailReason: reason_CouldNotBeAppliedLen,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).Len(t)
			}, tt.expectedFailReason)
		})
	}

	// other fail cases
	// actually.Got(1234).Expect(5).Len(t)
	// actually.Got(&[]int{1, 2, 3}).Expect(3).Len(t) // Ptr of slice cannot be applied builtin `len`
	// actually.Got(struct{i int}{i: 12}).Expect(1).Len(t)
	// actually.Got("foo").Expect(int64(3)).Len(t)
}
