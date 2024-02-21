# Actually

<a href="https://github.com/bayashi/actually/actions"><img src="https://github.com/bayashi/actually/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://goreportcard.com/report/github.com/bayashi/actually" title="actually report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/actually" alt="actually report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/actually" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/actually.svg" alt="Go Reference"></a>

Yet another assertion library, `actually`.

* Builder interface to make test code obvious
* Consistent method name to reduce things you have to remember
* Specific fail report to improve your Developer eXperience

-----

## Usage

[Try in playground](https://go.dev/play/p/Ut-hIr3vmYQ)

```go
package main

import (
    "testing"
    "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love, err := getLove()

    // Assert 1 object
    actually.Got(love).True(t)
    actually.Got(err).NoError(t)

    // Assert 2 objects
    actually.Got(love).Expect(true).Same(t)
    actually.Got(int32(1)).Expect(float64(1.0)).SameNumber(t)

    heart := &love
    body  := heart
    actually.Got(heart).Expect(body).SamePointer(t)
}

func getLove() (bool, error) {
    return true, nil
}
```

NOTE that `Got()` and `Expect()` should NOT be called multiple times in one chain.

-----

## Assertion Methods

### [Assertion for 1 object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-1-object)

* True
* False
* Nil
* NotNil
* NoError

### [Assertion for 2 objects](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-2-objects)

* Same
* SameNumber
* SamePointer
* SameType
* NotSame
* NotSameNumber
* NotSamePointer
* NotSameType

### [Assertion for string value by regexp](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-string-value-by-regexp)

* Match
* NotMatch

### [Assertion for length of an object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-length-of-an-object)

* Len

[Here is a Wiki of full API documentation](https://github.com/bayashi/actually/wiki).

-----

## Fail reports

Test code often breaks.

We often end up spending valuable time fixing failed tests written by someone who is no longer with us.

`actually` will help you with specific fail report:

```
nil_test.go:28:
            Test name:      TestNil
            Trace:          /path/to/src/github.com/bayashi/actually/nil_test.go:28
            Fail reason:    Expected <nil>, but it was NOT <nil>
            Type:           Got:string
            Actually got:   ""
```

Another example with diff:

```
same_test.go:19:
            Test name:      TestSame
            Trace:          /path/to/src/github.com/bayashi/actually/same_test.go:19
            Fail reason:    Not same value
            Expected:       Type: map[string]int, Dump: map[string]int{"foo":12}
            Actually got:   Type: map[string]int, Dump: map[string]int{"joo":12}
            Diff Details:   --- Expected
                            +++ Actually got
                            @@ -1,3 +1,3 @@
                             (map[string]int) (len=1) {
                            - (string) (len=3) "foo": (int) 12
                            + (string) (len=3) "joo": (int) 12
                             }
```

`actually` has a `X()` method to show raw strings in the fail report. It would be helpful to compare intricate strigns, like below. You don't need to keep commented-out lines to dump test data anymore:

```go
actually.Got(stringA).Expect(stringB).X().Same(t)
```

Below report would be lovely.

```
builder_test.go:133:
            Test name:      TestTree
            Trace:          /path/to/src/github.com/bayashi/goverview/builder_test.go:133
            Fail reason:    Not same
            Expected:       Dump: "\n┌ 001/\n├── .gitignore\n├── LICENSE: License MIT\n├── go.mod: go 1.18\n└───+ main.go: main\n      Func: X\n      const: X\n"
            Actually got:   Dump: "\n┌ 001/\n├── .gitignore\n├── LICENSE: License MIT\n├── go.mod: go 1.19\n└──* main.go: main\n      Func: X\n      Const: X\n"
            Diff Details:   --- Expected
                            +++ Actually got
                            @@ -4,6 +4,6 @@
                             ├── LICENSE: License MIT
                            -├── go.mod: go 1.18
                            -└───+ main.go: main
                            +├── go.mod: go 1.19
                            +└──* main.go: main
                                   Func: X
                            -      const: X
                            +      Const: X
            Raw Expect:     ---
                            ┌ 001/
                            ├── .gitignore
                            ├── LICENSE: License MIT
                            ├── go.mod: go 1.18
                            └───+ main.go: main
                                  Func: X
                                  const: X
                            ---
            Raw Got:        ---
                            ┌ 001/
                            ├── .gitignore
                            ├── LICENSE: License MIT
                            ├── go.mod: go 1.19
                            └──* main.go: main
                                  Func: X
                                  Const: X
```

There would be a notice message with a fail reason as a hint to pass:

```
same_test.go:64:
            Test name:      TestSamePointer
            Trace:          /path/to/src/github.com/bayashi/actually/same.go:53
                                    /path/to/src/github.com/bayashi/actually/same_test.go:64
            Fail reason:    `Got` is NOT type of Pointer
            Expected:       Type: *int, Dump: (*int)(0xc00001a528)
            Actually got:   Type: string, Dump: ""
            Notice:         It should be a Pointer for SamePointer() method
```

`actually` has a `Diff` method to see differences between 2 objects for debugging.

```go
Println(actually.Diff(objA, objB))
```

If objects are not string, even if these are objects, you can see the differences of dumped data.

[See more details in a Wiki](https://github.com/bayashi/actually/wiki).

-----

## Installation

    go get github.com/bayashi/actually

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi

## See Also

* https://github.com/bayashi/witness

## Special Thanks

Inspired by:

* https://github.com/stretchr/testify
* https://github.com/matryer/is
* https://github.com/fluentassert/verify
* https://metacpan.org/pod/Test::Arrow
