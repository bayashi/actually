package actually

const (
	envKey_FailNow = "ACTUALLY_FAIL_NOW"

	traceSeparator = "\n\t"

	panicReason_CalledGotTwice    = "Don't call `Got` method twice. It has been called already"
	panicReason_CalledExpectTwice = "Don't call `Expect` method twice. It has been called already"

	message_ExpectTrue  = "Boolean type <true>"
	message_ExpectFalse = "Boolean type <false>"

	// fail reason
	reason_WrongType             = "Different type"
	reason_NotSame               = "Not same value"
	reason_WrongPointerAddress   = "Wrong pointer address"
	reason_GotIsFunc             = "`Got` value is type of function"
	reason_ExpectIsFunc          = "`Expect` value is type of function"
	reason_GotIsNotPointer       = "`Got` is NOT type of Pointer"
	reason_ExpectIsNotPointer    = "`Expect` is NOT type of Pointer"
	reason_GotIsNilType          = "Type of `Got` is a <nil> value"
	reason_ExpectIsNilType       = "Type of `Expect` is a <nil> value"
	reason_ExpectIsNotNil        = "Expected other than <nil>, but got <nil>"
	reason_ExpectIsNotValidValue = "`Expect` value is NOT a valid value"
	reason_NotConvertibleTypes   = "The types of `Got` and `Expect` are NOT convertible"

	// notice_Method_*
	notice_Same_NotAcceptable        = "It's not acceptable in Same() method"
	notice_SamePointer_ShouldPointer = "It should be a Pointer for SamePointer() method"
	notice_SameNumber_ShouldNumber   = "It should be a number for SameNumber() method"

	template_Dump            = "Type: %Y, Dump: %#v"
	template_DumpStringType  = "Dump: %#v"
	template_DumpAsRawString = "---\n%s\n---"
)
