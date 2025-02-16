# Actually

<a href="https://github.com/bayashi/actually/actions"><img src="https://github.com/bayashi/actually/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://goreportcard.com/report/github.com/bayashi/actually" title="actually report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/actually" alt="actually report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/actually" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/actually.svg" alt="Go Reference"></a>

A testing library focused on turning failure into success, `actually`.

* Builder interface to make test code obvious
* Consistent method name to reduce things you have to remember
* Specific fail report to save your time
* There are helpers to see details of a failure

## Usage

```go
package main

import (
	"testing"

	a "github.com/bayashi/actually"
	pb "github.com/bayashi/actually/testpb"
)

func TestObject(t *testing.T) {
	love, err := getLove()

	a.Got(err).NoError(t)
	a.Got(love).True(t)
}

func getLove() (bool, error) {
	return true, nil
}

func TestObjects(t *testing.T) {
	x := map[string]int{
		"foo": 123,
	}
	y := map[string]int{
		"foo": 123,
	}

	// `Same` method verifies that two objects are same in value and type.
	// Function type value is not acceptable. And not verify pointer address.
	// It will be fail, int(1) and uint(1), because of type.
	a.Got(x).Expect(y).Same(t)

	// Cmp method gets the differences between two objects by go-cmp.Diff.
	a.Got(x).Expect(y).Cmp(t)
}

func TestProtoMessages(t *testing.T) {
	x := &pb.Foo{Id: 123}
	y := &pb.Foo{Id: 123}

	// CmpProto method gets the differences between two Protobuf messages
	// by go-cmp.Diff with protocmp.Transform option.
	a.Got(x).Expect(y).CmpProto(t)

	a.Got(x).Expect(y).SamePointer(t) // This test will be failed
}
```

## Assertion Methods

### [For 1 object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-1-object)

* True, False, Nil, NotNil, NoError

### [For 2 objects](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-2-objects)

* Same, SamePointer, SameType, SameConvertibleNumber
* NotSame, NotSamePointer, NotSameType, NotSameConvertibleNumber
* Cmp, CmpProto, CmpAllowUnexported, CmpIgnoreUnexported, (CmpOpt)

### [For panic](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-panic)

* Panic, NoPanic, PanicMessage

### [For string value by regexp](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-string-value-by-regexp)

* Match, NotMatch

### [For length of an object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-length-of-an-object)

* Len

[Here is a Wiki of full API documentation](https://github.com/bayashi/actually/wiki).

-----

## Fail reports

`actually` will help you with evident fail report:

```go
package foo

import (
	"testing"

	a "github.com/bayashi/actually"
)

func Test(t *testing.T) {
	x := "foo\nbar\nbaz"
	y := "foo\nbar\nbug"

	a.Got(x).Expect(y).Same(t)
}
```

Above code will put fail report like below:

```
=== RUN   Test
    foo_test.go:13: Test
        Fail reason:    Not same value
        Type:           Expect:string, Got:string
        Expected:       "foo\nbar\nbug"
        Actually got:   "foo\nbar\nbaz"
        Diff details:   --- Expected
                        +++ Actually got
                        @@ -2,2 +2,2 @@
                         bar
                        -bug
                        +baz
        Trace:          /path/to/foo/foo_test.go:13
--- FAIL: Test (0.00s)
```

## Helper Methods

There are helper functions: Name, Skip, Fi, X, Diff, Dump, etc...

### Debug

`actually` has the `Debug("label", any_variable)` method to show additional data only in fail report.

Like below, `src` variable will be dumped nicely only on fail.

```go
res := someFunc(src)
actually.Got(res).Debug("src", src).True(t)
```

[See more helper functions](https://github.com/bayashi/actually/wiki/Helper-functions).


## ACTUALLY_TRACE_SOURCE

If you set true value (i.e. "1", "true" or "TRUE" etc) into ENV:`ACTUALLY_TRACE_SOURCE` on running a test, then you can see a piece of source code for each stack trace in fail report.

```
=== RUN   Test
    foo_test.go:13: Test
        Fail reason:    Not same value
        Type:           Expect:string, Got:string
        Expected:       "foo\nbar\nbug"
        Actually got:   "foo\nbar\nbaz"
        Diff details:   --- Expected
                        +++ Actually got
                        @@ -2,2 +2,2 @@
                         bar
                        -bug
                        +baz
        Trace:          /path/to/foo/foo_test.go:13
                         10     x := "foo\nbar\nbaz"
                         11     y := "foo\nbar\nbug"
                         12
                         13>    a.Got(x).Expect(y).Same(t)
                         14  }
--- FAIL: Test (0.00s)
```

-----

## Installation

    go get github.com/bayashi/actually

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi

## Special Thanks

Inspired by:

* https://github.com/stretchr/testify
* https://github.com/matryer/is
* https://github.com/fluentassert/verify
* https://metacpan.org/pod/Test::Arrow
