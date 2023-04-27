// diff package provides a utility to handle diff data for a testing tool
package diff

// Most of the logic was copied from https://github.com/stretchr/testify

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

func Diff(expectv any, gotv any) string {
	if expectv == nil || gotv == nil {
		return ""
	}

	et, ek := typeAndKind(expectv)
	gt, _ := typeAndKind(gotv)

	if et != gt {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
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

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(g),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actually got",
		ToDate:   "",
		Context:  1,
	})

	return diff
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
