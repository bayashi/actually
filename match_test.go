package actually

import (
	"testing"
)

func TestMatch(t *testing.T) {
	Got("target string").Expect(`.ing$`).Match(t)
	Got("target string").Expect(`^[a-z]+$`).NotMatch(t)
}

func TestMatch_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got("target string").Expect(`^[a-z]+$`).Match(t)
	}, reason_NotMatch)

	stubConfirm(t, func() {
		Got("target string").Expect(`.ing$`).NotMatch(t)
	}, reason_UnexpectedlyMatch)
}
