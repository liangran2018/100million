package dial

import (
	"github.com/liangran2018/100million/news"
	"github.com/liangran2018/100million/base"
	"fmt"
	"github.com/liangran2018/100million/action"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
)

var companyspage = []string{"", "intro", "buy", "sell", "return"}

func companyPageShow() {
	pageShow(companyspage)
}

func Company() {
	for {
		cs := news.GetCompany()
		for k, v := range cs {
			if v.IsSt {
				base.Notice(fmt.Sprintf("%d. %s: 退市", k+1, v.Company.Name()))
				continue
			}
			base.Notice(fmt.Sprintf("%d. %s 价格: %.2f", k+1, v.Company.Name(), v.Price))
		}
		base.Notice(fmt.Sprintf("%d. %s", len(cs)+1, "返回"))

		i := base.InputNum()
		if i < 1 || i >= len(cs)+1 {
			return
		}

		singleCompany(cs[i])
	}
}

func singleCompany(c *news.NowCompany) {
	for {
		companyPageShow()
		i := base.InputNum()
		switch i {
		case 1:
			base.Notice(c.Company.Intro())
		case 2:
			if c.IsSt {
				base.Notice("can't buy")
				continue
			}
			base.Notice("most buy " + base.StrVal(int(float32(env.MoneyGet()) / c.Price)))
			base.Notice("num")
			i := base.InputNum()
			base.Notice(action.Invest(c.Company, c.Price, float32(i)))
		case 3:
			if c.IsSt {
				base.Notice("can't buy")
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
