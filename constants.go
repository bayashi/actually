package actually

const (
	TraceSeparator = "\n\t"

	FailReason_WrongType           = "Wrong Type"
	FailReason_NotSame             = "Not same"
	FailReason_WrongPointerAddress = "Wrong Type"
	FailReason_GotIsFunc           = "`Got` value is type of function. It cannot be used in Same() method"
	FailReason_ExpectIsFunc        = "`Expect` value is type of function. It cannot be used in Same() method"
	FailReason_GotIsNotPointer     = "`Got` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
	FailReason_ExpectIsNotPointer  = "`Expect` is NOT type of Pointer. It should be a Pointer for SamePointer() method"
)
