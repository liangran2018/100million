package main

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/event"
	"github.com/liangran2018/100million/dial"
	"github.com/liangran2018/100million/action"
)

func main() {
	base.NewLogFile()
	defer base.CloseLog()

	event.StartEvent()
	action.Next()
	dial.PersonShow()

	for env.GetAge() <= 70 {
		dial.HomePageShow()

		i := base.InputNum()
		switch i {
		case 1:
			dial.Market()
		case 2:
			dial.Company()
		case 3:
			
		case 6:
			action.Next()
		case 7:
			action.Retire()
		default:
			base.Notice("err")
			continue
		}

		dial.PersonShow()
	}
}


