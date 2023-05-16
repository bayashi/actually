package actually

import (
	"strings"
	"testing"
)

func TestSkip(t *testing.T) {
	Got(1).NotNil(t)
	Skip(t, "Skip reason")
	Got(2).NotNil(t)
	Got(3).NotNil(t)
}

func TestTraceinfo(t *testing.T) {
	trace := traceinfo()
	if !strings.Contains(trace, "helpers_test.go:16") {
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

func TestDiff(t *testing.T) {
	got := "\n" + Diff("bar", "bug")
	expect := `
--- a
+++ b
@@ -1 +1 @@
-bar
+bug
`
	Got(got).Expect(expect).X().Same(t)
}
