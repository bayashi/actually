# Test Actually

Yet another pithy testing framework `actually`.

Experimental yet :D

## Usage

```go
package you

import (
    "testing"
    "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love, err := getLove()
    actually.Got(love).True(t)
    actually.Got(err).Nil(t)
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

This is an example of fail case:

```
true_test.go:21:
            Trace:          /go/src/github.com/bayashi/actually/true.go:27
                                    /go/src/github.com/bayashi/actually/true_test.go:21
            Test func:      TestTrue()
            Expected:       true
            Actually got:   false
```

Another example:

```
same_test.go:69:
            Trace:          /go/src/github.com/bayashi/actually/same.go:99
                                    /go/src/github.com/bayashi/actually/same_test.go:69
            Test func:      TestSameNumber()
            Fail reason:    The types of `Got` and `Expect` are NOT convertible
            Expected:       Type:int, 1
            Actually got:   Type:string, "1"
```

Actually, you can write multiple assertions in one chain as below:

```go
package main

import (
	"testing"

	. "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love := true
    Got(love).NotNil(t).True(t).
        Expect(true).Same(t)
}
```

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
