package report

import (
	"testing"
)

func TestReport(t *testing.T) {
	r := New()
	r.Trace("trace").Diff("diff")
	got := r.Put()
	expect := "\tTrace:       \ttrace\n\tDiff Details:\tdiff\n"
	if expect != got {
		t.Errorf("Report is wrong.\n[Expect]\n%#v\n\n[Got]\n%#v\n", expect, got)
		t.Fail()
	}
}

func TestSetter(t *testing.T) {
	r := New()

	r = r.Trace("t")
	if r.trace != "t" {
		t.Errorf("Wrong Trace(): %#v", r.trace)
	}

	r = r.Function("f")
	if r.function != "f" {
		t.Errorf("Wrong Function(): %#v", r.function)
	}

	r = r.Name("n")
	if r.name != "n" {
		t.Errorf("Wrong Name(): %#v", r.name)
	}

	r = r.Reason("r")
	if r.reason != "r" {
		t.Errorf("Wrong Reason(): %#v", r.reason)
	}

	r = r.Notice("notice")
	if r.notice != "notice" {
		t.Errorf("Wrong Notice(): %#v", r.notice)
	}
	r = r.Noticef("notice:%s", "Keira")
	if r.notice != "notice:Keira" {
		t.Errorf("Wrong Noticef(): %#v", r.notice)
	}

	r = r.Got("g")
	if r.got != "g" {
		t.Errorf("Wrong Got(): %#v", r.got)
	}
	r = r.Gotf("gf")
	if r.got != "gf" {
		t.Errorf("Wrong Gotf(): %#v", r.got)
	}
	r = r.GotAsString("gs")
	if r.gotAsString != "gs" {
		t.Errorf("Wrong GotAsString(): %#v", r.gotAsString)
	}

	r = r.Expect("e")
	if r.expect != "e" {
		t.Errorf("Wrong Expect(): %#v", r.expect)
	}
	r = r.Expectf("ef")
	if r.expect != "ef" {
		t.Errorf("Wrong Expectf(): %#v", r.expect)
	}
	r = r.ExpectAsString("es")
	if r.expectAsString != "es" {
		t.Errorf("Wrong ExpectAsString(): %#v", r.expectAsString)
	}

	r = r.Diff("d")
	if r.diff != "d" {
		t.Errorf("Wrong Diff(): %#v", r.diff)
	}
}
