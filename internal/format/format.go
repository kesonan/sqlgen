package format

import (
	"go/format"

	"golang.org/x/tools/imports"
)

// Source formats go code and imports.
func Source(data []byte) ([]byte, error) {
	ret, err := format.Source(data)
	if err != nil {
		return nil, err
	}

	return imports.Process("", ret, nil)
}
