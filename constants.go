package actually

const (
	envKey_FailNow = "ACTUALLY_FAIL_NOW"

	message_label_Regexp = "Regexp"
	message_label_Target = "Target"

	panicReason_CalledGotTwice    = "Don't call `Got` method twice. It has been called already"
	panicReason_CalledExpectTwice = "Don't call `Expect` method twice. It has been called already"
	panicReason_NotCalledGot      = "You called assertion method, but you forgot to call Got()"

	message_ExpectTrue  = "It should be boolean type <true>"
	message_ExpectFalse = "It should be boolean type <false>"

	// fail reason
	reason_WrongType             = "Different type"
	reason_SameType              = "Unexpectedly same type"
	reason_NotSame               = "Not same value"
	reason_Same                  = "Unexpectedly same value"
	reason_NotMatch              = "Not matched the regexp"
	reason_UnexpectedlyMatch     = "Unexpectedly matched the regexp"
	reason_UnexpectedlyError     = "Error happened"
	reason_WrongPointerAddress   = "Wrong pointer address"
	reason_SamePointerAddress    = "Unexpectedly same pointer address"
	reason_GotIsFunc             = "`Got` value is type of function"
	reason_ExpectIsFunc          = "`Expect` value is type of function"
	reason_GotShouldFuncType     = "`Got` value should be type of function"
	reason_GotIsNotPointer       = "`Got` is NOT type of Pointer"
	reason_ExpectIsNotPointer    = "`Expect` is NOT type of Pointer"
	reason_GotIsNilType          = "Type of `Got` is a <nil> value"
	reason_ExpectIsNilType       = "Type of `Expect` is a <nil> value"
	reason_GotIsNotNumber        = "Type of `Got` is not a number"
	reason_ExpectIsNotNumber     = "Type of `Expect` is not a number"
	reason_ExpectNilButNotNil    = "Expected <nil>, but it was NOT <nil>"
	reason_ExpectIsNotNil        = "Expected other than <nil>, but got <nil>"
	reason_ExpectIsNotValidValue = "`Expect` value is NOT a valid value"
	reason_NotConvertibleTypes   = "The types of `Got` and `Expect` are NOT convertible"
	reason_ExpectvalueNotInt     = "`Expect` value should be `int` or `int32` type, but it's `%s`"
	reason_CouldNotBeAppliedLen  = "`Got` value could NOT be applied builtin `len`"
	reason_ShouldHaveItems       = "`Got` value should have %d item(s), but has %d"
	reason_ExpectPanic           = "Expected panic, but no panic"
	reason_PanicButMsgwrongType  = "Did panic, but a message from panic was wrong type"
	reason_PanicButMsgDifferent  = "Did panic, but a message from panic was unexpectedly different"
	reason_ExpectNoPanic         = "Expected no panic, but did panic"

	// notice_Method_*
	notice_Label                        = "Notice"
	notice_Same_NotAcceptable           = "It's not acceptable in Same() method"
	notice_SamePointer_ShouldPointer    = "It should be a Pointer for SamePointer() method"
	notice_SameNumber_ShouldNumber      = "It should be a number(int or float) for SameNumber() method"
	notice_NotSamePointer_ShouldPointer = "It should be a Pointer for NotSamePointer() method"

	gotFunc_Label = "Got func"
)
