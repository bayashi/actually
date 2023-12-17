package actually

import (
	"os"
	"testing"
)

const ExistingFileInThisModule = "README.md"
const NotExistingFileInThisModule = "not-found-01xt79a5jk.exe"

func TestNoError(t *testing.T) {
	_, err := os.Open(ExistingFileInThisModule)
	Got(err).NoError(t)

	// Fail case
	// _, err2 := os.Open(NotExistingFileInThisModule)
	// Got(err2).NoError(t)
}

func TestIsTypeOfError(t *testing.T) {
	if Got(123).isTypeOfError() != false {
		t.Error("Integer should NOT be a type of error")
	}

	_, err := os.Open(NotExistingFileInThisModule)
	if Got(err).isTypeOfError() != true {
		t.Errorf("%#v should be a type of error", err)
	}
}

func TestGotError(t *testing.T) {
	_, err := os.Open(NotExistingFileInThisModule)
	a := GotError(err)
	if a.got != err {
		t.Errorf("`GotError()` was broken. Expected:%#v, but Actual:%#v", err, a.got)
	}

	_, err2 := os.Open(ExistingFileInThisModule)
	a2 := GotError(err2)
	if a2.got != err2 {
		t.Errorf("`GotError()` was broken. Expected:%#v, but Actual:%#v", err2, a2.got)
	}
}

func TestActuallyGotError(t *testing.T) {
	_, err := os.Open(NotExistingFileInThisModule)
	a := &TestingA{}
	a.GotError(err)
	if a.got != err {
		t.Errorf("`actually.GotError()` was broken. Expected:%#v, but Actual:%#v", err, a.got)
	}

	_, err2 := os.Open(ExistingFileInThisModule)
	a2 := &TestingA{}
	a2.GotError(err2)
	if a2.got != err2 {
		t.Errorf("`actually.GotError()` was broken. Expected:%#v, but Actual:%#v", err2, a2.got)
	}
}
