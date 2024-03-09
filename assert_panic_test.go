package actually

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	Got(func() { panic("OMG") }).Panic(t)
}

func TestPanic_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got(nil).Panic(t)
	}, reason_GotShouldFuncType)
	stubConfirm(t, func() {
		Got(func() {}).Panic(t)
	}, reason_ExpectPanic)
}

func TestPanicMessage(t *testing.T) {
	Got(func() { panic("hidebu") }).Expect("hidebu").PanicMessage(t)
	Got(func() { panic(fmt.Errorf("abeshi")) }).Expect(fmt.Errorf("abeshi")).PanicMessage(t)
}

func TestPanicMessage_Fail(t *testing.T) {
	for tn, tt := range map[testName]testCase{
		"Go should be func type": {
			expected: "any", actuallyGot: nil, expectedFailReason: reason_GotShouldFuncType,
		},
		"Expect panic, but didn't": {
			expected: "any", actuallyGot: func() {}, expectedFailReason: reason_ExpectPanic,
		},
		"the types of expect and panic message are different": {
			expected: "OMG ERROR", actuallyGot: func() { panic(fmt.Errorf("OMG ERROR")) }, expectedFailReason: reason_PanicButMsgwrongType,
		},
		"expect and panic message are wrong": {
			expected: "OMG", actuallyGot: func() { panic("OOPS") }, expectedFailReason: reason_PanicButMsgDifferent,
		},
	} {
		t.Run(tn, func(t *testing.T) {
			stubConfirm(t, func() {
				Got(tt.actuallyGot).Expect(tt.expected).PanicMessage(t)
			}, tt.expectedFailReason)
		})
	}
}

func TestNoPanic(t *testing.T) {
	Got(func() {}).NoPanic(t)
}

func TestNoPanic_Fail(t *testing.T) {
	stubConfirm(t, func() {
		Got(nil).NoPanic(t)
	}, reason_GotShouldFuncType)
	stubConfirm(t, func() {
		Got(func() { panic("OMG") }).NoPanic(t)
	}, reason_ExpectNoPanic)
}
