package diff

import (
	"testing"
	"time"
)

func TestDiffStruct(t *testing.T) {
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

	got := "\n" + Diff(a, b) // "\n" is for comparing to heredoc.
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

func TestDiffString(t *testing.T) {
	a := "beer"
	b := "deer"

	got := "\n" + Diff(a, b)
	expect := `
--- Expected
+++ Actually got
@@ -1 +1 @@
-beer
+deer
`
	if got != expect {
		t.Errorf("Diff was wrong.\n[Got]\n%s\n\n[Expect]\n%s\n", got, expect)
	}
}

func TestDiffTime(t *testing.T) {
	a := time.Date(2023, 4, 23, 0, 0, 0, 0, time.UTC)
	b := time.Date(2023, 4, 24, 0, 0, 0, 0, time.UTC)

	got := "\n" + Diff(a, b)
	expect := `
--- Expected
+++ Actually got
@@ -1,2 +1,2 @@
-(time.Time) 2023-04-23 00:00:00 +0000 UTC
+(time.Time) 2023-04-24 00:00:00 +0000 UTC
 
` // What is last space? It would be created by difflib.GetUnifiedDiffString?? Anyway, I ignore :-)
	if got != expect {
		t.Errorf("Diff was wrong.\n[Got]\n%#v\n\n[Expect]\n%#v\n", got, expect)
	}
}

func TestDiffNil(t *testing.T) {
	if Diff(nil, nil) != "" {
		t.Error("<nil> input should be blank string")
	}

	if Diff(1, "string") != "" {
		t.Error("Input values should be same type")
	}

	if Diff(1, 2) != "" {
		t.Error("NOT supported int type")
	}
}
