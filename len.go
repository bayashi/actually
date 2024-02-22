package actually

import (
	"reflect"
	"testing"

	w "github.com/bayashi/witness"
)

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
func (a *testingA) Len(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	wi := w.Got(a.got).Expect(a.expect)

	if k, ok := isValidExpect(a.expect); !ok {
		return a.failf(wi, reason_ExpectvalueNotInt, k)
	}

	gotLen, ok := getLen(a.got)
	if !ok {
		return a.fail(wi, reason_CouldNotBeAppliedLen)
	}
	if gotLen != a.expect {
		return a.failf(wi, reason_ShouldHaveItems, a.expect, gotLen)
	}

	return a
}

func isValidExpect(e any) (reflect.Kind, bool) {
	k := reflect.ValueOf(e).Kind()

	return k, k == reflect.Int || k == reflect.Int32
}

// getLen tries to get the length of an object.
// It returns (0, false) if impossible.
func getLen(x interface{}) (length int, ok bool) {
	v := reflect.ValueOf(x)
	defer func() {
		ok = recover() == nil
	}()

	return v.Len(), true
}
