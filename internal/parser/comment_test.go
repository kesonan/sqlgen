package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anqiansong/sqlgen/internal/spec"
)

func Test_parseLineComment(t *testing.T) {
	var test = []struct {
		input    string
		expected spec.Comment
	}{
		{input: "", expected: spec.Comment{}},
		{input: "--fn:a", expected: spec.Comment{
			OriginText: "--fn:a",
			LineText:   []string{"--fn:a"},
			FuncName:   "a",
		}},
		{input: "-- fn: ab", expected: spec.Comment{
			OriginText: "-- fn: ab",
			LineText:   []string{"-- fn: ab"},
			FuncName:   "ab",
		}},
		{input: "--  fn : a1", expected: spec.Comment{
			OriginText: "--  fn : a1",
			LineText:   []string{"--  fn : a1"},
			FuncName:   "a1",
		}},
		{input: "-- fn:A", expected: spec.Comment{
			OriginText: "-- fn:A",
			LineText:   []string{"-- fn:A"},
			FuncName:   "A",
		}},
		{input: "--		fn:A1", expected: spec.Comment{
			OriginText: "--\t\tfn:A1",
			LineText:   []string{"--\t\tfn:A1"},
			FuncName:   "A1",
		}},
		{input: "-- fn : A_ ", expected: spec.Comment{
			OriginText: "-- fn : A_ ",
			LineText:   []string{"-- fn : A_ "},
			FuncName:   "A_",
		}},
		{input: "-- fn : A_\n-- plain text", expected: spec.Comment{
			OriginText: "-- fn : A_\n-- plain text",
			LineText:   []string{"-- fn : A_", "-- plain text"},
			FuncName:   "A_",
		}},
		{input: `-- fn: Insert
-- name: foo
-- 用户数据插入
insert into user (user, name, password, mobile)
values ('test', 'test', 'test', 'test');`, expected: spec.Comment{
			OriginText: `-- fn: Insert
-- name: foo
-- 用户数据插入
insert into user (user, name, password, mobile)
values ('test', 'test', 'test', 'test');`,
			LineText: []string{"-- fn: Insert", "-- name: foo", "-- 用户数据插入", "insert into user (user, name, password, mobile)", "values ('test', 'test', 'test', 'test');"},
			FuncName: "Insert",
		}},
	}
	for _, c := range test {
		actual, _ := parseLineComment(c.input, true)
		assert.Equal(t, c.expected, actual)
	}
}

func Test_parseFuncName(t *testing.T) {
	var test = []struct {
		input    string
		expected string
		err      bool
	}{
		{input: "", expected: ""},
		{input: "fn:", err: true},
		{input: "fn:a", expected: "a"},
		{input: "fn: ab", expected: "ab"},
		{input: "fn : a1", expected: "a1"},
		{input: "fn:A", expected: "A"},
		{input: "fn:A1", expected: "A1"},
		{input: " fn : A_ ", expected: "A_"},
		{input: "fn:2", err: true},
		{input: "fn:_", err: true},
		{input: "fn:-", err: true},
	}
	for _, c := range test {
		actual, err := parseFuncName(c.input)
		if c.err {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, c.expected, actual)
	}
}

func Test_trimComment(t *testing.T) {
	test := []struct {
		input  string
		expect string
		err    bool
	}{
		{},
		{
			input:  " ",
			expect: " ",
		},
		{
			input:  "-",
			expect: "-",
		},
		{
			input:  "--",
			expect: "",
		},
		{
			input: `--foo--bar
foo`,
			expect: "foo",
		},
		{
			input: `--foo--bar
foo
bar`,
			expect: "foo bar",
		},
		{
			input:  `/**/foo`,
			expect: "foo",
		},
		{
			input:  `foo/**/ bar`,
			expect: "foo bar",
		},
		{
			input: `foo/**/ 
--foo
/*--*/
bar`,
			expect: "foo bar",
		},
		{
			input: "/*",
			err:   true,
		},
		{
			input: "/**",
			err:   true,
		},
	}
	for _, v := range test {
		s := NewSqlScanner(v.input)
		actual, err := s.ScanAndTrim()
		if v.err {
			assert.Error(t, err)
			continue
		}
		assert.Equal(t, v.expect, actual)
	}
}
