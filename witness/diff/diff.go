package diff

import (
	"reflect"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
)

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
	MaxDepth:                10,
}

var spewConfigStringerEnabled = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	MaxDepth:                10,
}

// Diff method returns diff text of 2 testing objects
func Diff(expectv any, gotv any) string {
	e, g := pre(expectv, gotv)
	if e == "" && g == "" {
		return ""
	}

	return getDiff(e, g, "Expected", "Actually got")
}

// DiffSimple method returns diff text of 2 objects as "a and b"
func DiffSimple(av any, bv any) string {
	a, b := pre(av, bv)
	if a == "" && b == "" {
		return ""
	}

	return getDiff(a, b, "a", "b")
}

func pre(expectv any, gotv any) (string, string) {
	if expectv == nil || gotv == nil {
		return "", ""
	}

	et, ek := typeAndKind(expectv)
	gt, _ := typeAndKind(gotv)

	if et != gt {
		return "", ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return "", ""
	}

	var e, g string

	switch et {
	case reflect.TypeOf(""):
		e = reflect.ValueOf(expectv).String()
		g = reflect.ValueOf(gotv).String()
	case reflect.TypeOf(time.Time{}):
		e = spewConfigStringerEnabled.Sdump(expectv)
		g = spewConfigStringerEnabled.Sdump(gotv)
	default:
		e = spewConfig.Sdump(expectv)
		g = spewConfig.Sdump(gotv)
	}

	return e, g
}

func typeAndKind(v any) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}

func getDiff(e string, g string, fromFile string, toFile string) string {
	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(g),
		FromFile: fromFile,
		FromDate: "",
		ToFile:   toFile,
		ToDate:   "",
		Context:  1,
	})

	return diff
}
