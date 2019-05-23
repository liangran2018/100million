package event

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/env"
)

var startMsg = []string{"5000", "10000", "100000"}

func StartEvent() {
	i := base.Rand(5)

	var smsg string
	if i < 2 {
		smsg = startMsg[0]
		env.MoneyAdd(5000)
	} else if i < 4 {
		smsg = startMsg[1]
		env.MoneyAdd(10000)
	} else {
		smsg = startMsg[2]
		env.MoneyAdd(100000)
	}

	base.Output(smsg)
}