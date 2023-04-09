# Test Actually

Yet another pithy testing framework `actually`.

Experimental yet :D

## Usage

    package you

    import (
    	"testing"

    	"github.com/bayashi/actually"
    )

    func Test(t *testing.T) {
        love, err := getLove()

    	actually.Got(love).True(t)
    	actually.Got(err).Nil(t)
    }

    func getLove() (bool, error) {
        return true, nil
    }

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
