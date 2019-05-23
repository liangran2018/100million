package base

import (
	"fmt"
)

func Output(msg ...interface{}) {
	fmt.Println(msg...)
	Log(Info, msg...)
}
