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

func newsChoose() map[goods.GoodsIndex]newsIndex {
	g := base.RandMany(int(goods.GoodsEnd), 5)
	m := make(map[goods.GoodsIndex]newsIndex, 5)

	for _, gIdx := range g {
		p := base.Rand(100)
		for _, nIdx := range news[goods.GoodsIndex(gIdx)] {
			pmax, pmin := nIdx.pro()
			if p >= pmin && p < pmax {
				m[goods.GoodsIndex(gIdx)] = nIdx
			}
		}
	}

	return m
}

func Show() (news []string, goods map[string]int) {
	ns := make([]newsIndex, 0)
	goods = make(map[string]int, 0)
	n := newsChoose()
	for k, v := range n {
		ns = append(ns, v)
		if v.Style() == "G" {
			goods[k.Name()] = k.Max()
		} else if v.Style() == "B" {
			goods[k.Name()] = k.Min()
		} else {
			goods[k.Name()] = k.Normal()
		}
	}

	news = make([]string, 0)
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

	return
}
