package actually

import (
	"testing"

	w "github.com/bayashi/witness"
)

// Panic asserts that a test function you got panics
/*
	actually.Got(func(){ panic("OMG") }).Panic(t) // Pass
*/
func (a *testingA) Panic(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !isFuncType(a.got) {
		w := w.Got(a.got)
		return a.fail(w, reason_GotShouldFuncType)
	}

	if didPanic, _ := didPanic(a.got.(func())); !didPanic {
		w := w.Got(a.got)
		return a.fail(w, reason_ExpectPanic)
	}

	return a
}

// PanicMessage asserts that a test function you got panics, and
// a recovered panic message is same as you expect
/*
	actually.Got(func(){ panic("OMG") }).Expect("OMG").PanicMessage(t) // Pass
*/
func (a *testingA) PanicMessage(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !isFuncType(a.got) {
		w := w.Got(a.got)
		return a.fail(w, reason_GotShouldFuncType)
	}

	didPanic, panicMessage := didPanic(a.got.(func()))

	if !didPanic {
		w := w.Got(a.got)
		return a.fail(w, reason_ExpectPanic)
	}

	if !objectsAreSameType(a.expect, panicMessage) {
		w := w.Got(panicMessage).Message(gotFunc_Label, w.Dump(a.got)).Expect(a.expect)
		return a.fail(w, reason_PanicButMsgwrongType)
	}

	if !objectsAreSame(a.expect, panicMessage) {
		w := w.Got(panicMessage).Message(gotFunc_Label, w.Dump(a.got)).Expect(a.expect)
		return a.fail(w, reason_PanicButMsgDifferent)
	}

	return a
}

// NoPanic asserts that a test function you got doesn't panic
/*
	actually.Got(func(){ panic("OMG") }).NoPanic(t) // Fail
*/
func (a *testingA) NoPanic(t *testing.T, testNames ...string) *testingA {
	invalidCall(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !isFuncType(a.got) {
		w := w.Got(a.got)
		return a.fail(w, reason_GotShouldFuncType)
	}

	if didPanic, panicMessage := didPanic(a.got.(func())); didPanic {
		w := w.Got(panicMessage).Message(gotFunc_Label, w.Dump(a.got))
		return a.fail(w, reason_ExpectNoPanic)
	}

	return a
}

func didPanic(f func()) (did bool, panicMessage any) {
	did = true
	defer func() {
		panicMessage = recover()
	}()
	f()
	did = false

	return
}
