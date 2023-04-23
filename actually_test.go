package actually

import (
	"strings"
	"testing"
)

func TestGot(t *testing.T) {
	i := 12
	a := Got(i).Expect(i)

	if a.got.RawValue() != i {
		t.Errorf("`Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect.RawValue() != i {
		t.Errorf("`Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
}

func TestActuallyGot(t *testing.T) {
	i := 12
	a := &testingA{}
	a.Got(i).Expect(i)

	if a.got.RawValue() != i {
		t.Errorf("`actually.Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect.RawValue() != i {
		t.Errorf("`actually.Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
}

func TestExpect(t *testing.T) {
	i := 13
	a := Expect(i).Got(i)

	if a.got.RawValue() != i {
		t.Errorf("`Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect.RawValue() != i {
		t.Errorf("`Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
	}
	if a.failNow != false {
		t.Errorf("`FailNotNow()` was broken. Expected:%#v, but Actual:%#v", false, a.failNow)
	}
}

func TestActuallyExpect(t *testing.T) {
	i := 13
	a := &testingA{}
	a.Expect(i).Got(i)

	if a.got.RawValue() != i {
		t.Errorf("`actually.Got()` was broken. Expected:%#v, but Actual:%#v", i, a.got)
	}
	if a.expect.RawValue() != i {
		t.Errorf("`actually.Expect()` was broken. Expected:%#v, but Actual:%#v", i, a.expect)
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

func TestSkip(t *testing.T) {
	Got(1).NotNil(t)
	Skip(t)
	Got(2).NotNil(t)
	Got(3).NotNil(t)
}

func TestTraceinfo(t *testing.T) {
	trace := traceinfo()
	if !strings.Contains(trace, "actually_test.go:91") {
		t.Errorf("trace is wrong. Actual trace:%s", trace)
	}
}

func TestName(t *testing.T) {
	a := Got(1).Name("foo").NotNil(t)
	aa := Got(a.name).Expect("foo").Same(t, "bar")
	aaa := Got(aa.name).Expect("bar").Name("baz").Same(t, "aiko")
	Got(aaa.name).Expect("baz, aiko").Same(t)
}

func TestX(t *testing.T) {
	a := Got("beer").Expect("deer").X()
	Got(a.showRawData).True(t)
}
