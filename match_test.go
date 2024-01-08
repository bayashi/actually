package actually

import (
	"testing"
)

func TestMatch(t *testing.T) {
	Got("target string").Expect(`.ing$`).Match(t)
}

func TestNotMatch(t *testing.T) {
	Got("target string").Expect(`^[a-z]+$`).NotMatch(t)
}
