package base

import (
	"bufio"
	"os"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

func Input() string {
	i, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(i)
	if s == "" {
		panic(err)
	}

	return s
}

func InputNum() int {
	s := Input()
	i, err := IntVal(s)
	if err != nil {
		panic(err)
	}

	return i
}

//判断参数是否为空
func Empty(para string) bool {
	para = strings.TrimSpace(para)
	if para == "" {
		return true
	}

	return false
}
