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
	PhoneMore
	PhoneNormal
	PhoneLess
	BtcMore
	BtcNormal
	BtcLess
	GoldMore
	GoldNormal
	GoldLess
	CarMore
	CarNormal
	CarLess
	WineMore
	WineNormal
	WineLess
	DomainMore
	DomainNormal
	DomainLess
	StoneMore
	StoneNormal
	StoneLess
	DogFoodMore
	DogFoodNormal
	DogFoodLess
	BagMore
	BagNormal
	BagLess
	newsEnd
)

const (
	styleNormal = iota
	styleGood
	styleBad
)

const goodsNum = 5

func init() {
	news = make(map[goods.GoodsIndex][]newsIndex, goods.GoodsEnd)
	newsStr = make(map[newsIndex]newsFeature, newsEnd)

	news[goods.Phone] = []newsIndex{PhoneMore, PhoneLess, PhoneNormal}
	newsStr[PhoneMore] = newsFeature{intro:"山寨手机性能强大，可待机2年，自带打火机，销量高涨", style:styleGood, promin:80, promax:100}
	newsStr[PhoneLess] = newsFeature{intro:"网购兴起，山寨电子市场一片萧条", style:styleBad, promin:0, promax:25}
	newsStr[PhoneNormal] = newsFeature{style:styleNormal, promin:25, promax:80}

	news[goods.BTC] = []newsIndex{BtcMore, BtcLess, BtcNormal}
	newsStr[BtcMore] = newsFeature{intro:"数字加密货币时代到来，比特币价格疯狂上涨！", style:styleGood, promin:90, promax:100}
	newsStr[BtcLess] = newsFeature{intro:"遭遇多国监管高压，比特币价格暴跌！", style:styleBad, promin:0, promax:45}
	newsStr[BtcNormal] = newsFeature{style:styleNormal, promin:45, promax:90}

	news[goods.Gold] = []newsIndex{GoldMore, GoldLess, GoldNormal}
	newsStr[GoldMore] = newsFeature{intro:"中国大妈抢购黄金，对决资本大鳄，金价企稳回升", style:styleGood, promin:80, promax:100}
	newsStr[GoldLess] = newsFeature{intro:"多方金融机构集体唱空黄金，金价下跌", style:styleBad, promin:0, promax:30}
	newsStr[GoldNormal] = newsFeature{style:styleNormal, promin:30, promax:80}

	news[goods.Car] = []newsIndex{CarMore, CarLess, CarNormal}
	newsStr[CarMore] = newsFeature{intro:"国际汽车品牌成为土豪身份象征，价格高昂", style:styleGood, promin:85, promax:100}
	newsStr[CarLess] = newsFeature{intro:"港口开放，海外轿车低价进口", style:styleBad, promin:0, promax:20}
	newsStr[CarNormal] = newsFeature{style:styleNormal, promin:20, promax:85}

	news[goods.Wine] = []newsIndex{WineMore, WineLess, WineNormal}
	newsStr[WineMore] = newsFeature{intro:"飞天白酒一瓶难求！价格飞涨", style:styleGood, promin:80, promax:100}
	newsStr[WineLess] = newsFeature{intro:"专家称白酒投资存在泡沫，市民纷纷抛售", style:styleBad, promin:0, promax:30}
	newsStr[WineNormal] = newsFeature{style:styleNormal, promin:30, promax:80}

	news[goods.Domain] = []newsIndex{DomainMore, DomainLess, DomainNormal}
	newsStr[DomainMore] = newsFeature{intro:"互联网时代来临，珍稀域名成为天价", style:styleGood, promin:95, promax:100}
	newsStr[DomainLess] = newsFeature{intro:"大量普通域名无人问津", style:styleBad, promin:0, promax:40}
	newsStr[DomainNormal] = newsFeature{style:styleNormal, promin:40, promax:95}

	news[goods.Stone] = []newsIndex{StoneMore, StoneLess, StoneNormal}
	newsStr[StoneMore] = newsFeature{intro:"生活条件改善，越来越多的人开始收藏玉石", style:styleGood, promin:80, promax:100}
	newsStr[StoneLess] = newsFeature{intro:"古玩市场遇冷，价格动荡", style:styleBad, promin:0, promax:25}
	newsStr[StoneNormal] = newsFeature{style:styleNormal, promin:25, promax:80}

	news[goods.DogFood] = []newsIndex{DogFoodMore, DogFoodLess, DogFoodNormal}
	newsStr[DogFoodMore] = newsFeature{intro:"中美贸易战打响，各大货架上仅存的进口狗粮遭疯狂抢购", style:styleGood, promin:75, promax:100}
	newsStr[DogFoodLess] = newsFeature{intro:"明星官宣结婚，遍洒狗粮，缓解了狗粮危机", style:styleBad, promin:0, promax:25}
	newsStr[DogFoodNormal] = newsFeature{style:styleNormal, promin:25, promax:75}

	news[goods.Bag] = []newsIndex{BagMore, BagLess, BagNormal}
	newsStr[BagMore] = newsFeature{intro:"没有最贵，只有更贵！为维持高端地位，奢侈品牌持续涨价", style:styleGood, promin:75, promax:100}
	newsStr[BagLess] = newsFeature{intro:"包治百病，为抢占中国市场，大牌包包低价倾销", style:styleBad, promin:0, promax:30}
	newsStr[BagNormal] = newsFeature{style:styleNormal, promin:30, promax:75}

	newsStr[All] = newsFeature{intro:"今年市场平稳"}


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
	for k := range ns {
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

