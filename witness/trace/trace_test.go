package trace

import (
	"testing"

	tu "github.com/bayashi/actually/witness/testutil"
)

func TestInfo(t *testing.T) {
	trace := Info()
	if len(trace) != 2 {
		t.Errorf("trace length should be 2. but %#v", trace)
	}

	if ok, msg := tu.Match(`/witness/trace/trace\.go:\d+$`, trace[0]); !ok {
		t.Error(msg)
	}

	if ok, msg := tu.Match(`/witness/trace/trace_test\.go:\d+$`, trace[1]); !ok {
		t.Error(msg)
	}
}

func TestInfoWithFilter(t *testing.T) {
	filterFunc := func(file string) bool {
		ok, msg := tu.Match(`/witness/trace/trace`, file)
		if msg != "" {
			t.Error(msg)
		}
		return ok
	}
	trace := Info(filterFunc)
	if len(trace) != 1 {
		t.Errorf("trace length should be 1. but %d, %#v", len(trace), trace)
	}
}
