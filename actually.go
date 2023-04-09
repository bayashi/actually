package actually

import (
	"testing"
)

type testingA struct {
	got any
	expect any
	t *testing.T
	failNow bool
}

func Got(g any) *testingA {
	return &testingA{
		got: g,
	}
}

func (a *testingA) Got(g any) *testingA {
	a.got = g

	return a
}

func Expect(e any) *testingA {
	return &testingA{
		expect: e,
	}
}

func (a *testingA) Expect(e any) *testingA {
	a.expect = e

	return a
}

func Want(e any) *testingA {
	return Expect(e)
}
func (a *testingA) Want(e any) *testingA {
	return a.Expect(e)
}

func (a *testingA) FailNotNow() *testingA {
	a.failNow = false

	return a
}

func (a *testingA) FailNow() *testingA {
	a.failNow = true

	return a
}

func (a *testingA) fail(message string) {
	a.t.Helper()
	a.t.Error(message)
	if a.failNow {
		a.t.FailNow()
	} else {
		a.t.Fail()
	}
}
