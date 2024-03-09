package actually

import (
	"os"
	"testing"
)

const existingFileInThisModule = "README.md"
const notexistingFileInThisModule = "not-found-01xt79a5jk.exe"

func TestNoError(t *testing.T) {
	_, err := os.Open(existingFileInThisModule)
	Got(err).NoError(t)
}

func TestNoError_Fail(t *testing.T) {
	stubConfirm(t, func() {
		_, err := os.Open(notexistingFileInThisModule)
		Got(err).NoError(t)
	}, reason_UnexpectedlyError)
}

func TestIsTypeOfError(t *testing.T) {
	if Got(123).isTypeOfError() != false {
		t.Error("Integer should NOT be a type of error")
	}

	_, err := os.Open(notexistingFileInThisModule)
	if Got(err).isTypeOfError() != true {
		t.Errorf("%#v should be a type of error", err)
	}
}

func TestGotError(t *testing.T) {
	_, err := os.Open(notexistingFileInThisModule)
	a := GotError(err)
	if a.got != err {
		t.Errorf("`GotError()` was broken. Expected:%#v, but Actual:%#v", err, a.got)
	}

	_, err2 := os.Open(existingFileInThisModule)
	a2 := GotError(err2)
	if a2.got != err2 {
		t.Errorf("`GotError()` was broken. Expected:%#v, but Actual:%#v", err2, a2.got)
	}
}

func TestActuallyGotError(t *testing.T) {
	_, err := os.Open(notexistingFileInThisModule)
	a := &testingA{}
	a.GotError(err)
	if a.got != err {
		t.Errorf("`actually.GotError()` was broken. Expected:%#v, but Actual:%#v", err, a.got)
	}

	_, err2 := os.Open(existingFileInThisModule)
	a2 := &testingA{}
	a2.GotError(err2)
	if a2.got != err2 {
		t.Errorf("`actually.GotError()` was broken. Expected:%#v, but Actual:%#v", err2, a2.got)
	}
}
