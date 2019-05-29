package event

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/env"
)

var startMsg = []string{"20岁时白手起家，揣着5000元在社会闯荡",
	"20岁时作为一名中产阶级家庭的孩子，攒了10000元开始在社会打拼",
	"20岁时家族财力雄厚，家里给了100000元学做生意"}

var startI int

func StartEvent() {
	i := base.Rand(5)

	var smsg string
	if i < 2 {
		startI = 0
		smsg = startMsg[startI]
		env.MoneyAdd(5000)
	} else if i < 4 {
		startI = 1
		smsg = startMsg[startI]
		env.MoneyAdd(10000)
	} else {
		startI = 2
		smsg = startMsg[startI]
		env.MoneyAdd(100000)
	}

	base.Output(smsg)
}
