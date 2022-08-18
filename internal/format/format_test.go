package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	testData := []struct {
		input  []byte
		expect []byte
		error  bool
	}{
		{
			input: []byte(``),
			error: true,
		},
		{
			input: []byte(`package p`),
			expect: []byte(`package p
`),
			error: false,
		},
		{
			input: []byte(`package p;`),
			expect: []byte(`package p
`),
			error: false,
		},
		{
			input: []byte(`package p
		import "fmt"`),
			expect: []byte(`package p
`),
			error: false,
		},
		{
			input: []byte(`package p
import "foo'`),
			error: true,
		},
	}
	for _, data := range testData {
		actual, err := Source(data.input)
		if data.error {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, string(data.expect), string(actual))
	}
}
