package templatex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperCamel(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "Foo"},
		{input: "foo bar", expect: "FooBar"},
		{input: "foo_bar", expect: "FooBar"},
		{input: "foo-bar", expect: "FooBar"},
		{input: "_foobar", expect: "Foobar"},
		{input: "_foobar_", expect: "Foobar"},
	}
	for _, v := range testData {
		actual := UpperCamel(v.input)
		assert.Equal(t, v.expect, actual)
	}
}

func TestLowerCamel(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "Foo bar", expect: "fooBar"},
		{input: "Foo_bar", expect: "fooBar"},
		{input: "Foo-bar", expect: "fooBar"},
		{input: "_foobar", expect: "Foobar"},
		{input: "_foobar_", expect: "Foobar"},
		{input: "FooBar", expect: "fooBar"},
		{input: "Foo_Bar", expect: "fooBar"},
	}
	for _, v := range testData {
		actual := LowerCamel(v.input)
		assert.Equal(t, v.expect, actual)
	}
}

func TestJoin(t *testing.T) {
	var testData = []struct {
		input  []string
		expect string
	}{
		{input: []string{}, expect: ""},
		{input: []string{"foo"}, expect: "foo"},
		{input: []string{"foo", "bar"}, expect: "foo,bar"},
	}
	for _, v := range testData {
		actual := Join(v.input, ",")
		assert.Equal(t, v.expect, actual)
	}
}

func TestLineComment(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "foo\nbar", expect: "foo\n// bar"},
		{input: "foo\nbar\n", expect: "foo\n// bar"},
		{input: "\nfoo\nbar\n", expect: "foo\n// bar"},
	}
	for _, v := range testData {
		actual := LineComment(v.input)
		assert.Equal(t, v.expect, actual)
	}
}
