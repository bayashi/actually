package report

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/bayashi/actually/witness/obj"
)

const labelTag = "label"

type Failure struct {
	expect    *obj.Object `label:"Expected"`
	got       *obj.Object `label:"Actually got"`
	diff      string      `label:"Diff details"`
	rawExpect string      `label:"Raw Expect"`
	rawGot    string      `label:"Raw Got"`
	reason    string      `label:"Fail reason"`
	name      string      `label:"Test name"`
	trace     string      `label:"Trace"`
	messages  []map[string]string
	debugInfo []map[string][]*obj.Object `label:"Debug"`
}

func NewFailure() *Failure {
	return &Failure{}
}

func (f *Failure) Trace(trace string) *Failure {
	f.trace = trace
	return f
}

func (f *Failure) Name(name string) *Failure {
	f.name = name
	return f
}

func (f *Failure) Reason(reason string) *Failure {
	f.reason = reason
	return f
}

func (f *Failure) Got(o *obj.Object) *Failure {
	f.got = o
	return f
}

func (f *Failure) Expect(o *obj.Object) *Failure {
	f.expect = o
	return f
}

func (f *Failure) Diff(diff string) *Failure {
	f.diff = diff
	return f
}

func (f *Failure) RawGot(raw string) *Failure {
	f.rawGot = raw
	return f
}

func (f *Failure) RawExpect(raw string) *Failure {
	f.rawExpect = raw
	return f
}

func (f *Failure) Messages(msgs []map[string]string) *Failure {
	f.messages = msgs
	return f
}

func (f *Failure) DebugInfo(info []map[string][]*obj.Object) *Failure {
	f.debugInfo = info
	return f
}

func (f *Failure) Put() string {
	r := &Report{}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.Contents = f.buildContents()

	longestLen := 0
	for _, c := range r.Contents {
		if len(c.Label) > longestLen {
			longestLen = len(c.Label)
		}
	}

	lines := ""
	for i, c := range r.Contents {
		if i == 0 {
			lines += c.Body + "\n"
		} else {
			lines += fmt.Sprintf(
				"%s%s%s\t%s\n",
				c.Label,
				r.separator(),
				c.indentSpaces(longestLen),
				indentMessage(c.Body, longestLen),
			)
		}
	}

	return lines
}

func (f *Failure) buildTypeBody() string {
	var types []string
	if f.expect != nil && f.expect.Touch() {
		types = append(types, fmt.Sprintf("Expect:%s", f.expect.AsType()))
	}
	if f.got != nil && f.got.Touch() {
		types = append(types, fmt.Sprintf("Got:%s", f.got.AsType()))
	}

	return strings.Join(types, ", ")
}

func (f *Failure) buildContents() []*Content {
	var contents []*Content

	if f.name != "" {
		contents = append(contents, &Content{Label: f.fieldLabel("name"), Body: f.name})
	}
	if f.trace != "" {
		contents = append(contents, &Content{Label: f.fieldLabel("trace"), Body: f.trace})
	}
	if f.reason != "" {
		contents = append(contents, &Content{Label: f.fieldLabel("reason"), Body: f.reason})
	}

	if (f.expect != nil && f.expect.Touch()) || (f.got != nil && f.got.Touch()) {
		contents = append(contents, &Content{Label: "Type", Body: f.buildTypeBody()})
	}

	if f.expect != nil && f.expect.Touch() {
		contents = append(contents, &Content{Label: f.fieldLabel("expect"), Body: f.expect.AsString()})
	}
	if f.got != nil && f.got.Touch() {
		contents = append(contents, &Content{Label: f.fieldLabel("got"), Body: f.got.AsString()})
	}

	if f.diff != "" {
		contents = append(contents, &Content{Label: f.fieldLabel("diff"), Body: f.diff})
	}

	if f.rawExpect != "" {
		contents = append(contents, &Content{Label: "Raw Expect", Body: f.rawExpect})
	}
	if f.rawGot != "" {
		contents = append(contents, &Content{Label: "Raw Got", Body: f.rawGot})
	}

	for _, i := range f.messages {
		for label, body := range i {
			contents = append(contents, &Content{Label: label, Body: body})
		}
	}

	for _, i := range f.debugInfo {
		for label, objs := range i {
			label = strings.Join([]string{f.fieldLabel("debugInfo"), label}, " ")
			body := []string{}
			for _, o := range objs {
				body = append(body, o.AsDumpString())
			}
			contents = append(contents, &Content{Label: label, Body: strings.Join(body, "\n--\n")})
		}
	}

	return contents
}

func (f *Failure) fieldLabel(fieldName string) string {
	t := reflect.TypeOf(*f)
	field, _ := t.FieldByName(fieldName)

	return field.Tag.Get(labelTag)
}

type Report struct {
	Contents       []*Content
	LabelSeparator *string
	mutex          sync.RWMutex
}

func (r *Report) separator() string {
	if r.LabelSeparator != nil {
		return *r.LabelSeparator
	}

	return ":"
}

type Content struct {
	Label string
	Body  string
}

func (c *Content) indentSpaces(longestLen int) string {
	return strings.Repeat(" ", longestLen-len(c.Label))
}

func indentMessage(message string, len int) string {
	outBuf := new(bytes.Buffer)
	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		if i != 0 {
			outBuf.WriteString("\n\t" + strings.Repeat(" ", len+1) + "\t")
		}
		outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}
