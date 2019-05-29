package dial

import (
	"fmt"
	"strings"

	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/news"
	"github.com/liangran2018/100million/action"
)

var marketpage = []string{"", "goods", "quotations", "storehouse", "return"}
var goodspage = []string{"", "intro", "buy", "return"}
var storepage = []string{"", "intro", "sell", "return"}

func HomePageShow() {
	pageShow(homepage)
}

func MarketPageShow() {
	pageShow(marketpage)
}

func goodPageShow() {
	pageShow(goodspage)
}

func StorePageShow() {
	pageShow(storepage)
}

func Market() {
	for {
		MarketPageShow()
		i := base.InputNum()
		switch i {
		case 1:
			Goods()
		case 2:
			base.Notice(strings.Join(news.MarketNewsShow(), "\n"))
		case 3:
			Stores()
		default:
			return
		}
	}
}

func Goods() {
	for {
		gs := news.GetMarket()
		for k, v := range gs {
			base.Notice(fmt.Sprintf("%d. %s: %d", k+1, v.Goods.Name(), v.Price))
		}
		base.Notice(fmt.Sprintf("%d. %s", len(gs)+1, "return"))

		i := base.InputNum()
		if i < 1 || i >= len(gs)+1 {
			return
		}

		singleGoods(gs[i-1])
	}
}

func Stores() {
	for {
		gs := own.GetStore()
		if len(gs) == 0 {
			base.Notice("nothing")
			return
		}

		for k, v := range gs {
			base.Notice(fmt.Sprintf("%d. %s: %d %d", k, v.Goods.Name(), v.Num, v.Price))
		}
		base.Notice(fmt.Sprintf("%d. %s", len(gs)+1, "return"))

		i := base.InputNum()
		if i < 1 || i >= len(gs)+1 {
			return
		}

		singleStore(gs[i])
	}
}

func singleGoods(m *news.NowMarket) {
	for {
		goodPageShow()
		i := base.InputNum()
		switch i {
		case 1:
			base.Notice(m.Goods.Intro())
		case 2:
			base.Notice("most buy " + base.StrVal(own.MostBuy(m.Price)))
			base.Notice("num")
			i := base.InputNum()
			base.Notice(action.Buy(m.Goods, m.Price, i))
		default:
			return
		}
	}
}

func singleStore(s *own.NowStore) {
	for {
		StorePageShow()
		i := base.InputNum()
		switch i {
		case 1:
			base.Notice(s.Goods.Intro())
		case 2:
			base.Notice("num")
			i := base.InputNum()
			base.Notice(action.Sell(s.Goods, s.Price, i))
		default:
			return
		}
	}
}
