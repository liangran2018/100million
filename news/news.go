package news

import (
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/goods"
)

type newsFeature struct {
	intro string
	style string
	promin int
	promax int
}

type newsIndex int

var news map[goods.GoodsIndex][]newsIndex
var newsStr map[newsIndex]newsFeature

var newsSave map[goods.GoodsIndex]newsIndex

const (
	All newsIndex = iota
	MikeTeeMore
	MikeTeeLess
	newsEnd
)

func init() {
	news = make(map[goods.GoodsIndex][]newsIndex, goods.GoodsEnd)
	news[goods.MilkTee] = []newsIndex{All, MikeTeeMore, MikeTeeLess}

	newsStr = make(map[newsIndex]newsFeature, newsEnd)

	newsSave = make(map[goods.GoodsIndex]newsIndex)
}

func (n newsIndex) Intro() string {
	return newsStr[n].intro
}

func (n newsIndex) Style() string {
	return newsStr[n].style
}

func (n newsIndex) pro() (max, min int) {
	p := newsStr[n]
	return p.promax, p.promin
}

func newsGet() {
	reset()
	g := base.RandMany(int(goods.GoodsEnd), 5)
	for _, gIdx := range g {
		p := base.Rand(100)
		for _, nIdx := range news[goods.GoodsIndex(gIdx)] {
			pmax, pmin := nIdx.pro()
			if p >= pmin && p < pmax {
				newsSave[goods.GoodsIndex(gIdx)] = nIdx
			}
		}
	}
}

func reset() {
	for k := range newsSave {
		delete(newsSave, k)
	}
}

func NewsShow() []string {
	newsGet()
	ns := make([]newsIndex, 0)
	for k, v := range newsSave {
		ns = append(ns, v)
	}

	news := make([]string, 0)
	for k := range news {
		if ns[k].Style() == "N" {
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

func GoodsShow() map[string]int {
	goods := make(map[string]int, 0)
	for k, v := range newsSave {
		if v.Style() == "G" {
			goods[k.Name()] = k.Max()
		} else if v.Style() == "B" {
			goods[k.Name()] = k.Min()
		} else {
			goods[k.Name()] = k.Normal()
		}
	}

	return goods
}
