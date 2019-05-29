package main

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/event"
	"github.com/liangran2018/100million/dial"
	"fmt"
	"github.com/liangran2018/100million/action"
)

func main() {
	base.NewLogFile()
	defer base.CloseLog()

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("panic: ", r)
			base.Log(base.Wrong, r)
		}
	}()

	event.StartEvent()

	for env.GetAge() <= 70 {
		dial.PersonShow()
		dial.HomePageShow()

		i := base.InputNum()
		switch i {
		case 1:
			dial.Market()
		case 2:
			dial.Company()
		case 6:
			action.Next()
		default:
			base.Notice("err")
			continue
		}
	}
}


