package dial

import (
	"github.com/liangran2018/100million/news"
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/action"
	"fmt"
)

func Companys() {
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
