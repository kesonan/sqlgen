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
