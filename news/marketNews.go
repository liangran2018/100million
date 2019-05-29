package news

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/goods"
)

type newsFeature struct {
	intro string
	style int
	promin int
	promax int
}

type NowMarket struct {
	Goods goods.GoodsIndex
	News newsIndex
	Price int
}

type newsIndex int

var news map[goods.GoodsIndex][]newsIndex
var newsStr map[newsIndex]newsFeature
var save []*NowMarket

const (
	All newsIndex = iota
	PhoneNormal
	PhoneMore
	PhoneLess
	newsEnd
)

const (
	styleNormal = iota
	styleGood
	styleBetter
	styleBest
	styleBad
	styleWorse
	styleWorst
)

const goodsNum = 5

func init() {
	news = make(map[goods.GoodsIndex][]newsIndex, goods.GoodsEnd)
	news[goods.Phone] = []newsIndex{PhoneNormal, PhoneMore, PhoneLess}

	newsStr = make(map[newsIndex]newsFeature, newsEnd)
	newsStr[PhoneNormal] = newsFeature{style:styleNormal, promin:35, promax:80}
	newsStr[PhoneMore] = newsFeature{intro:"more", style:styleGood, promin:80, promax:100}
	newsStr[PhoneLess] = newsFeature{intro:"less", style:styleBad, promin:0, promax:35}
}

func (n newsIndex) Intro() string {
	return newsStr[n].intro
}

func (n newsIndex) Style() int {
	return newsStr[n].style
}

func (n newsIndex) pro() (max, min int) {
	p := newsStr[n]
	return p.promax, p.promin
}

func NewMarketNew() {
	save = make([]*NowMarket, goodsNum)

	g := base.RandMany(int(goods.GoodsEnd), goodsNum)
	for k, gIdx := range g {
		save[k] = &NowMarket{Goods:goods.GoodsIndex(gIdx)}
		p := base.Rand(100)
		for _, nIdx := range news[goods.GoodsIndex(gIdx)] {
			pmax, pmin := nIdx.pro()
			if p >= pmin && p < pmax {
				save[k].News = nIdx
				if nIdx.Style() == styleGood {
					save[k].Price = save[k].Goods.Max()
				} else if nIdx.Style() == styleBad {
					save[k].Price = save[k].Goods.Min()
				} else {
					save[k].Price = save[k].Goods.Normal()
				}
			}
		}
	}
}

func GetMarket() []*NowMarket {
	return save
}

func MarketNewsShow() []string {
	ns := make([]newsIndex, 0)
	for _, v := range save {
		ns = append(ns, v.News)
	}

	news := make([]string, 0)
	for k := range news {
		if ns[k].Style() == styleNormal {
			continue
		} else {
			news = append(news, ns[k].Intro())
		}
	}

	if len(news) == 0 {
		news = append(news, newsStr[All].intro)
	}

	return news
}

