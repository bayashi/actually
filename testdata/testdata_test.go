package testdata

import (
	"fmt"
	"testing"
)

type Example struct {
	id int
	name string
}

func TestTruncate(t *testing.T) {
	len := 10
	td := NewTestData(&Example{id: 12, name: "John Doe"}, len)

	tts := []struct {
		format string
		expect string
	}{
		{ format: "%v",  expect: "&{12 John <... truncated>"},
		{ format: "%+v", expect: "&{id:12 na<... truncated>"},
		{ format: "%#v", expect: "&testdata.<... truncated>"},
		{ format: "%s",  expect: "&{%!s(int=<... truncated>"},
		{ format: "%Y",  expect: "*testdata.Example"}, // The type of RawValue()
	}
	for _, tt := range tts {
		if got := fmt.Sprintf(tt.format, td); got != tt.expect {
			t.Errorf("got:%s != expect:%s", got, tt.expect)
		}
	}
}
