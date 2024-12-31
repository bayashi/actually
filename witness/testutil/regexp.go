package testutil

import (
	"fmt"
	"regexp"
)

func Match(re string, target string) (bool, string) {
	reg := regexp.MustCompile(re)
	if !reg.MatchString(target) {
		return false, fmt.Sprintf("Not matched the regexp `%s` for %q", reg.String(), target)
	}

	return true, ""
}
