package actually

const (
	traceSeparator = "\n\t"

	panicReason_CalledGotTwice    = "Don't call `Got` method twice. It has been called already"
	panicReason_CalledExpectTwice = "Don't call `Expect` method twice. It has been called already"

	message_ExpectTrue  = "Boolean type <true>"
	message_ExpectFalse = "Boolean type <false>"

	failReason_WrongType             = "Each value you set `Got` and `Expect` has a different type"
	failReason_NotSame               = "Not same value"
	failReason_WrongPointerAddress   = "Wrong Pointer"
	failReason_GotIsFunc             = "`Got` value is type of function"
	failReason_ExpectIsFunc          = "`Expect` value is type of function"
	failReason_GotIsNotPointer       = "`Got` is NOT type of Pointer"
	failReason_ExpectIsNotPointer    = "`Expect` is NOT type of Pointer"
	failReason_GotIsNilType          = "Type of `Got` is a <nil> value"
	failReason_ExpectIsNilType       = "Type of `Expect` is a <nil> value"
	failReason_ExpectIsNotNil        = "Expected other than <nil>, but got <nil>"
	failReason_ExpectIsNotValidValue = "`Expect` value is NOT a valid value"
	failReason_NotConvertibleTypes   = "The types of `Got` and `Expect` are NOT convertible"

	failNotice_NotAcceptableSameMethod  = "It's not acceptable in Same() method"
	failNotice_ShouldPointerSamePointer = "It should be a Pointer for SamePointer() method"
	failNotice_ShouldNumberSameNumber   = "It should be a number for SameNumber() method"

	template_Dump            = "Type: %Y, Dump: %#v"
	template_DumpStringType  = "Dump: %#v"
	template_DumpAsRawString = "---\n%s\n---"
)
