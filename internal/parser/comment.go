package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/scanner"

	"github.com/anqiansong/sqlgen/internal/spec"
	"github.com/anqiansong/sqlgen/internal/stringx"
)

const (
	plainTextMode = iota
	lineCommentMode
	docCommentOpenMode
	docCommentCloseMode
)

var (
	singleLineComment = []byte("--")
	fnPrefix          = "fn:"
	fnRegex           = `(?m)^[a-zA-Z]\w*`
)

type segment struct {
	start, end int
}

type sqlScanner struct {
	*scanner.Scanner
	mode       int
	source     string
	startIndex int
	segments   []segment
	trim       string
}

func NewSqlScanner(s string) *sqlScanner {
	var sc scanner.Scanner
	return &sqlScanner{
		Scanner: sc.Init(strings.NewReader(s)),
		mode:    plainTextMode,
		source:  s,
		trim:    s,
	}
}

// ScanAndTrim ignores non comment text.
func (s *sqlScanner) ScanAndTrim() (string, error) {
	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() {
		switch tok {
		case '-':
			s.enterLineCommentMode()
		case '/':
			err := s.enterDocCommentMode()
			if err != nil {
				return "", err
			}
		}
	}

	if len(s.segments) == 0 {
		return s.source, nil
	}

	var list []string
	for i := 0; i < len(s.segments); i++ {
		var segment string
		if i == 0 {
			if s.segments[i].start > 0 {
				segment = s.source[:s.segments[i].start]
				list = append(list, stringx.TrimNewLine(segment))
			}
		}
		if i != len(s.segments)-1 {
			segment = s.source[s.segments[i].end:s.segments[i+1].start]
		} else {
			segment = s.source[s.segments[i].end:]
		}

		s := stringx.TrimWhiteSpace(segment)
		if len(s) == 0 {
			continue
		}

		list = append(list, segment)
	}

	return stringx.FormatIdentifiers(strings.Join(list, "")), nil
}

func (s *sqlScanner) enterDocCommentMode() error {
	var position = s.CurrentPosition()
	s.startIndex = position - 1
	if s.startIndex < 0 {
		s.startIndex = 0
	}

	tok := s.Next()
	if tok != '*' {
		s.startIndex = 0
		return nil
	}

	s.mode = docCommentOpenMode
	for {
		tok := s.Next()
		if tok == scanner.EOF {
			s.mode = plainTextMode
			return fmt.Errorf("expected close flag '*/'")
		}

		if tok == '*' {
			s.mode = docCommentCloseMode
		} else if tok == '/' {
			if s.mode == docCommentCloseMode {
				s.segments = append(s.segments, segment{
					start: s.startIndex,
					end:   s.CurrentPosition(),
				})
			}
			break
		} else {
			if s.mode == docCommentCloseMode {
				s.mode = docCommentOpenMode
			}
		}
	}
	s.mode = plainTextMode
	return nil
}

func (s *sqlScanner) CurrentPosition() int {
	offset := s.Pos().Offset
	return offset
}

func (s *sqlScanner) enterLineCommentMode() {
	var position = s.CurrentPosition()
	s.startIndex = position - 1
	if s.startIndex < 0 {
		s.startIndex = 0
	}
	tok := s.Next()
	if tok != '-' {
		s.startIndex = 0
		return
	}

	s.mode = lineCommentMode
	for {
		tok = s.Next()
		if tok == scanner.EOF || tok == '\n' {
			s.segments = append(s.segments, segment{
				start: s.startIndex,
				end:   s.CurrentPosition(),
			})
			break
		}
	}
	s.mode = plainTextMode
}

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
		return "", errorMissingFunction
	}

	match, _ := regexp.MatchString(fnRegex, fn)
	if match {
		return fn, nil
	}

	return "", fmt.Errorf("invalid function name: %s", fn)
}
