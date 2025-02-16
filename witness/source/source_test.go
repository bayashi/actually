package source

import (
	"testing"
)

func TestGetSource(t *testing.T) {
	got, err := GetSouce("../../testdata/example_source.go", 8)
	if err != nil {
		t.Errorf("failed GetSouce %s", err)
	}

	expect := []string{
		"05  )",
		"06  ",
		"07  func main() {",
		"08> \tfmt.Println(\"Hello World\")",
		"09  }",
	}

	if len(expect) != len(got) {
		t.Errorf("expect `%#v`, got `%#v`", expect, got)
	}

	for i, e := range expect {
		if e != got[i] {
			t.Errorf("expect `%s`, but got `%s`", e, got[i])
		}
	}
}
