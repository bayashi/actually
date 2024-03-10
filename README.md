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
    a "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love, err := getLove()

    // Assert 1 object
    a.Got(love).True(t)
    a.Got(err).NoError(t)

    // Assert 2 objects
    a.Got(love).Expect(true).Same(t)
    a.Got(int32(1)).Expect(float64(1.0)).SameNumber(t)

    heart := &love
    body  := heart
    a.Got(heart).Expect(body).SamePointer(t)
}

func getLove() (bool, error) {
    return true, nil
}
```

NOTE that `Got()` and `Expect()` should NOT be called multiple times in one chain.

-----

## Assertion Methods

### [Assertion for 1 object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-1-object)

* True, False, Nil, NotNil, NoError

### [Assertion for 2 objects](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-2-objects)

* Same, SameNumber, SamePointer, SameType, NotSame, NotSameNumber, NotSamePointer, NotSameType

### [Assertion for panic](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-panic)

* Panic, PanicMessage, NoPanic

### [Assertion for string value by regexp](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-string-value-by-regexp)

* Match, NotMatch

### [Assertion for length of an object](https://github.com/bayashi/actually/wiki/All-assertion-methods#assertion-for-length-of-an-object)

* Len

[Here is a Wiki of full API documentation](https://github.com/bayashi/actually/wiki).

-----

## Fail reports

Test code often breaks.

We often end up spending valuable time fixing failed tests written by someone who is no longer with us.

`actually` will help you with evident fail report:

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

`actually` has the `Debug` method to show additional data in fail report.

Like below, `src` variable will be dumped nicely with Got value `res` on fail.

```go
res := someFunc(src)
actually.Got(res).Debug(src).True(t)
```

There are other helper methods too.

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
