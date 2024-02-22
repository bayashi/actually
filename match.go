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
func (a *testingA) Match(t *testing.T, testNames ...string) *testingA {
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
		wi := w.Message(message_label_Regexp, r.String()).Message(message_label_Target, target)
		return a.fail(wi, reason_NotMatch)
	}

	return a
}

// NotMatch method asserts that test data you got don't match expected value as regexp.
/*
	actually.Got("target string").Expect(`^[a-z]+$`).NotMatch(t) // pass
*/
func (a *testingA) NotMatch(t *testing.T, testNames ...string) *testingA {
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
		wi := w.Message(message_label_Regexp, r.String()).Message(message_label_Target, target)
		return a.fail(wi, reason_UnexpectedlyMatch)
	}

	return a
}
