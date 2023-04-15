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

func TestExpect(t *testing.T) {
	i := 13
	a := Expect(i).Got(i)

	if a.got != i {
		t.Errorf("`Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect != i {
		t.Errorf("`Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
	if a.failNow != false {
		t.Errorf("`FailNotNow()` was broken. Expected:%#v, but Actual:%#v", false, a.failNow)
	}
}

func TestFail(t *testing.T) {
	a := Got(nil)
	if a.failNow != false {
		t.Errorf("Default failNow is false, but Actual:%#v", a.failNow)
	}

	a.FailNow()
	if a.failNow != true {
		t.Errorf("`FailNow()` was broken. Expected:%#v, but Actual:%#v", true, a.failNow)
	}

	a.FailNotNow()
	if a.failNow != false {
		t.Errorf("`FailNotNow()` was broken. Expected:%#v, but Actual:%#v", false, a.failNow)
	}
}

func TestDuplicateCall(t *testing.T) {
	//Got(1).Got(1).NotNil(t)
	//Expect(1).Expect(1).NotNil(t)
}
