package actually

import (
	"testing"
)

func TestCmp(t *testing.T) {
	Got(123).Expect(123).Cmp(t)
}

func TestCmp_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got(123).Expect(456).Cmp(t)
	}, "Not same value")
}

func TestCmpAllowUnexported(t *testing.T) {
	x := struct {
		id   int
		Name string
	}{
		id:   1,
		Name: "aiko",
	}
	Got(x).Expect(x).CmpAllowUnexported(t)
}

func TestCmpAllowUnexported_Fail(t *testing.T) {
	x := struct {
		id   int
		Name string
	}{
		id:   1,
		Name: "aiko",
	}
	y := x
	y.id = 2
	stubConfirm(t, func() {
		Got(x).Expect(y).CmpAllowUnexported(t)
	}, "Not same value")
}
