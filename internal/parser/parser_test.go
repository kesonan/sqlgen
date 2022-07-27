package parser

import (
	_ "embed"
	"fmt"
	"log"
	"testing"
)

//go:embed test.sql
var testSql string

func TestParse(t *testing.T) {
	dxl, err := Parse(testSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dxl)
}
