package actually

import (
	"os"
	"testing"
)

func TestActual(t *testing.T) {
	a := Actual("g")
	if a.got != "g" {
		t.Errorf("Actual method was wrong. Actually got %s", a.got)
	}

	ac := &TestingA{}
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

	ac := &TestingA{}
	ac.Want("E")
	if ac.expect != "E" {
		t.Errorf("Want method was wrong. Actually got %s", ac.expect)
	}
}

func TestFatal(t *testing.T) {
	a := Got("g").Fatal()
	if !*a.failNow {
		t.Error("expect failNow is false")
	}
}

func TestFatalOn(t *testing.T) {
	if os.Getenv(envKey_FailNow) != "" {
		t.Error("Already set ENV somehow")
	}
	FatalOn(t)
	if os.Getenv(envKey_FailNow) == "" {
		t.Error("expect failNow is false")
	}
}
