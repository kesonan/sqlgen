package parameter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	instance := New()
	assert.NotNil(t, instance)
}

func TestP_Add(t *testing.T) {
	instance := New()
	instance.Add(Parameter{
		Column:   "foo",
		Type:     "bar",
		ThirdPkg: "baz",
	}, Parameter{
		Column:   "foo1",
		Type:     "bar1",
		ThirdPkg: "baz1",
	}, Parameter{
		Column:   "foo",
		Type:     "bar",
		ThirdPkg: "baz",
	})
}

func TestP_List(t *testing.T) {
	instance := New()
	expected := Parameters{
		Parameter{
			Column:   "foo",
			Type:     "bar",
			ThirdPkg: "baz",
		},
		Parameter{
			Column:   "foo1",
			Type:     "bar",
			ThirdPkg: "baz",
		},
		Parameter{
			Column:   "foo1",
			Type:     "bar1",
			ThirdPkg: "baz1",
		},
	}
	instance.Add(
		Parameter{
			Column:   "foo",
			Type:     "bar",
			ThirdPkg: "baz",
		},
		Parameter{
			Column:   "foo",
			Type:     "bar",
			ThirdPkg: "baz",
		},
		Parameter{
			Column:   "foo1",
			Type:     "bar1",
			ThirdPkg: "baz1",
		})
	actual := instance.List()
	assert.Equal(t, len(expected), len(actual))
	for idx, expectedItem := range expected {
		actualItem := actual[idx]
		assert.Equal(t, expectedItem, actualItem)
	}
}
