package obj

import (
	"fmt"
	"testing"

	tu "github.com/bayashi/actually/witness/testutil"
	"github.com/yassinebenaid/godump"
)

type Example struct {
	id   int
	name string
}

func TestTruncate(t *testing.T) {
	o := NewObjectWithMaxLen(&Example{id: 12, name: "John Doe"}, 10)

	tts := []struct {
		format string
		expect string
	}{
		{format: "%v", expect: "&{12 John <... truncated>"},
		{format: "%+v", expect: "&{id:12 na<... truncated>"},
		{format: "%#v", expect: "&obj.Examp<... truncated>"},
		{format: "%s", expect: "&{%!s(int=<... truncated>"},
	}
	for _, tt := range tts {
		if got := fmt.Sprintf(tt.format, o); got != tt.expect {
			t.Errorf("got:%s != expect:%s", got, tt.expect)
		}
	}
}

func TestRawValue(t *testing.T) {
	o := NewObject("John Doe")
	if o.AsRawValue() != "John Doe" {
		t.Error("RawValue() was wrong")
	}
}

func TestIsStringType(t *testing.T) {
	if o := NewObject("John Doe"); !o.IsStringType() {
		t.Error("IsStringType() was wrong")
	}
	if o := NewObject(7); o.IsStringType() {
		t.Error("IsStringType() was wrong")
	}
}

func TestIsDumpableRawType(t *testing.T) {
	if o := NewObject([]int{1, 2}); !o.IsDumpableRawType() {
		t.Error("IsDumpableRawType() was wrong")
	}
	if o := NewObject(7); o.IsDumpableRawType() {
		t.Error("IsDumpableRawType() was wrong")
	}
}

func TestIsPointerType(t *testing.T) {
	i := 123
	if o := NewObject(&i); !o.IsPointerType() {
		t.Error("IsPointerType() was wrong")
	}
	if o := NewObject(7); o.IsPointerType() {
		t.Error("IsPointerType() was wrong")
	}
}

func TestBooleanPointer(t *testing.T) {
	f := false
	o := NewObject(&f)

	if ok, msg := tu.Match(`\(\*bool\)\([0-9a-fx]+\)\(false\)`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestError(t *testing.T) {
	e := fmt.Errorf("foo error")
	o := NewObject(e)

	if ok, msg := tu.Match(`\(\*errors\.errorString\)\([0-9a-fx]+\)\(foo error\)`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestPointerValueInt(t *testing.T) {
	i := 123
	o := NewObject(&i)

	if ok, msg := tu.Match(`\(\*int\)\([0-9a-fx]+\)\(123\)`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestPointerValueFloat(t *testing.T) {
	var i float64 = 0.123
	o := NewObject(&i)

	if ok, msg := tu.Match(`\(*float64\)\([0-9a-fx]+\)\(0\.123\)`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestStructValue(t *testing.T) {
	i := struct{ ID int }{ID: 123}
	o := NewObject(i)

	if ok, msg := tu.Match(`struct { ID int }{ID:123}`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestStructPointerValue(t *testing.T) {
	i := struct{ ID int }{ID: 123}
	o := NewObject(&i)

	if ok, msg := tu.Match(`[0-9a-fx]+, &struct { ID int }{ID:123}`, o.AsString()); !ok {
		t.Error(msg)
	}
}

func TestDump(t *testing.T) {
	o := NewObject(123)
	if o.AsDumpString() != "(int) 123\n" {
		t.Error("Dump() was wrong")
	}
}

func TestCustomDumper(t *testing.T) {
	o := NewObjectWithDumper(map[int]int{1: 123}, func(d any) string {
		var dumper godump.Dumper
		return dumper.Sprint(d)
	})
	if o.AsDumpString() != "map[int]int:1 {\n   1: 123,\n}" {
		t.Error("Dump() was wrong")
	}
}
