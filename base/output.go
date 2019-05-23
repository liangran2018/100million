package base

import (
	"fmt"
)

func Notice(msg string) {
	fmt.Println(msg)
}

func Output(msg ...interface{}) {
	fmt.Println(msg...)
	Log(Info, msg...)
}
