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

func GoodPageShow() {
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

		singleGoods(gs[i])
	}
}

func Stores() {
	for {
		gs := own.GetStore()
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
		GoodPageShow()
		i := base.InputNum()
		switch i {
		case 1:
			base.Notice(m.Goods.Intro())
		case 2:
			base.Notice("num")
			i := base.InputNum()
			ok := action.Buy(m.Goods, m.Price, i)
			if !ok {
				base.Notice("fail")
			} else {
				base.Notice("ok")
			}
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
			ok := action.Sell(s.Goods, s.Price, i)
			if !ok {
				base.Notice("fail")
			} else {
				base.Notice("ok")
			}
		default:
			return
		}
	}
}
