package base

import (
	"bufio"
	"os"
	"strings"
)

var input = bufio.NewReader(os.Stdin)

func Input() (string, error) {
	i, err := input.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(i), nil
}

//判断参数是否为空
func Empty(para string) bool {
	para = strings.TrimSpace(para)
	if para == "" {
		return true
	}

	return false
}
