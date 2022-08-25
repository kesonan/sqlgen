package patterns

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPattern_Match(t *testing.T) {
	list := Pattern{}
	matched := list.Match()
	assert.Equal(t, []string(nil), matched)
	matched = list.Match("foo")
	assert.Equal(t, []string(nil), matched)

	list = Pattern{"*"}
	matched = list.Match()
	assert.Equal(t, []string(nil), matched)
	matched = list.Match("foo")
	assert.Equal(t, "foo", matched[0])
}

func TestNew(t *testing.T) {
	p := New()
	assert.Equal(t, "*", p[0])

	p = New("foo")
	assert.Equal(t, "foo", p[0])

	p = New("foo,bar,baz,baz")
	assert.Equal(t, "foo,bar,baz", strings.Join(p, ","))
}
