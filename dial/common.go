package dial

import (
	"fmt"

	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/base"
)

var homepage = []string{"", "market", "company", "consume", "lover", "100million", "next year", "retire"}

func PersonShow() {
	total := env.MoneyGet() + own.GoodsTotalPrice() + own.CompanyTotalPrice()
	base.Notice(fmt.Sprintf("age:%d health:%d repute:%d cash:%d total:%d", env.GetAge(), env.HealthGet(), env.ReputeGet(), env.MoneyGet(), total))
}

func pageShow(page []string) {
	for k, v := range page {
		if k == 0 { continue }
		base.Notice(fmt.Sprintf("%d. %s", k, v))
	}
}