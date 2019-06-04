package today

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/goods"
	"github.com/liangran2018/100million/news"
)

type NowMarket struct {
	Goods goods.GoodsIndex
	News news.NewsIndex
	Price int
}

var save []*NowMarket

const goodsNum = 5

func TodayNews() {
	save = make([]*NowMarket, goodsNum)

	g := base.RandMany(int(goods.GoodsEnd), goodsNum)
	for k, gIdx := range g {
		save[k] = &NowMarket{Goods:goods.GoodsIndex(gIdx)}
		p := base.Rand(100)
		for _, nIdx := range news.GetNewsByGoods(goods.GoodsIndex(gIdx)) {
			pmax, pmin := nIdx.Pro()
			if p >= pmin && p < pmax {
				save[k].News = nIdx
				if nIdx.Style() == news.StyleGood {
					save[k].Price = save[k].Goods.Max()
				} else if nIdx.Style() == news.StyleBad {
					save[k].Price = save[k].Goods.Min()
				} else {
					save[k].Price = save[k].Goods.Normal()
				}
			}
		}
	}
}

func GetTodayNews() []*NowMarket {
	return save
}

func TodayPrice(g goods.GoodsIndex) int {
	for k := range save {
		if save[k].Goods == g {
			return save[k].Price
		}
	}

	return -1
}

func marketNewsGet() []string {
	ns := make([]string, 0)
	for k := range save {
		if save[k].News.Style() == news.StyleNormal {
			continue
		} else {
			ns = append(ns, save[k].News.Intro())
		}
	}

	if len(ns) == 0 {
		ns = append(ns, news.All.Intro())
	}

	return ns
}

func TodayNewsShow() {
	str := marketNewsGet()
	base.Notice("*******************************************")
	for k := range str {
		base.Notice(base.StrVal(k+1) + ". " + str[k])
	}
	base.Notice("*******************************************")
	base.Log(base.Info, str)
}