package actually

import (
	"fmt"
	"testing"
)

func TestSame(t *testing.T) {
	Got(nil).Expect(nil).Same(t)
	Got("").Expect("").Same(t)
	Got("a").Expect("a").Same(t)
	Got(0).Expect(0).Same(t)
	Got(12).Expect(12).Same(t)
	Got([2]int{1, 2}).Expect([2]int{1, 2}).Same(t)
	Got([]string{}).Expect([]string{}).Same(t)
	Got([]string{"a"}).Expect([]string{"a"}).Same(t)
	Got(map[string]int{"foo": 12}).Expect(map[string]int{"foo": 12}).Same(t)

	foo := struct {
		bar string
	}{
		bar: "foo",
	}
	bar := struct {
		bar string
	}{
		bar: "foo",
	}
	Got(foo).Expect(bar).Same(t)
	Got(fmt.Sprintf("%p", &foo) == fmt.Sprintf("%p", &bar)).False(t)

	// test name
	Got(0).Expect(0).Same(t, "zero")

	// NOT SUPPORT chan YET
	// ch1 := make(chan string, 1)
	// ch1 <- "foo"
	// ch2 := make(chan string, 1)
	// ch2 <- "foo"
	// actually.Got(ch1).Expect(ch2).Same(t)
}

func TestSame_Fail(t *testing.T) {
	stub()
	Got("a").Expect("b").Same(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_NotSame {
		t.Errorf("expected `%s`, but got `%s`", reason_NotSame, stubRes)
	}

	stub()
	Got(int16(12)).Expect(int32(12)).Same(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongType {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongType, stubRes)
	}

	f := func() {}
	stub()
	Got(f).Expect(f).Same(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsFunc {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsFunc, stubRes)
	}
}

func TestSamePointer(t *testing.T) {
	i := 7
	ptr := &i
	ptr2 := ptr
	Got(ptr).Expect(ptr2).SamePointer(t)
	Got(ptr).Expect(&i).SamePointer(t)

	// test name
	Got(ptr).Expect(ptr2).SamePointer(t, "Same Pointer")

	// fail
	// actually.Got("").Expect(ptr).SamePointer(t)
	// actually.Got(ptr).Expect("").SamePointer(t)
	// j := 7
	// actually.Got(ptr).Expect(&j).SamePointer(t)
}

func TestSamePointer_Fail(t *testing.T) {
	i := 7
	ptr := &i

	stub()
	Got(1).Expect(ptr).SamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNotPointer {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNotPointer, stubRes)
	}
	stub()
	Got(ptr).Expect(1).SamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNotPointer {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNotPointer, stubRes)
	}

	j := 7
	stub()
	Got(ptr).Expect(&j).SamePointer(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongPointerAddress {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongPointerAddress, stubRes)
	}
}

func TestSameNumber(t *testing.T) {
	Got(int8(1)).Expect(int32(1)).SameNumber(t)
	Got(float32(1.0)).Expect(int64(1)).SameNumber(t)

	// test name
	Got(1).Expect(1).SameNumber(t, "Same Number")
}

func TestSameNumber_Fail(t *testing.T) {
	stub()
	Got("1").Expect(1).SameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNotNumber {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNotNumber, stubRes)
	}
	stub()
	Got(1).Expect("1").SameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNotNumber {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNotNumber, stubRes)
	}

	stub()
	Got(nil).Expect(0).SameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_GotIsNilType {
		t.Errorf("expected `%s`, but got `%s`", reason_GotIsNilType, stubRes)
	}
	stub()
	Got(0).Expect(nil).SameNumber(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_ExpectIsNilType {
		t.Errorf("expected `%s`, but got `%s`", reason_ExpectIsNilType, stubRes)
	}
}

func TestChain(t *testing.T) {
	Got(7).NotNil(t).
		Expect(7).SameNumber(t).Same(t)
}

func TestSameType(t *testing.T) {
	Got(nil).Expect(nil).SameType(t)
	Got(true).Expect(false).SameType(t) // both are boolean
	a := Got(t).Expect(t).SameType(t)
	Got(a).Expect(a).SameType(t)
}

func TestSameType_Fail(t *testing.T) {
	stub()
	Got("1").Expect(1).SameType(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongType {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongType, stubRes)
	}

	stub()
	Got(nil).Expect(0).SameType(t)
	if !stubFailed {
		t.Error(notCalledFail)
	}
	if stubRes != reason_WrongType {
		t.Errorf("expected `%s`, but got `%s`", reason_WrongType, stubRes)
	}
}
