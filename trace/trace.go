// trace package provides a function to get caller info for a testing tool
package trace

// The functions in this file were copied from github.com/stretchr/testify

import (
	"fmt"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Info(skipTraceRule func(filepath string) bool) []string {
	var pc uintptr
	var file string
	var line int
	var ok bool

	var funcName string
	var lastCaller string

	callers := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}

		// Avoid panic. https://github.com/stretchr/testify/issues/180
		if file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}

		funcName = f.Name()

		// testing.tRunner is the standard library function that calls
		// tests. Subtests are called directly by tRunner, without going through
		// the Test/Benchmark/Example function that contains the t.Run calls.
		if funcName == "testing.tRunner" {
			break
		}

		lastCaller = fmt.Sprintf("%s:%d", file, line)

		if len(strings.Split(file, "/")) > 1 && // https://github.com/stretchr/testify/pull/402
			!skipTraceRule(file) && !strings.Contains(file, "/actually") {
			callers = append(callers, lastCaller)
		}

		segments := strings.Split(funcName, ".")
		name := segments[len(segments)-1]
		if isGolangTestFunc(name) {
			break
		}
	}

	if len(callers) == 0 {
		callers = append(callers, lastCaller)
	}

	return callers
}

func isGolangTestFunc(name string) bool {
	for _, prefix := range [3]string{"Test", "Benchmark", "Example"} {
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		if len(name) == len(prefix) {
			return true
		}
		r, _ := utf8.DecodeRuneInString(name[len(prefix):])
		if !unicode.IsLower(r) {
			return true
		}
	}

	return false
}
