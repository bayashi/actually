# Actually

<a href="https://github.com/bayashi/actually/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png"></a>
<a href="https://github.com/bayashi/actually/actions"><img src="https://github.com/bayashi/actually/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://goreportcard.com/report/github.com/bayashi/actually" title="actually report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/actually" alt="actually report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/actually" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/actually.svg" alt="Go Reference"></a>

Yet another pithy testing framework, `actually`.

## Usage

This is an example code of tests in `actually`. ([Try in playground](https://go.dev/play/p/Ut-hIr3vmYQ))

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

You can write multiple assertions in one chain like below ([Try in playground](https://go.dev/play/p/GxCV1Ubg6Uo)):

```go
package main

import (
	"testing"

	a "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love := true
    a.Got(love).NotNil(t).True(t).
        Expect(true).Same(t) // Obviously pass
}
```

NOTE that `Got()` and `Expect()` should NOT be called multiple times in one chain.

[Here is a Wiki of full API documentation](https://github.com/bayashi/actually/wiki).

## Fail reports

We spend a lot of time dealing with failing tests, so `actually` can shorten that time.

This is an example of simple fail report:

```
nil_test.go:28:
            Trace:          /path/to/src/github.com/bayashi/actually/nil_test.go:28
            Function:       TestNil()
            Expected:       <nil>
            Actually got:   Type:string, ""
```

Another example with diff:

```
same_test.go:19:
            Trace:          /path/to/src/github.com/bayashi/actually/same_test.go:19
            Function:       TestSame()
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
            Trace:          /path/to/src/github.com/bayashi/goverview/builder_test.go:133
            Function:       TestTree()
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
            Expected Raw:   ---
                            ┌ 001/
                            ├── .gitignore
                            ├── LICENSE: License MIT
                            ├── go.mod: go 1.18
                            └───+ main.go: main
                                  Func: X
                                  const: X
                            ---
            Got Raw:        ---
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
            Trace:          /path/to/src/github.com/bayashi/actually/same.go:53
                                    /path/to/src/github.com/bayashi/actually/same_test.go:64
            Function:       TestSamePointer()
            Fail reason:    `Got` is NOT type of Pointer
            Notice:         It should be a Pointer for SamePointer() method
            Expected:       Type: *int, Dump: (*int)(0xc00001a528)
            Actually got:   Type: string, Dump: ""
```

`actually` has a `Diff` method to see differences between 2 objects.

```go
Println(actually.Diff(objA, objB))
```

If objects are not string, even if these are objects, you can see the differences of dumped data.

[See more details in a Wiki](https://github.com/bayashi/actually/wiki).

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
