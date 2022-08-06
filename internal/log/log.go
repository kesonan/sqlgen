package log

import (
	"fmt"
	"os"
)

func Must(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
