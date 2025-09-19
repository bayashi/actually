package actually

import (
	"fmt"
	"strings"
	"testing"
)

func TestSkip(t *testing.T) {
	Got(1).NotNil(t)
	Skip(t, "Skip reason")
	Got(2).NotNil(t)
	Got(3).NotNil(t)
}

func TestTestName(t *testing.T) {
	a := Got(1).Name("foo").NotNil(t)
	aa := Got(a.name).Expect("foo").Same(t, "bar")
	aaa := Got(aa.name).Expect("bar").Name("baz").Same(t, "aiko")
	Got(aaa.name).Expect("baz.aiko").Same(t)
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

func TestDump(t *testing.T) {
	Got(Dump()).Expect("").X().Same(t)

	v := map[string]int{
		"foo": 256,
	}
	{
		expect := strings.TrimLeft(`
map[string]int:1 {
   "foo": 256,
}
`, "\n")
		Got(Dump(v)).Expect(expect).X().Same(t)
	}

	{
		expect := strings.TrimLeft(`
[0]
map[string]int:1 {
   "foo": 256,
}
[1]
map[string]int:1 {
   "foo": 256,
}
`, "\n")
		Got(Dump(v, v)).Expect(expect).X().Same(t)
	}
}

func TestFi(t *testing.T) {
	isFailed := Got(nil).Nil(t).Fi()
	Got(isFailed).False(t) // Passed, so it should be `false`

	stubConfirm(t, func() {
		if fi := Got(isFailed).True(t).Fi(); !fi {
			t.Fatal("If the test got fail, fi must return `true`. Actually, got `false`, somehow")
		}
	}, message_ExpectTrue)
}

func TestDebug(t *testing.T) {
	a := Got(1).Debug("label", 123)
	a.t = t

	stub()
	a.fail(a.wi(), "reason")

	if !strings.Contains(fmt.Sprintf("%#v", stubWitness), `{"label":[]*obj.Object{(*obj.Object)(`) {
		t.Error("not include debug info")
	}
}

func TestDebugMultipleInfo(t *testing.T) {
	a := Got(1).Debug("label", 123, 456)
	a.t = t

	stub()
	a.fail(a.wi(), "reason")

	if !strings.Contains(fmt.Sprintf("%#v", stubWitness), `), (*obj.Object)(`) {
		t.Error("not include 2nd debug info")
	}
}
