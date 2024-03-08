// Yet another pithy testing framework `actually`
package actually

import (
	"testing"
)

// testingA is a context of the test
type testingA struct {
	got         any
	setGot      bool
	expect      any
	setExpect   bool
	t           *testing.T
	failNow     *bool
	showRawData bool
	name        string
	failed      bool
	debugInfo   []map[string]any
}

// Got sets the value you actually got. Got() creates *testingA and returns it.
func Got(g any) *testingA {
	return &testingA{
		got:    g,
		setGot: true,
	}
}

// Got sets the value you actually got.
func (a *testingA) Got(g any) *testingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = g
	a.setGot = true

	return a
}

// Expect sets the value you expect to be the same as the one you got.
// Expect creates *testingA and returns it.
func Expect(e any) *testingA {
	return &testingA{
		expect:    e,
		setExpect: true,
	}
}

// Expect sets the value you expect to be the same as the one you got.
func (a *testingA) Expect(e any) *testingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = e
	a.setExpect = true

	return a
}
