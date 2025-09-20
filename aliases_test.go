package actually

import (
	"testing"
)

func TestActual(t *testing.T) {
	a := Actual("g")
	if a.got != "g" {
		t.Errorf("Actual method was wrong. Actually got %s", a.got)
	}

	ac := &testingA{}
	ac.Actual("G")
	if ac.got != "G" {
		t.Errorf("Actual method was wrong. Actually got %s", ac.got)
	}
}

func TestWant(t *testing.T) {
	a := Want("e")
	if a.expect != "e" {
		t.Errorf("Want method was wrong. Actually got %s", a.expect)
	}

	ac := &testingA{}
	ac.Want("E")
	if ac.expect != "E" {
		t.Errorf("Want method was wrong. Actually got %s", ac.expect)
	}
}
