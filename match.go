package actually

import (
	"fmt"
	"regexp"
	"testing"

	w "github.com/bayashi/witness"
)

// Match method asserts that test data you got match expected value as regexp.
/*
	actually.Got("target string").Expect(`.ing$`).Match(t) // pass
*/
func (a *TestingA) Match(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	var r *regexp.Regexp
	if rr, ok := a.expect.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(a.expect))
	}

	target := fmt.Sprint(a.got)
	if !r.MatchString(target) {
		wi := w.Message("Regexp", r.String()).Message("Target", target)
		return a.fail(wi, "Not matched the regexp")
	}

	return a
}

// NotMatch method asserts that test data you got don't match expected value as regexp.
/*
	actually.Got("target string").Expect(`^[a-z]+$`).NotMatch(t) // pass
*/
func (a *TestingA) NotMatch(t *testing.T, testNames ...string) *TestingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	var r *regexp.Regexp
	if rr, ok := a.expect.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(a.expect))
	}

	target := fmt.Sprint(a.got)
	if r.MatchString(target) {
		wi := w.Message("Regexp", r.String()).Message("Target", target)
		return a.fail(wi, "Unexpectedly matched the regexp")
	}

	return a
}
