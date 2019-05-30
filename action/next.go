package action

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/news"
)

func Next() {
	env.NextYear()

	news.NewMarketNew()
	news.MarketNewsShow()

	//news.NewCompany()
	//nowCompany := news.GetCompany()
	//ownCompany := own.GetOwnCompany()
	//for k := range nowCompany {
	//	if nowCompany[k].IsSt {
	//		if _, ok := ownCompany[nowCompany[k].Company]; ok {
	//			base.Output(nowCompany[k].Company.Name() + "is st")
	//			env.MoneyAdd(own.St(nowCompany[k].Company))
	//		}
	//	}
	//}
}
