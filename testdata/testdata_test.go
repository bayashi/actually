package testdata

import (
	"fmt"
	"testing"
)

type Example struct {
	id   int
	name string
}

func TestTruncate(t *testing.T) {
	len := 10
	td := NewTestData(&Example{id: 12, name: "John Doe"}, len)

	tts := []struct {
		format string
		expect string
	}{
		{format: "%v", expect: "&{12 John <... truncated>"},
		{format: "%+v", expect: "&{id:12 na<... truncated>"},
		{format: "%#v", expect: "&testdata.<... truncated>"},
		{format: "%s", expect: "&{%!s(int=<... truncated>"},
		{format: "%Y", expect: "*testdata.Example"}, // The type of RawValue()
	}
	for _, tt := range tts {
		if got := fmt.Sprintf(tt.format, td); got != tt.expect {
			t.Errorf("got:%s != expect:%s", got, tt.expect)
		}
	}
}

func TestRawValue(t *testing.T) {
	td := NewTestData("John Doe", 1024)
	if td.RawValue() != "John Doe" {
		t.Error("RawValue() was wrong")
	}
}

func TestIsStringType(t *testing.T) {
	if td := NewTestData("John Doe", 1024); !td.IsStringType() {
		t.Error("IsStringType() was wrong")
	}
	if td := NewTestData(7, 1024); td.IsStringType() {
		t.Error("IsStringType() was wrong")
	}
}

func TestIsDumpableRawType(t *testing.T) {
	if td := NewTestData([]int{1, 2}, 1024); !td.IsDumpableRawType() {
		t.Error("IsDumpableRawType() was wrong")
	}
	if td := NewTestData(7, 1024); td.IsDumpableRawType() {
		t.Error("IsDumpableRawType() was wrong")
	}
}

func TestDump(t *testing.T) {
	td := NewTestData(123, 1024)
	if td.Dump() != "(int) 123\n" {
		t.Error("Dump() was wrong")
	}
}
