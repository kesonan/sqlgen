package templatex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	instance := New()
	assert.NotNil(t, instance)
}

func TestT_AppendFuncMap(t *testing.T) {
	instance := New()
	fn := func() string {
		return "any"
	}
	list := template.FuncMap{
		"foo": fn,
		"bar": fn,
		"baz": fn,
	}
	instance.AppendFuncMap(list)
	for k := range list {
		_, ok := instance.fm[k]
		assert.True(t, ok)
	}
}

func TestT_MustParse(t *testing.T) {
	instance := New()
	instance.AppendFuncMap(template.FuncMap{"foo": func() string { return "bar" }})
	ret := instance.MustParse("{{foo}}")
	assert.Equal(t, instance, ret)
}

func TestT_MustExecute(t *testing.T) {
	instance := New()
	instance.MustParse("{{.foo}}")
	ret := instance.MustExecute(map[string]string{
		"foo": "bar",
	})
	assert.Equal(t, instance, ret)
	assert.Equal(t, "bar", instance.buffer.String())
}

func TestT_MustSaveAs(t *testing.T) {
	t.Run("format_false", func(t *testing.T) {
		instance := New()
		instance.MustParse("{{.foo}}")
		instance.MustExecute(map[string]string{
			"foo": "bar",
		})
		tempFile := filepath.Join(t.TempDir(), "foo")
		instance.MustSaveAs(tempFile, false)
		data, err := ioutil.ReadFile(tempFile)
		assert.NoError(t, err)
		assert.Equal(t, "bar", string(data))
	})

	t.Run("format_true", func(t *testing.T) {
		instance := New()
		instance.MustParse(" package {{.foo}}")
		instance.MustExecute(map[string]string{
			"foo": "bar",
		})
		tempFile := filepath.Join(t.TempDir(), "foo")
		instance.MustSaveAs(tempFile, true)
		data, err := ioutil.ReadFile(tempFile)
		assert.NoError(t, err)
		assert.Equal(t, "package bar\n", string(data))
	})
}

func TestT_MustSave(t *testing.T) {
	instance := New()
	instance.MustParse("{{.foo}}")
	instance.MustExecute(map[string]string{
		"foo": "bar",
	})
	tempFile := filepath.Join(t.TempDir(), "foo")
	instance.MustSave(tempFile, false)
	data, err := ioutil.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, "bar", string(data))

	instance.MustExecute(map[string]string{
		"foo": "baz",
	})
	instance.MustSave(tempFile, false)
	data, err = ioutil.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, "bar", string(data))
}

func TestT_Write(t *testing.T) {
	t.Run("format_false", func(t *testing.T) {
		instance := New()
		instance.MustParse("{{.foo}}")
		instance.MustExecute(map[string]string{
			"foo": "bar",
		})
		tempFile := filepath.Join(t.TempDir(), "foo")
		file, err := os.OpenFile(tempFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		assert.NoError(t, err)
		t.Cleanup(func() {
			file.Close()
		})

		instance.Write(file, false)
		data, err := ioutil.ReadFile(tempFile)
		assert.NoError(t, err)
		assert.Equal(t, "bar", string(data))
	})

	t.Run("format_true", func(t *testing.T) {
		instance := New()
		instance.MustParse(" package {{.foo}}")
		instance.MustExecute(map[string]string{
			"foo": "bar",
		})
		tempFile := filepath.Join(t.TempDir(), "foo")
		file, err := os.OpenFile(tempFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		assert.NoError(t, err)
		t.Cleanup(func() {
			file.Close()
		})

		instance.Write(file, true)
		data, err := ioutil.ReadFile(tempFile)
		assert.NoError(t, err)
		assert.Equal(t, "package bar\n", string(data))
	})
}
