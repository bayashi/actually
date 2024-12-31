package witness

import (
	"testing"
)

// Set a value you actually got. The `Actual` method is an alias of the `Got` method.
func Actual(v any) *Witness {
	return Got(v)
}

// Set a value you actually got. The `Actual` method is an alias of the `Got` method.
func (w *Witness) Actual(v any) *Witness {
	return w.Got(v)
}

// Set a value you want. The `Want` method is an alias of the `Expect` method.
func Want(v any) *Witness {
	return Expect(v)
}

// Set a value you want. The `Want` method is an alias of the `Expect` method.
func (w *Witness) Want(v any) *Witness {
	return w.Expect(v)
}

// Do fail with report and stop running test right now. The `Fatal` method is an alias of the `FailNow` method.
func (w *Witness) Fatal(t *testing.T, reason string) {
	t.Helper()
	w.FailNow(t, reason)
}

// Do fail with report and stop running test right now. The `Fatal` method is an alias of the `FailNow` method.
func Fatal(t *testing.T, reason string, got any, expect ...any) {
	t.Helper()
	FailNow(t, reason, got, expect...)
}
