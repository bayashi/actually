package actually

import (
	"testing"
)

func TestMatch(t *testing.T) {
	Got("target string").Expect(`.ing$`).Match(t)
	Got("target string").Expect(`^[a-z]+$`).NotMatch(t)
}

func TestMatch_Fail(t *testing.T) {
	stub()
	Got("target string").Expect(`^[a-z]+$`).Match(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_NotMatch {
		t.Errorf("expected `%s`, but got `%s`", reason_NotMatch, stubRes)
	}

	stub()
	Got("target string").Expect(`.ing$`).NotMatch(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_UnexpectedlyMatch {
		t.Errorf("expected `%s`, but got `%s`", reason_UnexpectedlyMatch, stubRes)
	}
}
