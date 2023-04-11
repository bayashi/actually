package report

import (
	"testing"
)

func TestReport(t *testing.T) {
	r := New()
	r.Trace("trace").Diff("diff")
	got := r.Put()
	expect := "\tTrace:\ttrace\n\tDiff: \tdiff\n"
	if expect != got {
		t.Errorf("Report is wrong.\n[Expect]\n%s\n\n[Got]\n%s\n", expect, got)
		t.Fail()
	}
}
