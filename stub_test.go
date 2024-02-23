package actually

import (
	"testing"

	w "github.com/bayashi/witness"
)

const (
	notCalledFail = "test was unexpectedly OK. funcFail wouldn't be called"
)

var (
	stubFailed  bool
	stubA       *testingA
	stubWitness *w.Witness
	stubRes     string
)

func stub() {
	// reset values
	stubFailed = false
	stubA = nil
	stubWitness = nil
	stubRes = ""

	funcFail = func(a *testingA, w *w.Witness, reason string) {
		stubFailed = true
		stubA = a
		stubWitness = w
		stubRes = reason
	}
}

func stubConfirm(t *testing.T, f func(), res string) {
	t.Helper()

	stub()

	f()

	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != res {
		Fail(t, "fail messages are wrong", stubRes, res)
	}
}
