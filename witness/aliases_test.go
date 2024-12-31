package witness

import (
	"testing"

	tu "github.com/bayashi/actually/witness/testutil"
)

func TestActual(t *testing.T) {
	w := Actual(123)
	if w.got.AsRawValue() != 123 {
		t.Error("the `Actual` method could not set correct value to `got`")
	}

	w = New()
	w.Actual(255)
	if w.got.AsRawValue() != 255 {
		t.Error("the `Actual` method could not set correct value to `got`")
	}
}

func TestWant(t *testing.T) {
	w := Want(123)
	if w.expect.AsRawValue() != 123 {
		t.Error("the `Want` method could not set correct value to `expect`")
	}

	w = New()
	w.Want(255)
	if w.expect.AsRawValue() != 255 {
		t.Error("the `Want` method could not set correct value to `expect`")
	}
}

func TestFatal(t *testing.T) {
	stub()

	Fatal(t, "reason", "g")
	if !failnow {
		t.Error("the `Fatal` method is wrong")
	}
	if ok, msg := tu.Match(`Fail reason:\s*\treason`, res); !ok {
		t.Error(msg)
	}
	if ok, msg := tu.Match(`Actually got:\s*\t"g"`, res); !ok {
		t.Error(msg)
	}

	failnow = false
	w := Got("G")
	w.Fatal(t, "Reason")
	if !failnow {
		t.Error("the `Fatal` method is wrong")
	}
	if ok, msg := tu.Match(`Fail reason:\s*\tReason`, res); !ok {
		t.Error(msg)
	}
	if ok, msg := tu.Match(`Actually got:\s*\t"G"`, res); !ok {
		t.Error(msg)
	}
}
