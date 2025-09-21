package source

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ContentLineNumberBefore = 3
	ContentLineNumberAfter  = 3
)

func GetSource(filepath string, lineNumber int) ([]string, error) {
	lineNumberStart := lineNumber - ContentLineNumberBefore
	if lineNumberStart < 1 {
		lineNumberStart = 1
	}
	lineNumberEnd := lineNumber + ContentLineNumberAfter

	fh, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	if _, err := fh.Seek(0, 0); err != nil {
		return nil, err
	}

	source := []string{}
	count := 0
	digit := len(strconv.Itoa(lineNumberEnd))
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		count++
		if count < lineNumberStart {
			continue
		}
		if count > lineNumberEnd {
			break
		}
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		source = append(source, fmt.Sprintf(
			"%s%d%s %s",
			strings.Repeat("0", digit-len(strconv.Itoa(count))),
			count,
			delimitor(lineNumber, count),
			line,
		))
	}

	return source, nil
}

func delimitor(lineNumber int, count int) string {
	if lineNumber == count {
		return ">"
	}

	return " "
}
