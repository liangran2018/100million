package action

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/news"
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/own"
)

func Next() {
	env.NextYear()

	news.NewMarketNew()
	marketNews := news.MarketNewsShow()
	for k := range marketNews {
		base.Notice(marketNews[k])
	}
	base.Log(base.Info, marketNews)

	news.NewCompany()
	nowCompany := news.GetCompany()
	ownCompany := own.GetOwnCompany()
	for k := range nowCompany {
		if nowCompany[k].IsSt {
			if _, ok := ownCompany[nowCompany[k].Company]; ok {
				base.Output(nowCompany[k].Company.Name() + "is st")
				env.MoneyAdd(own.St(nowCompany[k].Company))
			}
		}
	}
}
