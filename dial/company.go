package dial

import (
	"github.com/liangran2018/100million/base"
	"fmt"
	"github.com/liangran2018/100million/action"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/today"
)

var companyspage = []string{"", "buy", "sell", "return"}

func companyPageShow() {
	pageShow(companyspage)
}

func Company() {
	for {
		today.TodayCompanyShow()
		cs := today.GetTodayCompany()
		base.Notice(fmt.Sprintf("%d. %s", len(cs)+1, "返回"))

		i := base.InputNum()
		if i < 1 || i >= len(cs)+1 {
			return
		}

		singleCompany(cs[i])
	}
}

func singleCompany(c *today.NowCompany) {
	for {
		companyPageShow()
		i := base.InputNum()
		switch i {
		case 1:
			if c.IsSt {
				base.Notice("can't buy")
				continue
			}
			base.Notice("most buy " + base.StrVal(int(float32(env.MoneyGet()) / c.Price)))
			base.Notice("num")
			i := base.InputNum()
			base.Notice(action.Invest(c.Company, c.Price, float32(i)))
		case 2:
			if c.IsSt {
				base.Notice("can't sell")
				continue
			}
			base.Notice("most sell " + base.StrVal(own.CompanyNum(c.Company)))
			base.Notice("num")
			i := base.InputNum()
			base.Notice(action.DisInvest(c.Company, c.Price, float32(i)))
		default:
			return
		}
	}
}
