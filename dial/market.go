package dial

import (
	"fmt"
	"github.com/liangran2018/100million/own"
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/news"
	"github.com/liangran2018/100million/action"
)

var marketpage = []string{"", "物品", "市场行情", "仓库", "返回"}
var goodspage = []string{"", "简介", "买入", "返回"}
var storepage = []string{"", "简介", "卖出", "返回"}

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
			news.MarketNewsShow()
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
		base.Notice(fmt.Sprintf("%d. %s", len(gs)+1, "返回"))

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
			base.Notice("空空如也")
			return
		}

		for k, v := range gs {
			base.Notice(fmt.Sprintf("%d. %s  数量: %d 买入价: %d", k+1, v.Goods.Name(), v.Num, v.Price))
		}
		base.Notice(fmt.Sprintf("%d. %s", len(gs)+1, "返回"))

		i := base.InputNum()
		if i < 1 || i >= len(gs)+1 {
			return
		}

		singleStore(gs[i-1])
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
			base.Notice(fmt.Sprintf("最多可买%d个", own.MostBuy(m.Price)))
			base.Notice("购买数量")
			i := base.InputNum()
			if i <= 0 {
				base.Notice("数量有误")
				return
			}
			base.Notice(action.Buy(m.Goods, m.Price, i))
			return
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
			base.Notice(fmt.Sprintf("最多可卖%d个", s.Num))
			base.Notice("卖出数量")
			i := base.InputNum()
			if i <= 0 {
				base.Notice("数量有误")
				return
			}
			base.Notice(action.Sell(s.Goods, s.Price, i))
			return
		default:
			return
		}
	}
}
