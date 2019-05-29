package dial

import (
	"fmt"

	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/base"
)

var homepage = []string{"", "市场", "公司", "消费", "约会", "一亿小目标", "下一年", "提前退休"}

func PersonShow() {
	total := env.MoneyGet() + own.GoodsTotalPrice() + own.CompanyTotalPrice()
	base.Notice(fmt.Sprintf("年龄:%d 健康:%d 名声:%d 现金:%d 总资产:%d", env.GetAge(), env.HealthGet(), env.ReputeGet(), env.MoneyGet(), total))
}

func pageShow(page []string) {
	for k, v := range page {
		if k == 0 { continue }
		base.Notice(fmt.Sprintf("%d. %s", k, v))
	}
}