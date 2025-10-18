package report

import (
	"testing"

	"github.com/bayashi/actually/witness/obj"
)

func TestReport(t *testing.T) {
	f := NewFailure()
	f.Trace("trace").Diff("diff")
	got := f.Put()
	expect := "\nDiff details:\tdiff\nTrace:       \ttrace"
	if expect != got {
		t.Errorf("Report is wrong.\n[Expect]\n%#v\n\n[Got]\n%#v\n", expect, got)
		t.Fail()
	}
}

func TestSetter(t *testing.T) {
	f := NewFailure()

	f = f.Trace("t")
	if f.trace != "t" {
		t.Errorf("Wrong Trace(): %#v", f.trace)
	}

	f = f.Name("n")
	if f.name != "n" {
		t.Errorf("Wrong Name(): %#v", f.name)
	}

	f = f.Reason("r")
	if f.reason != "r" {
		t.Errorf("Wrong Reason(): %#v", f.reason)
	}

	f = f.Got(obj.NewObject("g"))
	if f.got.AsRawValue() != "g" {
		t.Errorf("Wrong Got(): %#v", f.got)
	}
	f = f.RawGot("gs")
	if f.rawGot != "gs" {
		t.Errorf("Wrong GotAsRaw(): %#v", f.rawGot)
	}

	f = f.Expect(obj.NewObject("e"))
	if f.expect.AsRawValue() != "e" {
		t.Errorf("Wrong Expect(): %#v", f.expect)
	}
	f = f.RawExpect("es")
	if f.rawExpect != "es" {
		t.Errorf("Wrong ExpectAsRaw(): %#v", f.rawExpect)
	}

	f = f.Diff("d")
	if f.diff != "d" {
		t.Errorf("Wrong Diff(): %#v", f.diff)
	}
}

func TestBuildTypeBody(t *testing.T) {
	for name, tt := range map[string]struct {
		expect   *obj.Object
		got      *obj.Object
		expected string
	}{
		"different types": {
			expect:   obj.NewObject("hello"),
			got:      obj.NewObject(123),
			expected: "Expect:string, Got:int",
		},
		"expect only": {
			expect:   obj.NewObject("hello"),
			got:      nil,
			expected: "Expect:string",
		},
		"got only": {
			expect:   nil,
			got:      obj.NewObject(123),
			expected: "Got:int",
		},
		"both nil": {
			expect:   nil,
			got:      nil,
			expected: "",
		},
		"same type": {
			expect:   obj.NewObject("hello"),
			got:      obj.NewObject("world"),
			expected: "Expect:string, Got:string",
		},
		"pointer types": {
			expect:   obj.NewObject(&[]int{1, 2, 3}),
			got:      obj.NewObject(&[]string{"a", "b"}),
			expected: "Expect:*[]int, Got:*[]string",
		},
		"struct types": {
			expect:   obj.NewObject(struct{ ID int }{ID: 1}),
			got:      obj.NewObject(struct{ Name string }{Name: "test"}),
			expected: "Expect:struct { ID int }, Got:struct { Name string }",
		},
		"nil values with touch true": {
			expect:   obj.NewObject(nil),
			got:      obj.NewObject(nil),
			expected: "Expect:<nil>, Got:<nil>",
		},
		"expect only with nil value": {
			expect:   obj.NewObject(nil),
			got:      nil,
			expected: "Expect:<nil>",
		},
		"got only with nil value": {
			expect:   nil,
			got:      obj.NewObject(nil),
			expected: "Got:<nil>",
		},
	} {
		t.Run(name, func(t *testing.T) {
			f := NewFailure()
			if tt.expect != nil {
				f = f.Expect(tt.expect)
			}
			if tt.got != nil {
				f = f.Got(tt.got)
			}

			result := f.buildTypeBody()
			if result != tt.expected {
				t.Errorf("buildTypeBody() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestIsDifferentTypes(t *testing.T) {
	for name, tt := range map[string]struct {
		expect   *obj.Object
		got      *obj.Object
		expected bool
	}{
		"both nil": {
			expect:   nil,
			got:      nil,
			expected: false,
		},
		"expect nil, got not nil": {
			expect:   nil,
			got:      obj.NewObject("hello"),
			expected: true,
		},
		"expect not nil, got nil": {
			expect:   obj.NewObject("hello"),
			got:      nil,
			expected: true,
		},
		"same types": {
			expect:   obj.NewObject("hello"),
			got:      obj.NewObject("world"),
			expected: false,
		},
		"different types": {
			expect:   obj.NewObject("hello"),
			got:      obj.NewObject(123),
			expected: true,
		},
		"both nil values with touch true": {
			expect:   obj.NewObject(nil),
			got:      obj.NewObject(nil),
			expected: false,
		},
		"expect nil value with touch true, got nil": {
			expect:   obj.NewObject(nil),
			got:      nil,
			expected: true,
		},
		"expect nil, got nil value with touch true": {
			expect:   nil,
			got:      obj.NewObject(nil),
			expected: true,
		},
		"pointer types same": {
			expect:   obj.NewObject(&[]int{1, 2}),
			got:      obj.NewObject(&[]int{3, 4}),
			expected: false,
		},
		"pointer types different": {
			expect:   obj.NewObject(&[]int{1, 2}),
			got:      obj.NewObject(&[]string{"a", "b"}),
			expected: true,
		},
		"struct types same": {
			expect:   obj.NewObject(struct{ ID int }{ID: 1}),
			got:      obj.NewObject(struct{ ID int }{ID: 2}),
			expected: false,
		},
		"struct types different": {
			expect:   obj.NewObject(struct{ ID int }{ID: 1}),
			got:      obj.NewObject(struct{ Name string }{Name: "test"}),
			expected: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			f := NewFailure()
			if tt.expect != nil {
				f = f.Expect(tt.expect)
			}
			if tt.got != nil {
				f = f.Got(tt.got)
			}

			result := f.isDifferentTypes()
			if result != tt.expected {
				t.Errorf("isDifferentTypes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
