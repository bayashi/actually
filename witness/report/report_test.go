package report

import (
	"testing"

	"github.com/bayashi/actually/witness/obj"
)

func TestReport(t *testing.T) {
	f := NewFailure()
	f.Trace("trace").Diff("diff")
	got := f.Put()
	expect := "\n\tTrace:       \ttrace\n\tDiff details:\tdiff\n"
	if expect != got {
		t.Errorf("Report is wrong.\n[Expect]\n%#v\n\n[Got]\n%#v\n", expect, got)
		t.Fail()
	}
}

func TestSetter(t *testing.T) {
	f := NewFailure()

	f = f.Trace("t")
	if f.trace != "t" {
		t.Errorf("Wrong Trace(): %#v", f.trace)
	}

	f = f.Name("n")
	if f.name != "n" {
		t.Errorf("Wrong Name(): %#v", f.name)
	}

	f = f.Reason("r")
	if f.reason != "r" {
		t.Errorf("Wrong Reason(): %#v", f.reason)
	}

	f = f.Got(obj.NewObject("g"))
	if f.got.AsRawValue() != "g" {
		t.Errorf("Wrong Got(): %#v", f.got)
	}
	f = f.RawGot("gs")
	if f.rawGot != "gs" {
		t.Errorf("Wrong GotAsRaw(): %#v", f.rawGot)
	}

	f = f.Expect(obj.NewObject("e"))
	if f.expect.AsRawValue() != "e" {
		t.Errorf("Wrong Expect(): %#v", f.expect)
	}
	f = f.RawExpect("es")
	if f.rawExpect != "es" {
		t.Errorf("Wrong ExpectAsRaw(): %#v", f.rawExpect)
	}

	f = f.Diff("d")
	if f.diff != "d" {
		t.Errorf("Wrong Diff(): %#v", f.diff)
	}
}
