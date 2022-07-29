package parser

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/anqiansong/sqlgen/internal/spec"
)

//-- fn: Insert
//-- name: foo
//-- 用户数据插入
//insert into user (user, name, password, mobile)
//values ('test', 'test', 'test', 'test');

var (
	singleLineComment = []byte("--")
)

func parseLineComment(sql string) spec.Comment {
	r := bufio.NewReader(strings.NewReader(sql))
	var comment spec.Comment
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasPrefix(line, singleLineComment) {
			var text = string(line[2:])
			var funcName = parseFuncName(text)
			if len(funcName) > 0 { // it will override the previous comment.
				comment.FuncName = funcName
			}
			comment.LineText = append(comment.LineText, text)
		}
	}

	return comment
}

func parseFuncName(text string) string {
	return ""
}
