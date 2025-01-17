package actually

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/testing/protocmp"
)

// Cmp method gets the differences between two objects by go-cmp.Diff.
// https://pkg.go.dev/github.com/google/go-cmp/cmp#Diff
/*
	actually.Got(obj1).Expect(obj2).Cmp(t)
*/
// If you need to set cmp.Option, then you shoud use `CmpOpt(cmp.Option)` method before calling Cmp.
// Cmp method is just a wrapper of go-cmp.Diff. So, it's same that unexported fields are not compared by default;
// they result in panics unless suppressed by using an Ignore option. It may panic if it cannot compare the values.
func (a *testingA) Cmp(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if diff := cmp.Diff(a.expect, a.got, a.cmpOpts.cmpOpts...); diff != "" {
		return a.fail(reportForSame(a).Message("Diff details", diff), reason_NotSame)
	}

	return a
}

// CmpProto method gets the differences between two Protobuf messages by go-cmp.Diff with protocmp.Transform option.
/*
	actually.Got(protoMessage1).Expect(protoMessage2).CmpProto(t)
*/
func (a *testingA) CmpProto(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	a = a.CmpOpt(protocmp.Transform())

	if diff := cmp.Diff(a.expect, a.got, a.cmpOpts.cmpOpts...); diff != "" {
		return a.fail(reportForSame(a).Message("Diff details", diff), reason_NotSame)
	}

	return a
}

// CmpAllowUnexported method gets the differences between two structs by go-cmp.Diff with cmp.AllowUnexported option.
// It accepts unexported methods to compare instead panic. If you would like to ignore unexported methods,
// then you can use CmpIgnoreUnexported method.
func (a *testingA) CmpAllowUnexported(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !isStructType(a.got) {
		w := reportForSame(a).Message(notice_Label, notice_Cmp_ShouldStruct)
		return a.fail(w, reason_GotShouldStruct)
	}
	if !isStructType(a.expect) {
		w := reportForSame(a).Message(notice_Label, notice_Cmp_ShouldStruct)
		return a.fail(w, reason_ExpectShouldStruct)
	}

	a.CmpOpt(cmp.AllowUnexported(a.got, a.expect))

	if diff := cmp.Diff(a.expect, a.got, a.cmpOpts.cmpOpts...); diff != "" {
		return a.fail(reportForSame(a).Message("Diff details", diff), reason_NotSame)
	}

	return a
}

// CmpIgnoreUnexported method gets the differences between two structs by go-cmp.Diff with cmpopts.IgnoreUnexported option.
// It ignores unexported methods to compare instead panic. If you would like to compare with also unexported methods,
// then you can use CmpAllowUnexported method.
func (a *testingA) CmpIgnoreUnexported(t *testing.T, testNames ...string) *testingA {
	invalidCallForSame(a)
	a.name = a.naming(testNames...)
	a.t = t
	a.t.Helper()

	if !isStructType(a.got) {
		w := reportForSame(a).Message(notice_Label, notice_Cmp_ShouldStruct)
		return a.fail(w, reason_GotShouldStruct)
	}
	if !isStructType(a.expect) {
		w := reportForSame(a).Message(notice_Label, notice_Cmp_ShouldStruct)
		return a.fail(w, reason_ExpectShouldStruct)
	}

	a.CmpOpt(cmpopts.IgnoreUnexported(a.got, a.expect))

	if diff := cmp.Diff(a.expect, a.got, a.cmpOpts.cmpOpts...); diff != "" {
		return a.fail(reportForSame(a).Message("Diff details", diff), reason_NotSame)
	}

	return a
}

// CmpOpt method sets/adds options for Cmp* methods.
// There is no method to reset cmpOpts. Just set all opts at one time, or add opts.
/*
	actually.Got(obj1).Expect(obj2).CmpOpt(cmpopts.IgnoreFields(Foo{}, "Field")).Cmp(t)
*/
// ref:
// * https://pkg.go.dev/github.com/google/go-cmp/cmp#Option
// * https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts
func (a *testingA) CmpOpt(cmpOpts ...cmp.Option) *testingA {
	a.cmpOpts.mutex.Lock()
	defer a.cmpOpts.mutex.Unlock()
	a.cmpOpts.cmpOpts = append(a.cmpOpts.cmpOpts, cmpOpts...)

	return a
}
