# Test Actually

<a href="https://github.com/bayashi/actually/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png"></a>
<a href="https://github.com/bayashi/actually/actions"><img src="https://github.com/bayashi/actually/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://pkg.go.dev/github.com/bayashi/actually"><img src="https://pkg.go.dev/badge/github.com/bayashi/actually.svg" alt="Go Reference"></a>

Yet another pithy testing framework `actually`.

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

Actually, you can write multiple assertions in one chain like below:

```go
package main

import (
	"testing"

	. "github.com/bayashi/actually"
)

func Test(t *testing.T) {
    love := true
    Got(love).NotNil(t).True(t).
        Expect(true).Same(t) // Obviously pass
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
