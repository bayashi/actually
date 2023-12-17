// Yet another pithy testing framework `actually`
package actually

import (
	"testing"
)

// TestingA is a context of the test
type TestingA struct {
	got         any
	setGot      bool
	expect      any
	setExpect   bool
	t           *testing.T
	failNow     *bool
	showRawData bool
	name        string
}

// Got sets the value you actually got. Got() creates *TestingA and returns it.
func Got(g any) *TestingA {
	return &TestingA{
		got:    g,
		setGot: true,
	}
}

// Got sets the value you actually got.
func (a *TestingA) Got(g any) *TestingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = g
	a.setGot = true

	return a
}

// Expect sets the value you expect to be the same as the one you got.
// Expect creates *TestingA and returns it.
func Expect(e any) *TestingA {
	return &TestingA{
		expect:    e,
		setExpect: true,
	}
}

// Expect sets the value you expect to be the same as the one you got.
func (a *TestingA) Expect(e any) *TestingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = e
	a.setExpect = true

	return a
}
