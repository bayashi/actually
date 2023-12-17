package actually

// internal utils

import (
	"os"
	"strings"

	w "github.com/bayashi/witness"
)

func (a *TestingA) fail(w *w.Witness, reason string) *TestingA {
	a.t.Helper()
	if a.failNow != nil && !*a.failNow {
		w.Fail(a.t, reason)
	} else if (a.failNow != nil && *a.failNow) || len(os.Getenv(envKey_FailNow)) > 0 {
		w.FailNow(a.t, reason)
	} else {
		w.Fail(a.t, reason)
	}

	return a
}

func (a *TestingA) naming(testNames ...string) string {
	if a.name != "" {
		if len(testNames) > 0 {
			n := []string{a.name}
			n = append(n, testNames...)
			return strings.Join(n, ", ")
		} else {
			return a.name
		}
	} else {
		return strings.Join(testNames, ", ")
	}
}

func invalidCall(a *TestingA) {
	if !a.setGot {
		panic("You called assertion method, but you forgot to call Got().")
	}
}
