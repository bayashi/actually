package actually

import (
	"regexp"
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

func TestMatch_ByteSlice(t *testing.T) {
	// []byte is automatically converted to a string for matching.
	Got([]byte("target string")).Expect(`.ing$`).Match(t)
	Got([]byte("target string")).Expect(`^[a-z]+$`).NotMatch(t)
	Got([]byte("target string")).Expect(regexp.MustCompile(`.ing$`)).Match(t)
}

func TestMatch_ByteSlice_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got([]byte("target string")).Expect(`^[a-z]+$`).Match(t)
	}, reason_NotMatch)

	stubConfirm(t, func() {
		Got([]byte("target string")).Expect(`.ing$`).NotMatch(t)
	}, reason_UnexpectedlyMatch)
}
