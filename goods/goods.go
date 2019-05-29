package goods

import "github.com/liangran2018/100million/base"

type goodsFeature struct {
	name      string
	intro     string
	maxmax    int
	maxmin    int
	normalmax int
	normalmin int
	minmax    int
	minmin    int
}

type GoodsIndex int

var goods map[GoodsIndex]goodsFeature

const (
	Phone GoodsIndex = iota
	BTC
	Wine
	Bag
	Domain
	Car
	GoodsEnd

)

func init() {
	goods = make(map[GoodsIndex]goodsFeature, GoodsEnd)
	goods[Phone] = goodsFeature{name:"", intro:"", maxmax:29888, maxmin:12888, normalmax:10888, normalmin:5288, minmax:4488, minmin:888}
	goods[BTC] = goodsFeature{name:"", intro:"", maxmax:133728, maxmin:77392, normalmax:43853, normalmin:19978, minmax:15166, minmin:7219}
	goods[Wine] = goodsFeature{name:"", intro:"", maxmax:15000, maxmin:12000, normalmax:3000, normalmin:1500, minmax:1200, minmin:600}
	goods[Bag] = goodsFeature{name:"", intro:"", maxmax:80000, maxmin:40000}



}

func (g GoodsIndex) Name() string {
	return goods[g].name
}

func (g GoodsIndex) Intro() string {
	return goods[g].intro
}

func (g GoodsIndex) Max() int {
	return base.MMrand(goods[g].maxmax, goods[g].maxmin)
}

func (g GoodsIndex) Normal() int {
	return base.MMrand(goods[g].normalmax, goods[g].normalmin)
}

func (g GoodsIndex) Min() int {
	return base.MMrand(goods[g].minmax, goods[g].minmin)
}
