package report

// Most of the logic was copied from https://github.com/stretchr/testify

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type Report struct {
	trace   string `label:"Trace"`
	name    string `label:"Test func"`
	reason  string `label:"Fail reason"`
	expect  string `label:"Expected"`
	got     string `label:"Actually got"`
	diff    string `label:"Diff"`
	message string `label:"Message"`
}

type reportContent struct {
	label   string
	content string
}

const reportLabelTag = "label"

func New() *Report {
	return &Report{}
}

func (r *Report) Put() string {
	reportContents := r.buildReportContets()

	longestLen := 0
	for _, rc := range *reportContents {
		if len(rc.label) > longestLen {
			longestLen = len(rc.label)
		}
	}

	output := ""
	for _, rc := range *reportContents {
		label := fmt.Sprintf("%s:%s", rc.label, strings.Repeat(" ", longestLen - len(rc.label)))
		output += fmt.Sprintf("\t%s\t%s\n", label, indentMessage(rc.content, longestLen))
	}

	return output
}

func (r *Report) buildReportContets() *[]*reportContent {
	var rContents []*reportContent

	if r.trace != "" {
		rContents = append(rContents, &reportContent{label: r.label("trace"), content: r.trace})
	}
	if r.name != "" {
		rContents = append(rContents, &reportContent{label: r.label("name"), content: r.name})
	}
	if r.reason != "" {
		rContents = append(rContents, &reportContent{label: r.label("reason"), content: r.reason})
	}
	if r.expect != "" {
		rContents = append(rContents, &reportContent{label: r.label("expect"), content: r.expect})
	}
	if r.got != "" {
		rContents = append(rContents, &reportContent{label: r.label("got"), content: r.got})
	}
	if r.diff != "" {
		rContents = append(rContents, &reportContent{label: r.label("diff"), content: r.diff})
	}
	if r.message != "" {
		rContents = append(rContents, &reportContent{label: r.label("message"), content: r.message})
	}

	return &rContents
}

func (r *Report) label(field string) string {
	t := reflect.TypeOf(*r)
	f, _ := t.FieldByName(field)

	return f.Tag.Get(reportLabelTag)
}

func indentMessage(message string, longestLen int) string {
	outBuf := new(bytes.Buffer)
	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		if i != 0 {
			outBuf.WriteString("\n\t" + strings.Repeat(" ", longestLen + 1) + "\t")
		}
		outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}

func (r *Report) Trace(trace string) *Report {
	r.trace = trace
	return r
}

func (r *Report) Name(name string) *Report {
	r.name = name
	return r
}

func (r *Report) Reason(reason string) *Report {
	r.reason = reason
	return r
}

func (r *Report) Got(got string) *Report {
	r.got = got
	return r
}

func (r *Report) Gotf(format string, vars ...any) *Report {
	r.got = fmt.Sprintf(format, vars...)
	return r
}

func (r *Report) Expect(expect string) *Report {
	r.expect = expect
	return r
}

func (r *Report) Expectf(format string, vars ...any) *Report {
	r.expect = fmt.Sprintf(format, vars...)
	return r
}

func (r *Report) Diff(diff string) *Report {
	r.diff = diff
	return r
}

func (r *Report) Message(message string) *Report {
	r.message = message
	return r
}
