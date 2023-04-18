package diff

import (
	"testing"
)

func TestDiff(t *testing.T) {
	type S struct {
		id   int
		name string
	}

	a := &S{
		id:   123,
		name: "aiko",
	}
	b := &S{
		id:   124,
		name: "eiko",
	}

	got := "\n" + Diff(a, b) // "\n" is for heredoc.
	expect := `
--- Expected
+++ Actually got
@@ -1,4 +1,4 @@
 (*diff.S)({
- id: (int) 123,
- name: (string) (len=4) "aiko"
+ id: (int) 124,
+ name: (string) (len=4) "eiko"
 })
`
	if got != expect {
		t.Errorf("Diff was wrong.\n[Got]\n%s\n\n[Expect]\n%s\n", got, expect)
	}
}
