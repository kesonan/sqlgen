package templatex

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/anqiansong/sqlgen/internal/log"
)

const name = "_"

// T is a template helper.
type T struct {
	t      *template.Template
	buffer *bytes.Buffer
}

// New creates a new template helper.
func New() *T {
	var t = template.New(name)
	return &T{
		t:      t,
		buffer: bytes.NewBuffer(nil),
	}
}

// MustParse parses the template.
func (t *T) MustParse(text string) *T {
	t.t.Funcs(funcMap)
	_, err := t.t.Parse(text)
	log.Must(err)
	return t
}

// MustExecute executes the template.
func (t *T) MustExecute(data interface{}) *T {
	t.buffer.Reset()
	log.Must(t.t.Execute(t.buffer, data))
	return t
}

// MustSaveAs saves the template to the given filename, it will overwrite the file if it exists.
func (t *T) MustSaveAs(filename string) {
	log.Must(ioutil.WriteFile(filename, t.buffer.Bytes(), 0666))
}

// MustSave saves the template to the given filename, it will do nothing if it exists.
func (t *T) MustSave(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		t.MustSaveAs(filename)
	}
}
