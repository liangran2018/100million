package dial

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/base"
	"fmt"
	"github.com/liangran2018/100million/news"
)

var homepage = []string{"", "market"}
var marketpage = []string{"", "goods", "quotations", "storehouse"}
var goodspage = []string{"", "intro", "buy"}
var storepage = []string{"", "intro", "sell"}

func PersonShow() {
	total := env.MoneyGet() + own.GoodsTotalPrice()
	base.Notice(fmt.Sprintf("%d %d %d %d %d", env.GetAge(), env.HealthGet(), env.ReputeGet(), env.MoneyGet(), total))
}

func pageShow(page []string) int {
	for k, v := range page {
		if k == 0 { continue }
		base.Notice(fmt.Sprintf("%d. %s", k, v))
	}

	return len(page)
}

func HomePageShow() int {
	return pageShow(homepage)
}

func MarketPageShow() int {
	return pageShow(marketpage)
}

func GoodsPageShow() int {
	gs := news.GoodsShow()
	i := 1
	for k, v := range gs {
		base.Notice(fmt.Sprintf("%d. %s: %d", i, k, v))
		i++
	}

	return i
}

func

func HomePageChoose(i int) int {
	if i == 1 {
		return MarketPageShow()
	}
	return 0
}

func MarketPageChoose(i int) int {
	if i == 1 {
		return GoodsPageShow()
	} else if i == 2 {
		return news.NewsShow()
	}
}