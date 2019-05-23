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
	MilkTee GoodsIndex = iota
	GoodsEnd
)

func init() {
	goods = make(map[GoodsIndex]goodsFeature, GoodsEnd)
	goods[MilkTee] = goodsFeature{}
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
