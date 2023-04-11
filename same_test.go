package actually_test

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestSame(t *testing.T) {
	// pass
	actually.Got(nil).Expect(nil).Same(t)
	actually.Got("").Expect("").Same(t)
	actually.Got("a").Expect("a").Same(t)
	actually.Got(0).Expect(0).Same(t)
	actually.Got(12).Expect(12).Same(t)
	actually.Got([2]int{1, 2}).Expect([2]int{1, 2}).Same(t)
	actually.Got([]string{}).Expect([]string{}).Same(t)
	actually.Got([]string{"a"}).Expect([]string{"a"}).Same(t)
	actually.Got(map[string]int{"foo":12}).Expect(map[string]int{"foo":12}).Same(t)

	foo := struct {
		bar string
	}{
		bar: "foo",
	}
	bar := struct {
		bar string
	}{
		bar: "foo",
	}
	actually.Got(foo).Expect(bar).Same(t)
	// Actually, pass above structs test, but these are not same address.
	// Just data are same. So, below test will fail
	//actually.Got(fmt.Sprintf("%p", &foo)).Expect(fmt.Sprintf("%p", &bar)).Same(t)

	// NOT SUPPORT chan YET
	// ch1 := make(chan string, 1)
	// ch1 <- "foo"
	// ch2 := make(chan string, 1)
	// ch2 <- "foo"
	// actually.Got(ch1).Expect(ch2).Same(t)

	// fail
	//actually.Got("a").Expect("b").Same(t)
}
