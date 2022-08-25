package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoIncrement(t *testing.T) {
	test := map[string]string{
		"":     "",
		"1":    "2",
		"a":    "a1",
		"1a":   "1a1",
		"a1":   "a2",
		"a10":  "a11",
		"a1b1": "a1b2",
	}
	for input, expect := range test {
		actual := AutoIncrement(input, 1)
		assert.Equal(t, expect, actual)
	}
}

func TestTrimWhiteSpace(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "foo bar", expect: "foobar"},
		{input: "foo\nbar", expect: "foobar"},
		{input: "foo\rbar", expect: "foobar"},
		{input: "foo\tbar", expect: "foobar"},
		{input: "foo\n\r\tbar", expect: "foobar"},
	}
	for _, v := range testData {
		actual := TrimWhiteSpace(v.input)
		assert.Equal(t, v.expect, actual)
	}
}

func TestTrimNewLine(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "foo bar", expect: "foo bar"},
		{input: "foo\nbar", expect: "foobar"},
		{input: "foo\rbar", expect: "foobar"},
		{input: "foo\n\r\tbar", expect: "foo\tbar"},
	}
	for _, v := range testData {
		actual := TrimNewLine(v.input)
		assert.Equal(t, v.expect, actual)
	}
}

func TestTrimSpace(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "foo bar", expect: "foobar"},
		{input: "foo\nbar", expect: "foo\nbar"},
		{input: "foo\rbar", expect: "foo\rbar"},
		{input: "foo\n\r\tbar", expect: "foo\n\rbar"},
	}
	for _, v := range testData {
		actual := TrimSpace(v.input)
		assert.Equal(t, v.expect, actual)
	}
}

func TestRepeatJoin(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo,foo"},
	}
	for _, v := range testData {
		actual := RepeatJoin(v.input, ",", 2)
		assert.Equal(t, v.expect, actual)
	}
}

func TestFormatIdentifiers(t *testing.T) {
	var testData = []struct {
		input  string
		expect string
	}{
		{input: "", expect: ""},
		{input: "foo", expect: "foo"},
		{input: "foo bar", expect: "foo bar"},
		{input: "foo\nbar", expect: "foo bar"},
		{input: "foo\tbar", expect: "foo bar"},
		{input: "foo\rbar", expect: "foo bar"},
		{input: "foo\fbar", expect: "foo bar"},
		{input: "foo\n\t\r\fbar", expect: "foo bar"},
	}
	for _, v := range testData {
		actual := FormatIdentifiers(v.input)
		assert.Equal(t, v.expect, actual)
	}
}
