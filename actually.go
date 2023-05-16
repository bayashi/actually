// Yet another pithy testing framework `actually`
package actually

import (
	"testing"

	"github.com/bayashi/actually/testobject"
)

// TestingA is a context of the test
type TestingA struct {
	got         *testobject.TestObject
	setGot      bool
	expect      *testobject.TestObject
	setExpect   bool
	t           *testing.T
	failNow     *bool
	showRawData bool
	name        string
}

// Got sets the value you actually got. Got() creates *TestingA and returns it.
func Got(g any) *TestingA {
	return &TestingA{
		got:    testobject.NewTestObject(g, 0),
		setGot: true,
	}
}

// Got sets the value you actually got.
func (a *TestingA) Got(g any) *TestingA {
	if a.setGot {
		panic(panicReason_CalledGotTwice)
	}

	a.got = testobject.NewTestObject(g, 0)
	a.setGot = true

	return a
}

// Expect sets the value you expect to be the same as the one you got.
// Expect creates *TestingA and returns it.
func Expect(e any) *TestingA {
	return &TestingA{
		expect:    testobject.NewTestObject(e, 0),
		setExpect: true,
	}
}

// Expect sets the value you expect to be the same as the one you got.
func (a *TestingA) Expect(e any) *TestingA {
	if a.setExpect {
		panic(panicReason_CalledExpectTwice)
	}

	a.expect = testobject.NewTestObject(e, 0)
	a.setExpect = true

	return a
}
