package news

import "github.com/liangran2018/100million/goods"

type newsFeature struct {
	intro string
	style int
	promin int
	promax int
}

type NewsIndex int

var news map[goods.GoodsIndex][]NewsIndex
var newsStr map[NewsIndex]newsFeature

const (
	All NewsIndex = iota
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
	StyleNormal = iota
	StyleGood
	StyleBad
)

func init() {
	news = make(map[goods.GoodsIndex][]NewsIndex, goods.GoodsEnd)
	newsStr = make(map[NewsIndex]newsFeature, newsEnd)

	news[goods.Phone] = []NewsIndex{PhoneMore, PhoneLess, PhoneNormal}
	newsStr[PhoneMore] = newsFeature{intro:"山寨手机性能强大，可待机2年，自带打火机，销量高涨", style:StyleGood, promin:80, promax:100}
	newsStr[PhoneLess] = newsFeature{intro:"网购兴起，山寨电子市场一片萧条", style:StyleBad, promin:0, promax:25}
	newsStr[PhoneNormal] = newsFeature{style:StyleNormal, promin:25, promax:80}

	news[goods.BTC] = []NewsIndex{BtcMore, BtcLess, BtcNormal}
	newsStr[BtcMore] = newsFeature{intro:"数字加密货币时代到来，比特币价格疯狂上涨！", style:StyleGood, promin:90, promax:100}
	newsStr[BtcLess] = newsFeature{intro:"遭遇多国监管高压，比特币价格暴跌！", style:StyleBad, promin:0, promax:45}
	newsStr[BtcNormal] = newsFeature{style:StyleNormal, promin:45, promax:90}

	news[goods.Gold] = []NewsIndex{GoldMore, GoldLess, GoldNormal}
	newsStr[GoldMore] = newsFeature{intro:"中国大妈抢购黄金，对决资本大鳄，金价企稳回升", style:StyleGood, promin:80, promax:100}
	newsStr[GoldLess] = newsFeature{intro:"多方金融机构集体唱空黄金，金价下跌", style:StyleBad, promin:0, promax:30}
	newsStr[GoldNormal] = newsFeature{style:StyleNormal, promin:30, promax:80}

	news[goods.Car] = []NewsIndex{CarMore, CarLess, CarNormal}
	newsStr[CarMore] = newsFeature{intro:"国际汽车品牌成为土豪身份象征，价格高昂", style:StyleGood, promin:85, promax:100}
	newsStr[CarLess] = newsFeature{intro:"港口开放，海外轿车低价进口", style:StyleBad, promin:0, promax:20}
	newsStr[CarNormal] = newsFeature{style:StyleNormal, promin:20, promax:85}

	news[goods.Wine] = []NewsIndex{WineMore, WineLess, WineNormal}
	newsStr[WineMore] = newsFeature{intro:"飞天白酒一瓶难求！价格飞涨", style:StyleGood, promin:80, promax:100}
	newsStr[WineLess] = newsFeature{intro:"专家称白酒投资存在泡沫，市民纷纷抛售", style:StyleBad, promin:0, promax:30}
	newsStr[WineNormal] = newsFeature{style:StyleNormal, promin:30, promax:80}

	news[goods.Domain] = []NewsIndex{DomainMore, DomainLess, DomainNormal}
	newsStr[DomainMore] = newsFeature{intro:"互联网时代来临，珍稀域名成为天价", style:StyleGood, promin:95, promax:100}
	newsStr[DomainLess] = newsFeature{intro:"大量普通域名无人问津", style:StyleBad, promin:0, promax:40}
	newsStr[DomainNormal] = newsFeature{style:StyleNormal, promin:40, promax:95}

	news[goods.Stone] = []NewsIndex{StoneMore, StoneLess, StoneNormal}
	newsStr[StoneMore] = newsFeature{intro:"生活条件改善，越来越多的人开始收藏玉石", style:StyleGood, promin:80, promax:100}
	newsStr[StoneLess] = newsFeature{intro:"古玩市场遇冷，价格动荡", style:StyleBad, promin:0, promax:25}
	newsStr[StoneNormal] = newsFeature{style:StyleNormal, promin:25, promax:80}

	news[goods.DogFood] = []NewsIndex{DogFoodMore, DogFoodLess, DogFoodNormal}
	newsStr[DogFoodMore] = newsFeature{intro:"中美贸易战打响，各大货架上仅存的进口狗粮遭疯狂抢购", style:StyleGood, promin:75, promax:100}
	newsStr[DogFoodLess] = newsFeature{intro:"明星官宣结婚，遍洒狗粮，缓解了狗粮危机", style:StyleBad, promin:0, promax:25}
	newsStr[DogFoodNormal] = newsFeature{style:StyleNormal, promin:25, promax:75}

	news[goods.Bag] = []NewsIndex{BagMore, BagLess, BagNormal}
	newsStr[BagMore] = newsFeature{intro:"没有最贵，只有更贵！为维持高端地位，奢侈品牌持续涨价", style:StyleGood, promin:75, promax:100}
	newsStr[BagLess] = newsFeature{intro:"包治百病，为抢占中国市场，大牌包包低价倾销", style:StyleBad, promin:0, promax:30}
	newsStr[BagNormal] = newsFeature{style:StyleNormal, promin:30, promax:75}

	newsStr[All] = newsFeature{intro:"今年市场平稳"}


}

func GetNewsByGoods(g goods.GoodsIndex) []NewsIndex {
	return news[g]
}

func (n NewsIndex) Intro() string {
	return newsStr[n].intro
}

func (n NewsIndex) Style() int {
	return newsStr[n].style
}

func (n NewsIndex) Pro() (max, min int) {
	p := newsStr[n]
	return p.promax, p.promin
}
