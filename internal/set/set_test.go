package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSet(t *testing.T) {
	listSet := From("a", "b", "c")
	var result = listSet.String()
	assert.Equal(t, []string{"a", "b", "c"}, result)
	listSet.Add("b", "c", "d")
	result = listSet.String()
	assert.Equal(t, []string{"a", "b", "c", "d"}, result)
	listSet.Remove("a")
	result = listSet.String()
	assert.Equal(t, []string{"b", "c", "d"}, result)
	exists := listSet.Exists("b")
	assert.True(t, exists)
	listSet.Add(1, 2, 3)
	intResult := listSet.Int()
	assert.Equal(t, []int{1, 2, 3}, intResult)
	listSet.Add(int32(4), int32(5), int32(6))
	int32Result := listSet.Int32()
	assert.Equal(t, []int32{4, 5, 6}, int32Result)
	listSet.Add(int64(7), int64(8), int64(9))
	int64Result := listSet.Int64()
	assert.Equal(t, []int64{7, 8, 9}, int64Result)
}

func TestFromString(t *testing.T) {
	s := FromString()
	assert.NotNil(t, s)
	s = FromString("foo", "bar", "baz")
	var result = s.String()
	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestListSet_AddStringList(t *testing.T) {
	s := FromString()
	s.AddStringList([]string{"foo", "bar", "baz"})
	var result = s.String()
	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestListSet_Remove(t *testing.T) {
	s := From("foo", "bar")
	s.Remove("baz")
	s.Remove("bar")
	var result = s.String()
	assert.Equal(t, []string{"foo"}, result)
}

func TestListSet_Init(t *testing.T) {
	s := From("foo", "bar", "baz")
	s.Init()
	var result = s.String()
	assert.Equal(t, []string(nil), result)
}

func TestListSet_Range(t *testing.T) {
	var str []string
	s := From()
	s.Range(func(v interface{}) {
		str = append(str, fmt.Sprint(v))
	})
	s = From("foo")
	s.Range(func(v interface{}) {
		str = append(str, fmt.Sprint(v))
	})

	assert.Equal(t, []string{"foo"}, str)
}
