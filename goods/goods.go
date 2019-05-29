package goods

import "github.com/liangran2018/100million/base"

type goodsFeature struct {
	name             string
	intro            string
	max, normal, min base.Pro
}

type GoodsIndex int

var goods map[GoodsIndex]goodsFeature

const (
	Phone GoodsIndex = iota
	BTC
	Gold
	Car
	Wine
	Domain
	Stone
	DogFood
	Bag

	GoodsEnd
)

func init() {
	goods = make(map[GoodsIndex]goodsFeature, GoodsEnd)
	goods[Phone] = goodsFeature{name: "山寨手机", intro: "华强北汹涌的人潮来了又去，弹指一挥间",
		max: base.Pro{Max: 5288, Min: 3588}, normal: base.Pro{Max: 1888, Min: 469}, min: base.Pro{Max: 888, Min: 199}}

	goods[BTC] = goodsFeature{name: "比特币", intro: "割还是被割？推荐先研读《韭菜锻炼宝典》",
		max: base.Pro{Max: 89999, Min: 69999}, normal: base.Pro{Max: 6999, Min: 2999}, min: base.Pro{Max: 3499, Min: 1599}}

	goods[Gold] = goodsFeature{name: "黄金", intro: "全民投资的时代，到底应该买什么？",
		max: base.Pro{Max: 17999, Min: 9999}, normal: base.Pro{Max: 6099, Min: 4099}, min: base.Pro{Max: 2999, Min: 1899}}

	goods[Car] = goodsFeature{name: "进口轿车", intro: "汽车产业是改革开放发展的注脚",
		max: base.Pro{Max: 139999, Min: 109999}, normal: base.Pro{Max: 52999, Min: 37999}, min: base.Pro{Max: 25999, Min: 19999}}

	goods[Wine] = goodsFeature{name: "飞天白酒", intro: "股王能否继续起飞？跌宕起伏的A股又将走向何处？",
		max: base.Pro{Max: 14999, Min: 6199}, normal: base.Pro{Max: 1899, Min: 499}, min: base.Pro{Max: 999, Min: 259}}

	goods[Domain] = goodsFeature{name: "网站域名", intro: "若回到互联网尚未兴起之时，你可以孤注一掷吗？",
		max: base.Pro{Max: 669999, Min: 559999}, normal: base.Pro{Max: 25999, Min: 19999}, min: base.Pro{Max: 13999, Min: 10099}}

	goods[Stone] = goodsFeature{name: "玉石", intro: "文化人，也可以赚钱",
		max: base.Pro{Max: 15999, Min: 8888}, normal: base.Pro{Max: 3799, Min: 2399}, min: base.Pro{Max: 2099, Min: 1199}}

	goods[DogFood] = goodsFeature{name: "狗粮", intro: "对住在微博的我来说，这是主食",
		max: base.Pro{Max: 13999, Min: 8599}, normal: base.Pro{Max: 2599, Min: 999}, min: base.Pro{Max: 1099, Min: 499}}

	goods[Bag] = goodsFeature{name: "名牌包包", intro: "让国际友人见识见识我们强盛的国力",
		max: base.Pro{Max: 89999, Min: 49999}, normal: base.Pro{Max: 29999, Min: 9999}, min: base.Pro{Max: 6999, Min: 2299}}
}

func (g GoodsIndex) Name() string {
	return goods[g].name
}

func (g GoodsIndex) Intro() string {
	return goods[g].intro
}

func (g GoodsIndex) Max() int {
	return base.MMrand(goods[g].max.Max, goods[g].max.Min)
}

func (g GoodsIndex) Normal() int {
	return base.MMrand(goods[g].normal.Max, goods[g].normal.Min)
}

func (g GoodsIndex) Min() int {
	return base.MMrand(goods[g].min.Max, goods[g].min.Min)
}
