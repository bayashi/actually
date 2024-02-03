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
	stubA       *TestingA
	stubWitness *w.Witness
	stubRes     string
)

func stub() {
	// reset values
	stubFailed = false
	stubA = nil
	stubWitness = nil
	stubRes = ""

	funcFail = func(a *TestingA, w *w.Witness, reason string) {
		stubFailed = true
		stubA = a
		stubWitness = w
		stubRes = reason
	}
}

func stubConfirm(t *testing.T, f func(), res string) {
	stub()

	f()

	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != res {
		t.Errorf("expected `%s`, but actually got `%s`", res, stubRes)
	}
}
