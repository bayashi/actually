package actually

import "fmt"

type failReport struct {
	message string
	got any
	expect any
	diff string
}

func (fr *failReport) ToString() string {
	return ""
}

func (fr *failReport) gotAsTypeString() string {
	return fmt.Sprintf("%T", fr.got)
}

func (fr *failReport) expectAsTypeString() string {
	return fmt.Sprintf("%T", fr.expect)
}

func (fr *failReport) gotAsLiteralString() string {
	return fmt.Sprintf("%#v", fr.got)
}

func (fr *failReport) expectAsLiteralString() string {
	return fmt.Sprintf("%#v", fr.expect)
}