package actually

import (
	"fmt"
	"regexp"
	"testing"
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

	matched, r := matchRegexp(a.expect, a.got)
	if !matched {
		wi := a.wi().Message(message_label_Regexp, r.String()).Message(message_label_Target, matchTarget(a.got))
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

	matched, r := matchRegexp(a.expect, a.got)
	if matched {
		wi := a.wi().Message(message_label_Regexp, r.String()).Message(message_label_Target, matchTarget(a.got))
		return a.fail(wi, reason_UnexpectedlyMatch)
	}

	return a
}

// []byte is implicitly converted to a string for matching.
func matchRegexp(pattern any, got any) (bool, *regexp.Regexp) {
	var r *regexp.Regexp
	if rr, ok := pattern.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(pattern))
	}

	switch v := got.(type) {
	case []byte:
		return r.Match(v), r
	case string:
		return r.MatchString(v), r
	default:
		return r.MatchString(fmt.Sprint(v)), r
	}
}

// matchTarget returns a readable string for fail reports.
func matchTarget(got any) string {
	switch v := got.(type) {
	case []byte:
		return string(v)
	case string:
		return v
	default:
		return fmt.Sprint(v)
	}
}
