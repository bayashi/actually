package actually

// internal utils

import (
	"fmt"
	"os"
	"strings"

	w "github.com/bayashi/witness"
)

func (a *testingA) fail(w *w.Witness, reason string) *testingA {
	a.t.Helper()
	a.doFail(w, reason)

	return a
}

func (a *testingA) failf(w *w.Witness, reasonFormat string, args ...any) *testingA {
	a.t.Helper()
	a.doFail(w, fmt.Sprintf(reasonFormat, args...))

	return a
}

var funcFail = func(a *testingA, w *w.Witness, reason string) {
	a.t.Helper()
	if a.failNow != nil && !*a.failNow {
		w.Fail(a.t, reason)
	} else if (a.failNow != nil && *a.failNow) || len(os.Getenv(envKey_FailNow)) > 0 {
		w.FailNow(a.t, reason)
	} else {
		w.Fail(a.t, reason)
	}
}

func (a *testingA) doFail(w *w.Witness, reason string) {
	a.t.Helper()
	a.failed = true
	funcFail(a, w, reason)
}

func (a *testingA) naming(testNames ...string) string {
	if a.name != "" {
		if len(testNames) > 0 {
			n := []string{a.name}
			n = append(n, testNames...)
			return strings.Join(n, ".")
		} else {
			return a.name
		}
	} else {
		return strings.Join(testNames, ".")
	}
}

func invalidCall(a *testingA) {
	if !a.setGot {
		panic(panicReason_NotCalledGot)
	}
}
