package actually

import (
	"testing"

	w "github.com/bayashi/actually/witness"
)

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

	a.t = t

	stub()
	a.fail(w.New(), "reason")
	if stubRes != "reason" {
		t.Error("a.fail method is wrong")
	}

	stub()
	a.failf(w.New(), "reason %s", "foo")
	if stubRes != "reason foo" {
		t.Error("a.failf method is wrong")
	}
}

func TestNaming(t *testing.T) {
	expect := "FooTest"
	if n := Got(12).naming(expect); n != expect {
		t.Errorf("Expect `%s`, but got `%s`", expect, n)
	}

	expect2 := "Name.FooTest"
	if n := Got(12).Name("Name").naming(expect); n != expect2 {
		t.Errorf("Expect `%s`, but got `%s`", expect2, n)
	}

	if n := Got(12).Name("Name").naming(expect, expect); n != expect2+"."+expect {
		t.Errorf("Expect `%s`, but got `%s`", expect2+", "+expect, n)
	}
}

func TestInvalidCall(t *testing.T) {
	f := false
	defer func() {
		err := recover()
		if err != panicReason_NotCalledGot {
			t.Errorf("expect error %s, but got %+v", panicReason_NotCalledGot, err)
		}
		f = true
	}()
	Expect(nil).Nil(t) // Without calling Got should be panic
	if !f {
		t.Error("panic wouldn't happen")
	}
}
