package witness

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bayashi/actually/witness/report"
	tu "github.com/bayashi/actually/witness/testutil"
)

// Global variables to check result
var res string
var failnow bool

func stub() {
	res = ""
	funcFail = func(t *testing.T, r *report.Failure) {
		res = r.Put()
		failnow = false
	}
	funcFailNow = func(t *testing.T, r *report.Failure) {
		funcFail(t, r)
		failnow = true
	}
}

func TestNil(t *testing.T) {
	stub()

	Got(nil).Expect(nil).Fail(t, "oops")

	// Fail reason:    oops
	// Expected:       <nil>
	// Actually got:   <nil>

	if ok, msg := tu.Match(`Expected:\s*\t<nil>`, res); !ok {
		t.Error(msg)
	}

	if ok, msg := tu.Match(`Actually got:\s*\t<nil>`, res); !ok {
		t.Error(msg)
	}
}

func TestError(t *testing.T) {
	stub()

	err := fmt.Errorf("error example %d", 123)
	Got(err).Fail(t, "oops")

	// Fail reason:    oops
	// Type:           Got:*errors.errorString
	// Actually got:   error example 123

	if ok, msg := tu.Match(`Type:\s*\tGot:\*errors\.errorString`, res); !ok {
		t.Error(msg)
	}

	if ok, msg := tu.Match(`Actually got:\s*\t\(\*errors\.errorString\)\([0-9a-fx]+\)\(error example 123\)`, res); !ok {
		t.Error(msg)
	}
}

func TestFailGot(t *testing.T) {
	stub()

	got := "got string"
	reason := "failure reason"
	Got(got).Name("Gotcha").Fail(t, reason)

	// Fail reason:    failure reason
	// Type:           Got:string
	// Actually got:   "got string"

	if !strings.HasPrefix(res, "TestFailGot/Gotcha\n") {
		t.Errorf("Expected to be the prefix string `TestFailGot/Gotcha`, but not: %q", res)
	}
	if !strings.Contains(res, "TestFailGot/Gotcha") {
		t.Errorf("Expected to be contained the string `Test name:`, but not: %q", res)
	}

	if !strings.Contains(res, "Trace:") {
		t.Errorf("Expected to be contained the string `Trace:`, but not: %q", res)
	}

	if !strings.Contains(res, "Fail reason:") {
		t.Errorf("Expected to be contained the string `Fail reason:`, but not: %q", res)
	}

	if !strings.Contains(res, reason) {
		t.Errorf("Expected to be contained the string `%s`, but not: %q", reason, res)
	}

	if !strings.Contains(res, "Got:string") {
		t.Errorf("Expected to be contained type, but not: %q", res)
	}

	if ok, msg := tu.Match(fmt.Sprintf("Actually got:\\s*\\t%q", got), res); !ok {
		t.Error(msg)
	}
}

func TestFailGotExpect(t *testing.T) {
	stub()

	got := "got string"
	expect := "expect string"
	reason := "failure reason"
	Got(got).Expect(expect).Fail(t, reason)

	// Fail reason:    failure reason
	// Expected:       "expect string"
	// Actually got:   "got string"

	if !strings.HasPrefix(res, "TestFailGotExpect\n") {
		t.Errorf("Expected to be the prefix string `TestFailGotExpect`, but not: %q", res)
	}
	if !strings.Contains(res, "Trace:") {
		t.Errorf("Expected to be contained the string `Trace:`, but not: %q", res)
	}

	if !strings.Contains(res, "Fail reason:") {
		t.Errorf("Expected to be contained the string `Fail reason:`, but not: %q", res)
	}

	if !strings.Contains(res, reason) {
		t.Errorf("Expected to be contained the string `%s`, but not: %q", reason, res)
	}

	if ok, msg := tu.Match(fmt.Sprintf("Actually got:\\s*\\t%q", got), res); !ok {
		t.Error(msg)
	}

	if ok, msg := tu.Match(fmt.Sprintf("Expected:\\s*\\t%q", expect), res); !ok {
		t.Error(msg)
	}
}

func TestFailWithDiff(t *testing.T) {
	stub()

	w := New(ShowDiff, NotShowRaw)

	got := "a\nb\nc"
	expect := "a\nd\nc"
	reason := "not same string"
	w.Got(got).Expect(expect).Fail(t, reason)

	// Fail reason:    not same string
	// Type:           Expect:string, Got:string
	// Expected:       "a\nd\nc"
	// Actually got:   "a\nb\nc"
	// Diff details:   --- Expected
	//                 +++ Actually got
	//                 @@ -1,3 +1,3 @@
	//                  a
	//                 -d
	//                 +b
	//                  c

	if !strings.Contains(res, "Diff details:") {
		t.Errorf("Expected to be contained the string `Diff details:`, but not: %q", res)
	}

	if ok, msg := tu.Match("\\s*a\n\\s*-d\n\\s*\\+b\n\\s*c", res); !ok {
		t.Error(msg)
	}
}

func TestFailWithRawData(t *testing.T) {
	stub()

	w := New(NotShowDiff, ShowRaw)

	got := "a\nb\nc"
	expect := "a\nd\nc"
	reason := "not same string"
	w.Got(got).Expect(expect).Fail(t, reason)

	// Fail reason:    not same string
	// Type:           Expect:string, Got:string
	// Expected:       "a\nd\nc"
	// Actually got:   "a\nb\nc"
	// Raw Expect:     ---
	//                 a
	//                 d
	//                 c
	//                 ---
	// Raw Got:        ---
	//                 a
	//                 b
	//                 c
	//                 ---

	if ok, msg := tu.Match("Raw Expect:\\s*\t---\n\\s*a\n\\s*d", res); !ok {
		t.Error(msg)
	}

	if ok, msg := tu.Match("Raw Got:\\s*\t---\n\\s*a\n\\s*b", res); !ok {
		t.Error(msg)
	}
}

func TestFailWithAdditionalMessage(t *testing.T) {
	stub()

	g := "a\nb\nc"
	e := "a\nd\nc"

	Got(g).Expect(e).Message("Example Label", "Some info").Fail(t, "Not same")

	// Fail reason:    Not same
	// Type:           Expect:string, Got:string
	// Expected:       "a\nd\nc"
	// Actually got:   "a\nb\nc"
	// Example Label:  Some info

	if ok, msg := tu.Match("Example Label:\\s*\tSome info\n", res); !ok {
		t.Error(msg)
	}
}

func TestFailWithDebugInfo(t *testing.T) {
	stub()

	Got(1).Expect(2).Debug("label", "debug info").Fail(t, "Not same")

	// Fail reason:  Not same
	// Type:         Expect:int, Got:int
	// Expected:     2
	// Actually got: 1
	// Debug label:  (string) (len=10) "debug info"

	if ok, msg := tu.Match(`Debug label:\s*\t\(string\) \(len=10\) "debug info"`, res); !ok {
		t.Error(msg)
	}
}

func TestFailWithMultipleDebugInfo(t *testing.T) {
	stub()

	Got(1).Expect(2).Debug("label", "debug info", "one more debug info").Fail(t, "Not same")

	// Fail reason:  Not same
	// Type:         Expect:int, Got:int
	// Expected:     2
	// Actually got: 1
	// Debug label:  (string) (len=10) "debug info"
	//               --
	//               (string) (len=19) "one more debug info"

	if ok, msg := tu.Match(`Debug label:[\t\s]+\(string\) \(len=10\) "debug info"\n[\t\s]+--\n[\t\s]+\(string\) \(len=19\) "one more debug info"`, res); !ok {
		t.Error(msg)
	}
}
