package stringx

import "strings"

func TrimNewLine(s string) string {
	var replacer = strings.NewReplacer("\r", "", "\n", "")
	return replacer.Replace(s)
}

func TrimSpace(s string) string {
	var r = strings.NewReplacer(" ", "", "\t", "")
	return r.Replace(s)
}

func RepeatJoin(s, sep string, count int) string {
	var list []string
	for i := 0; i < count; i++ {
		list = append(list, s)
	}

	return strings.Join(list, sep)
}
