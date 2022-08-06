package buffer

import (
	"fmt"
	"strings"
)

type b struct {
	list []string
}

func New() *b {
	return &b{}
}

func (b *b) Reset() {
	b.list = nil
}

func (b *b) Write(format string, a ...any) {
	b.list = append(b.list, fmt.Sprintf(format, a...))
}

func (b *b) String() string {
	return strings.Join(b.list, "\n")
}
