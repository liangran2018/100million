package main

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/event"
	"github.com/liangran2018/100million/dial"
)

func main() {
	base.NewLogFile()
	defer base.CloseLog()

	event.StartEvent()
	dial.PersonShow()

	for env.GetAge() <= 70 {
		lhome := dial.HomePageShow()
		i := base.InputNum()
		if i == 0 || i >= lhome {
			base.Notice("err")
			continue
		}

		lNext := dial.HomePageChoose(i)
		i = base.InputNum()
		if i == 0 || i >= lNext {
			base.Notice("err")
			continue
		}

	}


}


