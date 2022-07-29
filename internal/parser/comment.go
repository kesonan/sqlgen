package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/stringx"
)

var (
	singleLineComment = []byte("--")
	fnPrefix          = "fn:"
	fnRegex           = `(?m)^[a-zA-Z]\w*`
)

func parseLineComment(sql string) (spec.Comment, error) {
	r := bufio.NewReader(strings.NewReader(sql))
	var comment spec.Comment
	comment.OriginText = sql
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		comment.LineText = append(comment.LineText, string(line))
		if bytes.HasPrefix(line, singleLineComment) {
			var text = strings.TrimSpace(string(line[2:]))
			funcName, err := parseFuncName(text)
			if err != nil {
				return spec.Comment{}, err
			}

			if len(funcName) > 0 { // it will override the previous comment.
				comment.FuncName = funcName
			}
		}
	}

	return comment, nil
}

func parseFuncName(text string) (string, error) {
	s := stringx.TrimSpace(text)
	if !strings.HasPrefix(s, fnPrefix) {
		return "", nil
	}

	fn := strings.TrimPrefix(s, fnPrefix)
	if len(fn) == 0 {
		return "", fmt.Errorf("missing function name")
	}

	match, _ := regexp.MatchString(fnRegex, fn)
	if match {
		return fn, nil
	}
	return "", fmt.Errorf("invalid function name: %s", fn)
}
