package buffer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	instance := New()
	assert.NotNil(t, instance)
}

func TestB_Reset(t *testing.T) {
	instance := New()
	instance.Write("foo")
	instance.Reset()
	assert.Equal(t, 0, len(instance.list))
}

func TestB_Write(t *testing.T) {
	testData := map[string]string{
		"":        "",
		"foo":     "foo",
		"foo,bar": "foo\nbar",
	}
	instance := New()
	for input, expected := range testData {
		instance.Reset()
		fields := strings.FieldsFunc(input, func(r rune) bool {
			return r == ','
		})
		for _, field := range fields {
			instance.Write(field)
		}
		assert.Equal(t, expected, instance.String())
	}
}

func TestB_String(t *testing.T) {
	instance := New()
	instance.Write("foo")
	assert.Equal(t, "foo", instance.String())
}
