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
		dial.HomePageShow()

		i := base.InputNum()
		switch i {
		case 1:
			dial.Market()
		case 2:
			
		default:
			base.Notice("err")
			continue
		}
	}


}


