package actually

const (
	TraceSeparator = "\n\t"

	FailReason_WrongType             = "Wrong Type"
	FailReason_NotSame               = "Not same"
	FailReason_WrongPointerAddress   = "Wrong Type"
	FailReason_GotIsFunc             = "`Got` value is type of function. It cannot be used in Same() method"
	FailReason_ExpectIsFunc          = "`Expect` value is type of function. It cannot be used in Same() method"
	FailReason_GotIsNotPointer       = "`Got` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
	FailReason_ExpectIsNotPointer    = "`Expect` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
	FailReason_GotIsNilType          = "Type of `Got` is a nil value"
	FailReason_ExpectIsNilType       = "Type of `Expect` is a nil value"
	FailReason_ExpectIsNotValidValue = "`Expect` value is NOT a valid value"
	FailReason_NotConvertibleTypes   = "The types of `Got` and `Expect` are NOT convertible"
)
