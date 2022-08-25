package stringx

import (
	"fmt"
	"strconv"
	"strings"
)

func TrimWhiteSpace(s string) string {
	ret := TrimNewLine(s)
	return TrimSpace(ret)
}

func TrimNewLine(s string) string {
	var replacer = strings.NewReplacer("\r", "", "\n", "")
	return replacer.Replace(s)
}

func TrimSpace(s string) string {
	var r = strings.NewReplacer(" ", "", "\t", "")
	return r.Replace(s)
}

func RepeatJoin(s, sep string, count int) string {
	if len(s) == 0 {
		return ""
	}

	var list []string
	for i := 0; i < count; i++ {
		list = append(list, s)
	}

	return strings.Join(list, sep)
}

func AutoIncrement(s string, step int) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	for i := 0; i < length; i++ {
		r := s[i]
		if r >= '0' && r <= '9' {
			if num, ok := IsNumber(s[i:]); ok {
				return fmt.Sprintf("%s%d", s[:i], num+uint64(step))
			}
		}
	}
	return fmt.Sprintf("%s%d", s, step)
}

func IsNumber(s string) (uint64, bool) {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

func FormatIdentifiers(s string) string {
	var list = strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\r' || r == '\n' || r == '\f'
	})
	var target []string
	for _, v := range list {
		if len(v) > 0 {
			target = append(target, v)
		}
	}

	return strings.Join(target, " ")
}
