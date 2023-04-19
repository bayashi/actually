package actually

const (
	traceSeparator = "\n\t"

	panicReason_CalledGotTwice    = "Don't call `Got` method twice. It has been called already"
	panicReason_CalledExpectTwice = "Don't call `Expect` method twice. It has been called already"

	failReason_WrongType             = "Wrong Type"
	failReason_NotSame               = "Not same"
	failReason_WrongPointerAddress   = "Wrong Type"
	failReason_GotIsFunc             = "`Got` value is type of function. It cannot be used in Same() method"
	failReason_ExpectIsFunc          = "`Expect` value is type of function. It cannot be used in Same() method"
	failReason_GotIsNotPointer       = "`Got` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
	failReason_ExpectIsNotPointer    = "`Expect` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
	failReason_GotIsNilType          = "Type of `Got` is a nil value"
	failReason_ExpectIsNilType       = "Type of `Expect` is a nil value"
	failReason_ExpectIsNotValidValue = "`Expect` value is NOT a valid value"
	failReason_NotConvertibleTypes   = "The types of `Got` and `Expect` are NOT convertible"

	template_Dump             = "Type: %Y\nDump: %#v"
	template_DumpStringType   = "Dump: %#v"
	template_DumpAsStringAlso = "---\n%s\n---"
)
