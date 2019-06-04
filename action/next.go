package action

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/today"
)

func Next() {
	env.NextYear()

	today.TodayNews()
	today.TodayNewsShow()

	today.TodayCompany()
	//today.TodayCompanyShow()
}
