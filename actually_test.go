package actually

import (
	"testing"
)

func TestGot(t *testing.T) {
	i := 12
	a := Got(i).Expect(i)

	if a.got != i {
		t.Errorf("`Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect != i {
		t.Errorf("`Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
}

func TestActuallyGot(t *testing.T) {
	i := 12
	a := &TestingA{}
	a.Got(i).Expect(i)

	if a.got != i {
		t.Errorf("`actually.Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect != i {
		t.Errorf("`actually.Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
}

func TestExpect(t *testing.T) {
	i := 13
	a := Expect(i).Got(i)

	if a.got != i {
		t.Errorf("`Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect != i {
		t.Errorf("`Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
	if a.failNow != nil && *a.failNow != false {
		t.Errorf("`FailNotNow()` was broken. Expected:%#v, but Actual:%#v", false, a.failNow)
	}
}

func TestActuallyExpect(t *testing.T) {
	i := 13
	a := &TestingA{}
	a.Expect(i).Got(i)

	if a.got != i {
		t.Errorf("`actually.Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect != i {
		t.Errorf("`actually.Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
}

func TestFail(t *testing.T) {
	a := Got(nil)
	if a.failNow != nil && *a.failNow != false {
		t.Errorf("Default failNow is false, but Actual:%#v", a.failNow)
	}

	a.FailNow()
	if *a.failNow != true {
		t.Errorf("`FailNow()` was broken. Expected:%#v, but Actual:%#v", true, a.failNow)
	}

	a.FailNotNow()
	if *a.failNow != false {
		t.Errorf("`FailNotNow()` was broken. Expected:%#v, but Actual:%#v", false, a.failNow)
	}
}

func TestDuplicateCall(t *testing.T) {
	f := false
	defer func() {
		err := recover()
		if err != panicReason_CalledGotTwice {
			t.Errorf("expect error %s, but got %+v", panicReason_CalledGotTwice, err)
		}
		f = true
	}()
	Got(1).Got(2) // duplicate calling Got should be panic
	if !f {
		t.Error("panic wouldn't happen")
	}

	f = false
	defer func() {
		err := recover()
		if err != panicReason_CalledExpectTwice {
			t.Errorf("expect error %s, but got %+v", panicReason_CalledExpectTwice, err)
		}
		f = true
	}()
	Expect(1).Expect(1) // duplicate calling Expect should be panic
	if !f {
		t.Error("panic wouldn't happen")
	}
}
