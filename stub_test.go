package actually

import (
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
